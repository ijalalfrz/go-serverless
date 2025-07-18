package db

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	cfg "github.com/ijalalfrz/go-serverless/internal/app/config"
)

func InitDynamoDB(appConfig cfg.Config) *dynamodb.Client {
	fmt.Println("appConfig.DynamoDB.Endpoint", appConfig.DynamoDB.Endpoint)
	fmt.Println("appConfig.DynamoDB.Region", appConfig.DynamoDB.Region)
	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(appConfig.DynamoDB.Region),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: appConfig.DynamoDB.Endpoint}, nil
			},
		)),
	)
	if err != nil {
		slog.Error("failed to load default config", "error", err)
		panic(err)
	}

	return dynamodb.NewFromConfig(awsCfg)
}
