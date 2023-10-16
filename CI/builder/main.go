package main

import (
	"context"
	"dagger.io/dagger"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sync"
)

func main() {
	ctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		test(ctx)
		wg.Done()
	}()
	go func() {
		build(ctx)
		wg.Done()
	}()
	wg.Wait()
}

func test(ctx context.Context) {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	errorCheck(err)
	root := client.Host().Directory(".")
	defer client.Close()
	out, err := client.Container().
		From("golang:latest").
		WithDirectory("/", root).
		WithWorkdir("/").
		WithExec([]string{"go", "test"}).
		Stderr(ctx)
	errorCheck(err)
	fmt.Println(out)
}

func build(ctx context.Context) *dagger.Container {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	errorCheck(err)
	root := client.Host().Directory("../../App/")
	defer client.Close()
	return root.DockerBuild()
}

func publish(ctx context.Context) {
	image := build(ctx)
	ref, err := image.Publish(ctx, fmt.Sprintf("ttl.sh/hello-dagger-%.0f", math.Floor(rand.Float64()*10000000))) //#nosec
	errorCheck(err)
	fmt.Printf("Published image to :%s\n", ref)
}

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}
