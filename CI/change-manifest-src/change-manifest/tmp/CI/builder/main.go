package main

import (
	"context"
	"dagger.io/dagger"
	"fmt"
	"math"
	"math/rand"
	"os"
	"gopkg.in/yaml.v3"
)



type Port struct {
	ContainerPort int `yaml:"containerPort"`
}
	
type Container struct {
	Name string `yaml:"name"`
	Image string `yaml:"image"`
	Ports []Port `yaml:"ports"`
}

type TemplateSpec struct {
	Containers  []Container `yaml:"containers"`
}

type TemplateMetadata struct {
	Labels Labels
}
	
type Template struct {
	Metadata  TemplateMetadata `yaml:"metadata"`
	Spec TemplateSpec `yaml:"spec"`
}
type Selector struct {
	MatchLabels  Labels `yaml:"matchLabels"`
}

type Spec struct {
	Replias int `yaml:"replicas"`
	Selector Selector `yaml:"selector"`
	Template Template `yaml:"template"`
}
	
type Labels struct {
	App string `yaml:"app"`
}		
	
type Metadata struct {
	Name string `yaml:"name"`
	Namespace string `yaml:"namespace"`
	Labels Labels
}

type Manifest struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind string `yaml:"kind"`
	Metadata Metadata `yaml:"metadata"`
	Spec Spec `yaml:"spec"`
}
	
func readManifest(manifestPath string) Manifest {
	data, err := os.ReadFile(manifestPath)
	errorCheck(err)
	var manifest Manifest
	errorUnmarshal := yaml.Unmarshal(data, &manifest)
	errorCheck(errorUnmarshal)
	return manifest
}
	
func updateManifest(manifest Manifest, imageName string) Manifest {
	copy := manifest  
	copy.Spec.Template.Spec.Containers[0].Image = imageName
	return copy
}
	
func writeManifest(manifest Manifest, path string) {
	data, errorMarshal := yaml.Marshal(&manifest)
	errorCheck(errorMarshal)
	errorWrite := os.WriteFile(path, data, 0644)
	errorCheck(errorWrite)
}	

func main() {
	ctx := context.Background()
	imageRef := publish(ctx)
	manifestPath := "../../ArgoCD/deployment.yml"
	manifest := readManifest(manifestPath)
	new_manifest := updateManifest(manifest, imageRef)
	writeManifest(new_manifest, manifestPath)
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
	root := client.Host().Directory(".")
	defer client.Close()
	return root.DockerBuild().WithDirectory("/App", root)
}

func publish(ctx context.Context) string {
	image := build(ctx)
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
