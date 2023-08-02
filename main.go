package main

import (
	"fmt"
	"io"
	"net/http"
	"ondemanddeployer/constants"
	"ondemanddeployer/utils"
	"time"

	_ "ondemanddeployer/components/bashscript"
	_ "ondemanddeployer/routers"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func Subscribe() {
	var sess = session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String(constants.AWS_REGION),
		},
	))

	var snsService = sns.New(sess)

	var ipAddress string = ""
	if resp, err := http.Get(fmt.Sprintf("%v/%v", constants.AWS_INSTANCE_METADATA_ENDPOINT, "public-ipv4")); err == nil {
		defer resp.Body.Close()

		if bodyBytes, err := io.ReadAll(resp.Body); err == nil {
			ipAddress = string(bodyBytes)
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}

	inp := &sns.SubscribeInput{
		Endpoint: aws.String(fmt.Sprintf("https://%v:%v/%v", ipAddress, constants.HTTP_PORT, constants.AWS_SNS_SUBSCRIPTION_PATH)),
		Protocol: aws.String("https"),
		TopicArn: aws.String(constants.AWS_SNS_TOPIC_ARN),
	}

	output, err := snsService.Subscribe(inp)
	if err != nil {
		utils.Log("Subscription Error: ", err.Error())
	}

	utils.Log("Subscription Request Successful: ", output.String())
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	} else {
		time.AfterFunc(constants.AWS_SNS_SUBSCRIPTION_WAIT_TIME, func() {
			Subscribe()
		})
	}

	utils.Log("Server is up")

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
	}))

	beego.Run()
}
