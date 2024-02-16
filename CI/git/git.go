package git

import (
    "github.com/go-git/go-git/v5"
    "github.com/go-git/go-git/v5/plumbing/transport/http"
    "github.com/go-git/go-git/v5/plumbing"
    "github.com/go-git/go-git/v5/plumbing/object"
    "os"
    "time"
    "fmt"
)

func Clone(path string, url string, token string) *git.Repository {
    repo , err := git.PlainClone(path, false, &git.CloneOptions{
        URL:      url,
        Progress: os.Stdout,
        Auth: &http.BasicAuth{
            Username: "abc123", // anything but an empty string
            Password: token,
        },
    })
    errorCheck(err)
    return repo
}

func Commit(filePath string, repo *git.Repository) plumbing.Hash {
    worktree, worktreeError := repo.Worktree()
    errorCheck(worktreeError)
    hash, addError := worktree.Add(filePath)
    fmt.Println(hash)
    errorCheck(addError)
    commit, errCommit := worktree.Commit("automatic k8s template update", &git.CommitOptions{
        Author: &object.Signature{
            Name:  "mean mug CI",
            Email: "meanMugCI@doe.org",
            When:  time.Now(),
        },
    })
    errorCheck(errCommit)
    return commit
}

func Push(repo *git.Repository, token string)  {
    err := repo.Push(&git.PushOptions{
        RemoteName: "origin",
        Auth: &http.BasicAuth{
            Username: "abc123", // anything but an empty string
            Password: token,
        },
    }) 
    errorCheck(err)
}

func errorCheck(err error) {
    if err != nil {
        panic(err)
    }
}
