package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Github
type BaseController struct {
	beego.Controller
}

// @Title HealthCheck
// @Description Validate the health of the app
// @Success 200 {object}
// @router /healthcheck [get]
func (u *BaseController) HealthCheck() {
	u.Data["json"] = map[string]interface{}{"status": "The server is healthy"}
	u.ServeJSON()
}
