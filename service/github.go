package service

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

// repoUrl: git repo url
// path: git clone path in host
func CloneRepo(repoUrl string, path string) (err error) {
	_, err = git.PlainClone(path, false, &git.CloneOptions{
		URL:      repoUrl,
		Progress: os.Stdout,
	})
	return
}

// path: repo path in host
// name: added remote name
// url: added remote url
func MirrorRepo(path, name, url string) {
	repository, err := git.PlainOpen(path)
	if err != nil {
		LogSugar.Errorf("open repo failed, err: %v\n", err)
	}
	err = repository.Fetch(&git.FetchOptions{
		RemoteName: "origin",
		Force:      true,
	})
	if err != nil {
		LogSugar.Errorf("git fetch failed, err: %v\n", err)
	}

	_, err = repository.CreateRemote(&config.RemoteConfig{
		Name: name,
		URLs: []string{url},
	})
	if err != nil {
		LogSugar.Errorf("create remote failed, err: %v\n", err)
	}
	err = repository.Push(&git.PushOptions{
		RemoteName: name,
		RefSpecs: []config.RefSpec{
			"refs/remotes/origin/*",
			"refs/heads/*",
		},
		Prune: true,
	})
	if err != nil {
		LogSugar.Errorf("push new remote failed, err: %v\n", err)
	}

}
