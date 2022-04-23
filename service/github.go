package service

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/serialt/gopkg"
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
		RefSpecs: []config.RefSpec{
			"+refs/heads/*:refs/remotes/origin/*",
		},
	})
	if err != nil {
		LogSugar.Errorf("git fetch failed, err: %v\n", err)
	} else {
		LogSugar.Info("git fetch succeed,")
	}

	_, err = repository.CreateRemote(&config.RemoteConfig{
		Name: name,
		URLs: []string{url},
	})
	if err != nil {
		LogSugar.Errorf("create remote failed, err: %v\n", err)
	}

	// auth, _ := ssh.NewPublicKeysFromFile("git", "/Users/serialt/.ssh/id_rsa", "")

	// err = repository.Push(&git.PushOptions{
	// 	RemoteName: name,
	// 	// RefSpecs: []config.RefSpec{
	// 	// 	"+refs/heads/*:refs/remotes/origin/*",
	// 	// },
	// 	// Prune: true,
	// 	Force: true,
	// 	Auth:  auth,
	// })
	// if err != nil {
	// 	LogSugar.Errorf("push new remote failed, err: %v\n", err)
	// }
	_, err = gopkg.FindCommandPath("git")
	if err != nil {
		LogSugar.Error("command git not find")
		os.Exit(555)
	}

	gitConfigStr := fmt.Sprintf("cd %s && %s", path, "git config --global core.sshCommand \"ssh -i ~/.ssh/id_rsa -o IdentitiesOnly=yes -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no\"")
	_, err = gopkg.RunCMD(gitConfigStr)
	if err != nil {
		LogSugar.Errorf("git config UserKnownHostsFile failed, err: %v", err)
	}

	cmdStr := fmt.Sprintf("cd %s && git push --tags --force --prune %s \"refs/remotes/origin/*:refs/heads/*\" ", path, name)
	result, err := gopkg.RunCMD(cmdStr)
	if err != nil {
		LogSugar.Errorf("git push failed, err: %v", err)
	} else {
		LogSugar.Infof("git push result:%s", result)
	}
}
