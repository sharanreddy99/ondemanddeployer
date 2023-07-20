package controllers

import (
	"encoding/json"
	"ondemanddeployer/components/aws"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Github
type AWSController struct {
	beego.Controller
}

// @Title PublishSNS
// @Description Publish a message to SNS Topic
// @Success 200 {object}
// @router /publishSNS [post]
func (u *AWSController) PublishSNS() {
	var message string
	json.Unmarshal(u.Ctx.Input.RequestBody, &message)
	aws.PublishMessage(message)
	u.Data["json"] = map[string]string{"status": "Published the message successfully"}
	u.ServeJSON()
}
