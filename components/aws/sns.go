package aws

import (
	"context"
	"encoding/json"
	"net/http"
	"ondemanddeployer/constants"
	"ondemanddeployer/utils"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

var snsClient *sns.SNS

type RequestType string

type SubscriptionConfirmationRequest struct {
	Type              RequestType
	MessageId         string
	TopicArn          string
	Message           string
	Timestamp         string
	SignatureVersion  string
	SigningCertURL    string
	SubscribeURL      string
	Token             string
	MessageAttributes MessageAttribute
}

type NotificationRequest struct {
	Type              RequestType
	MessageId         string
	TopicArn          string
	Message           string
	Timestamp         string
	SignatureVersion  string
	SigningCertURL    string
	MessageAttributes MessageAttribute
}

type MessageAttribute struct {
	Event Event
}

type Event struct {
	Type  string
	Value EventValue
}

type EventValue string

func (s *SubscriptionConfirmationRequest) Bind(r *http.Request) error {
	return json.NewDecoder(r.Body).Decode(s)
}

func (n *NotificationRequest) Bind(r *http.Request) error {
	return json.NewDecoder(r.Body).Decode(n)
}

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

// Identify incoming notification request and process it.
func SubscribeMessage(r *http.Request) map[string]interface{} {
	var resp map[string]interface{}
	var err error

	switch r.Header.Get("X-Amz-Sns-Message-Type") {
	case "SubscriptionConfirmation":
		err = confirmSubscription(r)
	case "Notification":
		err = handleReceivedMessage(r)
	default:
		resp["status"] = "FAIL"
		resp["message"] = "invalid message type"
	}

	if err != nil {
		resp["status"] = "FAIL"
		resp["message"] = err.Error()
		utils.Log("Error while handling message's type: ", err.Error())
	} else {
		resp["status"] = "SUCCESS"
		resp["message"] = "Successfully handled notification"
	}

	return resp
}

// Subscription confirmation. Takes place only once. Once confirmed, duplicated
// CLI subscription command executions will not trigger another request.
func confirmSubscription(r *http.Request) error {
	var req SubscriptionConfirmationRequest
	if err := req.Bind(r); err != nil {
		utils.Log("Request Binding: %w", err)
		return err
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()

	if _, err := snsClient.ConfirmSubscriptionWithContext(ctx, &sns.ConfirmSubscriptionInput{
		Token:    aws.String(req.Token),
		TopicArn: aws.String(req.TopicArn),
	}); err != nil {
		utils.Log("Subscription Cnfirmation Error: ", err)
		return err
	}

	utils.Log("Subscription Confirmation Successful!")
	return nil
}

// Consumes published events. Called as many times as the client publishes.
func handleReceivedMessage(r *http.Request) error {
	// var req NotificationRequest
	// if err := req.Bind(r); err != nil {
	// 	return fmt.Errorf("request binding: %w", err)
	// }

	// switch req.MessageAttributes.Event.Value {
	// case Upload:
	// 	log.Printf("uploading %s ...", req.Message)
	// case Download:
	// 	log.Printf("downloading %s ...", req.Message)
	// default:
	// 	return fmt.Errorf("unknwon event value")
	// }

	return nil
}

func init() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	snsClient = sns.New(sess)
}
