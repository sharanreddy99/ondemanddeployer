package constants

import "fmt"

var (
	GITHUB_USERNAME       string = "sharanreddy99"
	GITHUB_REPOS_LIST_URL string = fmt.Sprintf("https://api.github.com/users/%s/repos", GITHUB_USERNAME)

	AWS_SNS_TOPIC_ARN string = "arn:aws:sns:us-east-1:004517456981:sharankonda"
)
