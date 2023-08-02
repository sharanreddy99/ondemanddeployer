package constants

import (
	"fmt"
	"time"
)

var (
	HTTP_PORT string = "8080"

	GITHUB_USERNAME            string        = "sharanreddy99"
	GITHUB_REPOS_LIST_URL      string        = fmt.Sprintf("https://api.github.com/users/%s/repos", GITHUB_USERNAME)
	GITHUB_LANGUAGE_LIST_URL   string        = fmt.Sprintf("https://api.github.com/repos/%s/{{repoName}}/languages", GITHUB_USERNAME)
	GITHUB_ALLOWED_REPOS       []string      = []string{"housemate", "securechat", "ceta", "elevator_multithreading", "image_slideshow", "polls_viewer"}
	GITHUB_LANGUAGES_DATA_PATH string        = "data/github/languages.json"
	GITHUB_REPOS_DATA_PATH     string        = "data/github/repos.json"
	GITHUB_CACHE_EXPIRY_TIME   time.Duration = 24 * time.Hour

	AWS_REGION                     string        = "us-east-1"
	AWS_SNS_TOPIC_ARN              string        = "arn:aws:sns:us-east-1:004517456981:sharankonda"
	AWS_SNS_PROTOCOL               string        = "https"
	AWS_SNS_HOST_IP                string        = "project.sharankonda.com"
	AWS_SNS_SUBSCRIPTION_PATH      string        = "v1/aws/subscribeSNS"
	AWS_SNS_SUBSCRIPTION_WAIT_TIME time.Duration = 5 * time.Second
	AWS_INSTANCE_METADATA_ENDPOINT string        = "http://169.254.169.254/latest/meta-data"
)
