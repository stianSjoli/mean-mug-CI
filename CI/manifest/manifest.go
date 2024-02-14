package manifest

import (
    "gopkg.in/yaml.v3"
    "os"
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

func ReadManifest(manifestPath string) Manifest {
    data, err := os.ReadFile(manifestPath)
    errorCheck(err)
    var manifest Manifest
    errorUnmarshal := yaml.Unmarshal(data, &manifest)
    errorCheck(errorUnmarshal)
    return manifest
}
    
func UpdateManifest(manifest Manifest, imageName string) Manifest {
    copy := manifest  
    copy.Spec.Template.Spec.Containers[0].Image = imageName
    return copy
}
    
func WriteManifest(manifest Manifest, path string) {
    data, errorMarshal := yaml.Marshal(&manifest)
    errorCheck(errorMarshal)
    errorWrite := os.WriteFile(path, data, 0644)
    errorCheck(errorWrite)
}

func errorCheck(err error) {
    if err != nil {
        panic(err)
    }
}