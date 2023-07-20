package aws

import (
	"ondemanddeployer/constants"
	"ondemanddeployer/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

var snsClient *sns.SNS

func PublishMessage(message string) {
	messageObj := &sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String(constants.AWS_SNS_TOPIC_ARN),
	}

	res, err := snsClient.Publish(messageObj)
	if err != nil {
		utils.Log("Publishing SNS Error: ", err.Error())
	}

	utils.Log("Message published successfully: ", message, *res.MessageId)
}

func init() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	snsClient = sns.New(sess)
}
