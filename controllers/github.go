package controllers

import (
	"ondemanddeployer/components/github"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Github
type GithubController struct {
	beego.Controller
}

// @Title FetchAllRepos
// @Description Fetch all repositories list from the users profile
// @Success 200 {object}
// @router /getReposList [get]
func (u *GithubController) FetchAllRepos() {
	u.Data["json"] = github.FetchAllReposList()
	u.ServeJSON()
}

// @Title FetchAllLangugages
// @Description Fetch all repositories languages list from the users profile
// @Success 200 {object}
// @router /getLanguagesList [get]
func (u *GithubController) FetchAllLanguages() {
	u.Data["json"] = github.FetchAllLanguagesList()
	u.ServeJSON()
}

// @Title FetchGithubMetaData
// @Description Fetch all repositories info from users profile
// @Success 200 {object}
// @router /getReposMetadata [get]
func (u *GithubController) FetchGithubMetaData() {
	u.Data["json"] = github.FetchGithubMetaData()
	u.ServeJSON()
}
