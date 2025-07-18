//go:build unit

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	t.Run("Load config successfully", func(t *testing.T) {
		config := MustInitConfig("../../../.env.sample")

		assert.Equal(t, LogLeveler("info"), config.LogLevel)
		assert.Equal(t, "http://dynamodb:8000", config.DynamoDB.Endpoint)
		assert.Equal(t, "ap-southeast-1", config.DynamoDB.Region)
		assert.Equal(t, "devices_rizal_alfarizi_local", config.DynamoDB.TableName)
	})
}
