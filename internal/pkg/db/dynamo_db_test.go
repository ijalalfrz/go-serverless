//go:build unit

package db

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/ijalalfrz/go-serverless/internal/app/config"
	"github.com/stretchr/testify/assert"
)

func TestInitDynamoDB(t *testing.T) {
	initDynamoDBTest := func(name string, cfg config.Config, wantErr bool) func(t *testing.T) {
		return func(t *testing.T) {
			var client *dynamodb.Client

			if wantErr {
				assert.Panics(t, func() {
					client = InitDynamoDB(cfg)
				})
				return
			}

			assert.NotPanics(t, func() {
				client = InitDynamoDB(cfg)
			})
			assert.NotNil(t, client)
		}
	}

	t.Run("success_local", initDynamoDBTest(
		"success_local",
		config.Config{
			DynamoDB: config.DynamoDB{
				Endpoint:  "http://localhost:8000",
				Region:    "ap-southeast-1",
				TableName: "test_table",
			},
		},
		false,
	))

	t.Run("success_production", initDynamoDBTest(
		"success_production",
		config.Config{
			DynamoDB: config.DynamoDB{
				Region:    "ap-southeast-1",
				TableName: "test_table",
			},
		},
		false,
	))

	t.Run("empty_region", initDynamoDBTest(
		"empty_region",
		config.Config{
			DynamoDB: config.DynamoDB{
				Endpoint: "http://localhost:8000",
				// Region intentionally left empty
				TableName: "test_table",
			},
		},
		false, // AWS SDK uses default region if empty
	))

	t.Run("invalid_endpoint", initDynamoDBTest(
		"invalid_endpoint",
		config.Config{
			DynamoDB: config.DynamoDB{
				Endpoint:  "not-a-valid-url",
				Region:    "ap-southeast-1",
				TableName: "test_table",
			},
		},
		false, // SDK validates endpoint when used, not during initialization
	))
}
