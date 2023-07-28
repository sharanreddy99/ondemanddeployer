package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"ondemanddeployer/components/bashscript"
	"ondemanddeployer/constants"
	"ondemanddeployer/utils"
	"strings"
	"time"
)

type GithubRepoObj struct {
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	URL             string    `json:"clone_url"`
	SVNUrl          string    `json:"svn_url"`
	StargazersCount int       `json:"stargazers_count"`
	LanguagesUrl    string    `json:"languages_url"`
	PushedAt        string    `json:"pushed_at"`
	Branch          string    `json:"default_branch"`
	IsActive        bool      `json:"is_active"`
	Timestamp       time.Time `json:"timestamp"`
}

type GithubLanguageObj struct {
	Name      string         `json:"name"`
	Data      map[string]int `json:"data"`
	Timestamp time.Time      `json:"timestamp"`
}

func FetchAllReposList() []GithubRepoObj {
	var inputReposObj []GithubRepoObj = make([]GithubRepoObj, 0)

	inputObj := utils.ReadFromCache(constants.GITHUB_REPOS_DATA_PATH)
	inputReposObj, _ = inputObj.([]GithubRepoObj)

	expiryDate := time.Now().Add(-constants.GITHUB_CACHE_EXPIRY_TIME)
	isExpired := false
	isLoopRan := false

	for _, repoObj := range inputReposObj {
		isLoopRan = true
		if repoObj.Timestamp.Before(expiryDate) {
			isExpired = true
			break
		}
	}

	if !isLoopRan {
		isExpired = true
	}

	if !isExpired {
		var filteredRespJson []GithubRepoObj = make([]GithubRepoObj, 0)
		for _, allowedRepo := range constants.GITHUB_ALLOWED_REPOS {
			for i := 0; i < len(inputReposObj); i++ {
				inputReposObj[i].IsActive = inputReposObj[i].Name == bashscript.ActiveProject
				inputReposObj[i].Timestamp = time.Now()
				if inputReposObj[i].Name == allowedRepo {
					filteredRespJson = append(filteredRespJson, inputReposObj[i])
					break
				}
			}
		}

		fmt.Println("cached apis")
		return filteredRespJson
	}

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

	var filteredRespJson []GithubRepoObj = make([]GithubRepoObj, 0)

	for _, allowedRepo := range constants.GITHUB_ALLOWED_REPOS {
		for i := 0; i < len(respJson); i++ {
			respJson[i].IsActive = respJson[i].Name == bashscript.ActiveProject
			respJson[i].Timestamp = time.Now()
			if respJson[i].Name == allowedRepo {
				filteredRespJson = append(filteredRespJson, respJson[i])
				break
			}
		}
	}

	utils.WriteToCache(constants.GITHUB_REPOS_DATA_PATH, filteredRespJson)

	fmt.Println("New APIS")
	return filteredRespJson
}

func FetchAllLanguagesList() []GithubLanguageObj {
	var inputLanguagesObj []GithubLanguageObj = make([]GithubLanguageObj, 0)

	inputObj := utils.ReadFromCache(constants.GITHUB_LANGUAGES_DATA_PATH)
	inputLanguagesObj, _ = inputObj.([]GithubLanguageObj)

	expiryDate := time.Now().Add(-constants.GITHUB_CACHE_EXPIRY_TIME)
	isExpired := false
	isLoopRan := false

	for _, languageObj := range inputLanguagesObj {
		isLoopRan = true
		if languageObj.Timestamp.Before(expiryDate) {
			isExpired = true
			break
		}
	}

	if !isLoopRan {
		isExpired = true
	}

	if !isExpired {
		fmt.Println("cached apis")
		return inputLanguagesObj
	}

	var respJson []GithubLanguageObj = make([]GithubLanguageObj, 0)
	var response interface{}

	for _, repoName := range constants.GITHUB_ALLOWED_REPOS {
		parsedUrl, err := url.Parse(strings.Replace(constants.GITHUB_LANGUAGE_LIST_URL, "{{repoName}}", repoName, 1))
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

		currentObj := GithubLanguageObj{Name: repoName, Data: make(map[string]int), Timestamp: time.Now()}

		if err = json.Unmarshal(responseBytes, &currentObj.Data); err != nil {
			utils.Log("Error occured while unmarshalling response: ", err.Error())
			return respJson
		}

		respJson = append(respJson, currentObj)
	}

	utils.WriteToCache(constants.GITHUB_LANGUAGES_DATA_PATH, respJson)

	fmt.Println("new apis")
	return respJson
}
