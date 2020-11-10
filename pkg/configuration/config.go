package configuration

import (
	"flag"

	"github.com/codingnagger/cloudready-todolist-blog-example/pkg/storage/aws"
)

// Config is the structure used for configuration
type Config struct {
	AwsConfig aws.TodoConfig
}

// BuildConfigFromFlags defines and parses flags in order to generate a configuration that is returned in a configuration.Config struct
func BuildConfigFromFlags() Config {
	var tableName string
	var awsEndpoint string

	flag.StringVar(&tableName, "tableName", "TodoList", "DynamoDB table")
	flag.StringVar(&awsEndpoint, "endpoint", "", "AWS endpoint (used for local dev)")

	flag.Parse()

	return Config{
		AwsConfig: aws.TodoConfig{
			TableName: tableName,
			Endpoint:  awsEndpoint},
	}
}
