package main

import (
	//"fmt"
	"os"
	"gopkg.in/yaml.v3"
	"flag"
	//"github.com/go-git/go-git/v5"
	//"github.com/go-git/go-git/v5/plumbing/object"
	//"time"
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
/*
func commitChanges(file string) {
	repo, errGit := git.PlainOpen("../../.git")
	errorCheck(errGit)
	//w, errHead := repo.
	//errorCheck(errHead)
	//w.Add(file)
	repo.WorkTree.Add(file)
	commit, errCommit := repo.Commit("automatic k8s template update", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "mean mug CI",
			Email: "meanMugCI@doe.org",
			When:  time.Now(),
		},
	})
	errorCheck(errCommit)
	commitObject, errCommitObject := repo.CommitObject(commit)
	errorCheck(errCommitObject)
	fmt.Println(commitObject)
}
*/
func main() {	
	image := flag.String("image", "", "docker image with sha")
	manifestPath := flag.String("manifestPath", "", "relative location to manifest file")
	flag.Parse()
	manifest := readManifest(*manifestPath)
	new_manifest := updateManifest(manifest, *image)
	writeManifest(new_manifest, *manifestPath)
	//commitChanges(*manifestPath)
}



func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}
