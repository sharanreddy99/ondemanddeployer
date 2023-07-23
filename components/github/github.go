package github

import (
	"encoding/json"
	"net/http"
	"net/url"
	"ondemanddeployer/components/bashscript"
	"ondemanddeployer/constants"
	"ondemanddeployer/utils"
)

type GithubRepoObj struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	URL             string `json:"clone_url"`
	SVNUrl          string `json:"svn_url"`
	StargazersCount int    `json:"stargazers_count"`
	LanguagesUrl    int    `json:"languages_url"`
	PushedAt        string `json:"pushed_at"`
	Branch          string `json:"default_branch"`
	IsActive        bool   `json:"is_active"`
}

func FetchAllReposList() []GithubRepoObj {
	var err error
	var respJson []GithubRepoObj
	var response interface{}

	parsedUrl, err := url.Parse(constants.GITHUB_REPOS_LIST_URL)
	if err != nil {
		utils.Log("Error occured while parsing external API: ", err.Error())
		return respJson
	}

	req := &http.Request{Method: "GET", URL: parsedUrl}
	response = utils.ExecuteRequest(req)
	if err != nil {
		utils.Log("Error occured while parsing external API: ", err.Error())
		return respJson
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		utils.Log("Error occured while marshalling response: ", err.Error())
		return respJson
	}

	if err = json.Unmarshal(responseBytes, &respJson); err != nil {
		utils.Log("Error occured while unmarshalling response: ", err.Error())
		return respJson
	}

	for i := 0; i < len(respJson); i++ {
		respJson[i].IsActive = respJson[i].Name == bashscript.ActiveProject
	}

	return respJson
}
