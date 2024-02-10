//go:build mage

package main

import (
	"context"
	"dagger.io/dagger"
	"fmt"
	"math"
	"math/rand"
	"os"
)

func Test(ctx context.Context) {
	client, errConnect := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	errorCheck(errConnect)
	root := client.Host().Directory(".")
	defer client.Close()
	_, err := client.Container().
		From("golang:latest").
		WithDirectory("/", root).
		WithWorkdir("/").
		WithExec([]string{"go", "test"}).
		Stderr(ctx)
	errorCheck(err)
}

func Build(ctx context.Context) *dagger.Container {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	errorCheck(err)
	root := client.Host().Directory(".")
	defer client.Close()
	return root.DockerBuild().WithDirectory("/App", root)
}

func Publish(ctx context.Context) string {
	image := Build(ctx)
	ref, err := image.Publish(ctx, fmt.Sprintf("ttl.sh/app-%.0f", math.Floor(rand.Float64()*10000000)))
	errorCheck(err)
	fmt.Printf("Published image to :%s\n", ref)
	return ref 
}

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}
