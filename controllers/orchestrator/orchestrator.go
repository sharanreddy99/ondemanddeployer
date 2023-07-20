package github

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
