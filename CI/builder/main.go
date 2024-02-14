//go:build mage

package main

import (
	"context"
	"dagger.io/dagger"
	//"fmt"
	//"math"
	//"math/rand"
	"os"
	"strings"
)

func appDirectory() string {
	current, err := os.Getwd()
	errorCheck(err)
	return strings.Replace(current, "CI/builder", "App", 1)
}

func Test(ctx context.Context) {
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

func Build(ctx context.Context) {
	client, errConnect := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	defer client.Close()
	errorCheck(errConnect)
	root := client.Host().Directory(appDirectory())
	_, err := client.Container().
		From("golang:latest").
		WithMountedDirectory("/App", root).
		WithWorkdir("/App").
		WithExec([]string{"go", "build -o main"}).
		Stderr(ctx)
	errorCheck(err)
}

/*
func Publish(ctx context.Context) string {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	errorCheck(err)
	root := client.Host().Directory(".")
	defer client.Close()
	root.DockerBuild().WithDirectory("/App", root)
	ref, err := image.Publish(ctx, fmt.Sprintf("ttl.sh/app-%.0f", math.Floor(rand.Float64()*10000000)))
	errorCheck(err)
	fmt.Printf("Published image to :%s\n", ref)
	return ref 
}
*/
func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}
