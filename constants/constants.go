package constants

import (
	"fmt"
	"time"
)

var (
	HTTP_PORT string = "8080"

	GITHUB_USERNAME       string = "sharanreddy99"
	GITHUB_REPOS_LIST_URL string = fmt.Sprintf("https://api.github.com/users/%s/repos", GITHUB_USERNAME)

	AWS_SNS_TOPIC_ARN              string        = "arn:aws:sns:us-east-1:004517456981:sharankonda"
	AWS_SNS_SUBSCRIPTION_PATH      string        = "v1/aws/subscribeSNS"
	AWS_SNS_SUBSCRIPTION_WAIT_TIME time.Duration = 5 * time.Second
	AWS_INSTANCE_METADATA_ENDPOINT string        = "http://169.254.169.254/latest/meta-data"
)
