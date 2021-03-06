package service

import (
	"github.com/parnurzeal/gorequest"
	"go.uber.org/zap"
)

var LogSugar *zap.SugaredLogger

func (c *GiteeClient) GiteeCreateRepo(repo string, private string, description string) (resp gorequest.Response, err []error) {
	url := "https://gitee.com/api/v5/user/repos"

	formData := map[string]string{
		"access_token": c.AccessToken,
		"name":         repo,
		"has_issues":   "true",
		"has_wiki":     "true",
		"can_comment":  "true",
		"description":  description,
		"private":      private,
		"path":         repo,
	}

	request := gorequest.New()

	resp, _, err = request.Post(url).
		Set("Content-Type", "application/json;charset=UTF-8").
		SendMap(formData).
		End()

	if err != nil {
		LogSugar.Errorf("request failed, reponame: %v  err: %v ", repo, err)
	}
	if resp.StatusCode == 422 {
		LogSugar.Infof("repo exist")
	} else {
		LogSugar.Infof("resp code: %v", resp.StatusCode)
	}

	return

}
