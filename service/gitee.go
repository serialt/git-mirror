package service

import (
	"github.com/parnurzeal/gorequest"
	"go.uber.org/zap"
)

var LogSugar *zap.SugaredLogger

func (c *GiteeClient) GiteeCreateRepo(repo string, private bool) (resp gorequest.Response, err []error) {
	url := "https://gitee.com/api/v5/user/repos"

	formData := map[string]string{
		"access_token": c.AccessToken,
		"name":         repo,
		"has_issues":   "true",
		"has_wiki":     "true",
		"can_comment":  "true",
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
	LogSugar.Infof("resp code: %v", resp.StatusCode)
	return

}
