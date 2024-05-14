package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"ondemanddeployer/components/bashscript"
	"ondemanddeployer/constants"
	"ondemanddeployer/utils"
	"strings"
	"time"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
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
	Name      string           `json:"name"`
	Data      map[string]int64 `json:"data"`
	Timestamp time.Time        `json:"timestamp"`
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

		currentObj := GithubLanguageObj{Name: repoName, Data: make(map[string]int64), Timestamp: time.Now()}
		if err = json.Unmarshal(responseBytes, &currentObj.Data); err != nil {
			fmt.Println(string(responseBytes))
			utils.Log("Error occured while unmarshalling response: ", err.Error())
			return respJson
		}

		respJson = append(respJson, currentObj)
	}

	utils.WriteToCache(constants.GITHUB_LANGUAGES_DATA_PATH, respJson)
	return respJson
}

func FetchGithubMetaData() map[string]interface{} {
	cacheObj := utils.ReadFromCache(constants.GITHUB_OPENSOURCE_DATA_PATH)
	cacheMap, _ := cacheObj.(map[string]interface{})
	expiryDate := time.Now().Add(-constants.GITHUB_CACHE_EXPIRY_TIME)
	isExpired := true

	if fetchedAt, ok := cacheMap["fetchedAt"].(time.Time); ok && !fetchedAt.Before(expiryDate) {
		isExpired = false
	}

	if !isExpired {
		return cacheMap
	}

	client := githubv4.NewClient(oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: constants.GITHUB_TOKEN},
	)))

	variables := map[string]interface{}{
		"username": githubv4.String(constants.GITHUB_USERNAME),
	}

	allStatsQuery := AllStatsQuery{}
	if err := client.Query(context.Background(), &allStatsQuery, variables); err != nil {
		utils.Log("Error occured while hitting GraphQL APIs: ", err.Error())
		return map[string]interface{}{"message": "Error occured while querying data"}
	}

	var result map[string]interface{}

	jsonBytes, err := json.Marshal(allStatsQuery)
	if err != nil {
		utils.Log("Error occurred while marshalling graphql response into string", err.Error())
		return map[string]interface{}{"message": "Error occurred while marshalling graphql response into string"}
	}

	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		utils.Log("Error occurred while unmarshalling graphql marshalled string into map", err.Error())
		return map[string]interface{}{"message": "Error occurred while unmarshalling graphql marshalled string into map"}
	}

	result["fetchedAt"] = time.Now()
	utils.WriteToCache(constants.GITHUB_OPENSOURCE_DATA_PATH, result)

	return result
}
