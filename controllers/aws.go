package controllers

import (
	"encoding/json"
	"fmt"
	"ondemanddeployer/components/aws"

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
	var message string
	json.Unmarshal(u.Ctx.Input.RequestBody, &message)
	aws.PublishMessage(message)
	u.Data["json"] = map[string]string{"status": "Published the message successfully"}
	u.ServeJSON()
}

// @Title SubscribeSNS
// @Description Receives the published messages from the SNS topic
// @Success 200 {object}
// @router /subscribeSNS [post]
func (u *AWSController) SubscribeSNS() {
	fmt.Println(string(u.Ctx.Input.RequestBody))
	u.Data["json"] = map[string]string{"status": "Received the message successfully"}
	u.ServeJSON()
}
