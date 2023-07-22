package controllers

import (
	"ondemanddeployer/components/aws"
	"ondemanddeployer/utils"

	beego "github.com/beego/beego/v2/server/web"
)

type AWSController struct {
	beego.Controller
}

// @Title PublishSNS
// @Description Publish a message to SNS Topic
// @Success 200 {object}
// @router /publishSNS [post]
func (u *AWSController) PublishSNS() {
	aws.PublishMessage(string(u.Ctx.Input.RequestBody))
	u.Data["json"] = map[string]string{"status": "Published the message successfully"}
	u.ServeJSON()
}

// @Title SubscribeSNS
// @Description Receives the published messages from the SNS topic
// @Success 200 {object}
// @router /subscribeSNS [post]
func (u *AWSController) SubscribeSNS() {
	utils.Log(string(u.Ctx.Input.RequestBody))
	resp := aws.SubscribeMessage(u.Ctx.Request)
	u.Data["json"] = resp

	u.ServeJSON()
}
