package main

// import (
// 	_ "ondemanddeployer/routers"

// 	beego "github.com/beego/beego/v2/server/web"
// )

// func main() {
// 	if beego.BConfig.RunMode == "dev" {
// 		beego.BConfig.WebConfig.DirectoryIndex = true
// 		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
// 	}

// 	beego.Run()
// }

import (
	"fmt"
	"io"
	"net/http"
	"ondemanddeployer/constants"
	"ondemanddeployer/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func Subscribe() {
	var sess = session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String("ap-south-1"),
		},
	))

	var snsService = sns.New(sess)

	var ipAddress string = ""
	if resp, err := http.Get(fmt.Sprintf("%v/%v", constants.AWS_INSTANCE_METADATA_ENDPOINT, "local-ipv4")); err != nil {
		defer resp.Body.Close()

		if bodyBytes, err := io.ReadAll(resp.Body); err != nil {
			ipAddress = string(bodyBytes)
		}
	}

	inp := &sns.SubscribeInput{
		Endpoint: aws.String(fmt.Sprintf("%v/%v", ipAddress, constants.AWS_SNS_SUBSCRIPTION_PATH)),
		Protocol: aws.String("http"),
		TopicArn: aws.String(constants.AWS_SNS_TOPIC_ARN),
	}

	output, err := snsService.Subscribe(inp)
	if err != nil {
		utils.Log("Subscription Error: ", err.Error())
	}

	utils.Log("Subscription Request Successful: ", output.String())
}

func main() {
	Subscribe()
}
