package db

import (
	"context"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	cfg "github.com/ijalalfrz/go-serverless/internal/app/config"
)

func InitDynamoDB(appConfig cfg.Config) *dynamodb.Client {

	opts := []func(*config.LoadOptions) error{
		config.WithRegion(appConfig.DynamoDB.Region),
	}

	if appConfig.DynamoDB.Endpoint != "" {
		opts = append(opts, config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: appConfig.DynamoDB.Endpoint}, nil
			},
		)))
	}

	awsCfg, err := config.LoadDefaultConfig(context.TODO(), opts...)
	if err != nil {
		slog.Error("failed to load default config", "error", err)
		panic(err)
	}

	return dynamodb.NewFromConfig(awsCfg)
}
