//go:build mage

package main

import (
	"context"
	"dagger.io/dagger"
	"os"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"example.com/git"
    "example.com/manifest"
	"github.com/magefile/mage/mg"
)

type App mg.Namespace



func appDirectory() string {
	current, err := os.Getwd()
	errorCheck(err)
	return strings.Replace(current, "CI/builder", "App", 1)
}

func (App) Test(ctx context.Context) {
	client, errConnect := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	defer client.Close()
	errorCheck(errConnect)
	root := client.Host().Directory(appDirectory())
	_, err := client.Container().
		From("golang:latest").
		WithMountedDirectory("/App", root).
		WithWorkdir("/App").
		WithExec([]string{"go", "test"}).
		Stderr(ctx)
	errorCheck(err)
}

func (App) Build(ctx context.Context) {
	client, errConnect := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	defer client.Close()
	errorCheck(errConnect)
	root := client.Host().Directory(appDirectory())
	_, err := client.Container().
		From("golang:latest").
		WithMountedDirectory("/App", root).
		WithWorkdir("/App").
		WithExec([]string{"go", "build"}).
		Stderr(ctx)
	errorCheck(err)
}

func (App) Deploy(ctx context.Context, token string) {
	client, errConnect := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	errorCheck(errConnect)
	root := client.Host().Directory(appDirectory())
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
		ref, err := prodImage.Publish(ctx, fmt.Sprintf("ttl.sh/app-%.0f", math.Floor(rand.Float64()*10000000))) //default time to live is 1 hour
		errorCheck(err)
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
    errorCheck(errRemove)
}

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}
