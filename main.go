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
	"ondemanddeployer/constants"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/beego/beego"

	"fmt"
	"os"
)

func Middleware(componentName string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {

		opts := []nethttp.MWOption{
			nethttp.MWComponentName(componentName),
			nethttp.OperationNameFunc(func(r *http.Request) string {
				return "HTTP " + r.Method + " " + r.URL.Path
			}),
			nethttp.MWURLTagFunc(func(u *url.URL) string {
				return u.String()
			}),
			nethttp.MWSpanFilter(func(r *http.Request) bool {
				return true
			}),
			nethttp.MWSpanObserver(func(span opentracing.Span, r *http.Request) {

			}),
		}

		return nethttp.Middleware(opentracing.GlobalTracer(), h, opts...)
	}
}

func main() {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	beego.RunWithMiddleWares = 

	svc := sns.New(sess)
	

	// svc.Subsc

	// result, err := svc.Publish(&sns.PublishInput{
	// 	Message:  aws.String("hello sharan"),
	// 	TopicArn: aws.String(constants.AWS_SNS_TOPIC_ARN),
	// })

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(1)
	// }

	// fmt.Printf("%+v", result)
}
