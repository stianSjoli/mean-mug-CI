package main

import (
    "fmt"
    "example.com/git"
    "example.com/manifest"
    "flag"
)


func main() {
    token := flag.String("token", "", "git token with write access")
    manifestfile := flag.String("manifestfile", ".", "k8 file with full filepath")
    repoUrl := flag.String("repoUrl", "", "url to remote git repository")
    imageRef := flag.String("imageRef", "", "imageRef to new image")
    flag.Parse()
    if(*token == "" || *repoUrl == "") {
        fmt.Println("missing input values")
    } else {
        dirName := "./tmp"
        manifestPath := dirName + "/" + *manifestfile
        repo := git.Clone(dirName, *repoUrl, *token)
        m := manifest.ReadManifest(manifestPath)
        fmt.Println(m)
        updatedM := manifest.UpdateManifest(m, *imageRef)
        manifest.WriteManifest(updatedM, manifestPath)
        fmt.Println(git.Commit(*manifestfile,repo))
        git.Push(repo, *token)
    }
}