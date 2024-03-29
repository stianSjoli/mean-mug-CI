//go:build mage

package main

import (
	"context"
	"dagger.io/dagger"
	"os"
	"sync"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"example.com/git"
    "example.com/manifest"
)

type TargetContext struct {
	cancel context.CancelFunc
	ctx context.Context
}

func appDirectory(cancel context.CancelFunc) string {
	current, err := os.Getwd()
	errorCheck(err, cancel)
	return strings.Replace(current, "CI/builder", "App", 1)
}

func test(tc TargetContext, wg *sync.WaitGroup) {
	client, errConnect := dagger.Connect(tc.ctx, dagger.WithLogOutput(os.Stderr))
	defer client.Close()
	errorCheck(errConnect, tc.cancel)
	root := client.Host().Directory(appDirectory(tc.cancel))
	_, err := client.Container().
		From("golang:latest").
		WithMountedDirectory("/App", root).
		WithWorkdir("/App").
		WithExec([]string{"go", "test"}).
		Stderr(tc.ctx)
	errorCheck(err, tc.cancel)
	wg.Done()
}

func build(tc TargetContext, wg *sync.WaitGroup) {
	client, errConnect := dagger.Connect(tc.ctx, dagger.WithLogOutput(os.Stderr))
	defer client.Close()
	errorCheck(errConnect, tc.cancel)
	root := client.Host().Directory(appDirectory(tc.cancel))
	_, err := client.Container().
		From("golang:latest").
		WithMountedDirectory("/App", root).
		WithWorkdir("/App").
		WithExec([]string{"go", "build"}).
		Stderr(tc.ctx)
	errorCheck(err, tc.cancel)
	wg.Done()
}

func deploy(tc TargetContext, token string) {
	client, errConnect := dagger.Connect(tc.ctx, dagger.WithLogOutput(os.Stderr))
	errorCheck(errConnect, tc.cancel)
	root := client.Host().Directory(appDirectory(tc.cancel))
	defer client.Close()
	imageRef := make(chan string)
	go func() {
		builder := client.Container().
			From("golang:latest").
			WithDirectory("/src", root).
			WithWorkdir("/src").
			WithEnvVariable("CGO_ENABLED", "0").
			WithExec([]string{"go", "build", "-o", "app"})

		prodImage := client.Container().
			From("alpine").
			WithFile("/bin/app", builder.File("/src/app")).
			WithEntrypoint([]string{"/bin/app"})
		ref, err := prodImage.Publish(tc.ctx, fmt.Sprintf("ttl.sh/app-%.0f", math.Floor(rand.Float64()*10000000))) //default time to live is 1 hour
		errorCheck(err, tc.cancel)
		imageRef <- ref 
	}()
	
	dirPath := "./tmp"
	manifestPath := "ArgoCD/deployment.yml"
	repoUrl := "https://github.com/stianSjoli/mean-mug-CI.git"
	repo := git.Clone(dirPath, repoUrl, token)
    currentManifest := manifest.ReadManifest(dirPath + "/" + manifestPath)
    newManifest := manifest.UpdateManifest(currentManifest, <- imageRef)
    manifest.WriteManifest(newManifest, dirPath + "/" + manifestPath)
    git.Commit(manifestPath, repo)
    git.Push(repo, token)
    errRemove := os.RemoveAll(dirPath)
    errorCheck(errRemove, tc.cancel)
}

func CI() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(2)
	tc := TargetContext{cancel:cancel, ctx: ctx}
	go test(tc, &wg)
	go build(tc, &wg)
	wg.Wait()
}

func CD(token string) {
	CI()
	ctx, cancel := context.WithCancel(context.Background())
	tc := TargetContext{cancel:cancel, ctx: ctx}
	deploy(tc, token)
}
func errorCheck(err error, cancel context.CancelFunc) {
	if err != nil {
		cancel()
	}
}
