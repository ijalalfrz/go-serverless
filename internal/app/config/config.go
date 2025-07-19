//nolint:tagliatelle // exclude this linter until it supports UPPER_SNAKE_CASE
package config

import (
	"log/slog"
	"time"
)

type LogLeveler string

func (l LogLeveler) Level() slog.Level {
	var level slog.Level

	_ = level.UnmarshalText([]byte(l))

	return level
}

// Config holds the server configuration.
type Config struct {
	LogLevel         LogLeveler `mapstructure:"LOG_LEVEL"`
	ProfilingEnabled bool       `mapstructure:"PROFILING_ENABLED"`
	DynamoDB         DynamoDB   `mapstructure:",squash"`
	HTTP             HTTP       `mapstructure:",squash"`
	Locales          Locales    `mapstructure:",squash"`
}

type DynamoDB struct {
	Endpoint  string `mapstructure:"DYNAMODB_ENDPOINT"`
	Region    string `mapstructure:"DYNAMODB_REGION"`
	TableName string `mapstructure:"DYNAMODB_TABLE_NAME"`
}

type HTTP struct {
	Timeout       time.Duration `mapstructure:"HTTP_TIMEOUT"`
	PprofEnabled  bool          `mapstructure:"PPROF_ENABLED"`
	PprofPort     int           `mapstructure:"PPROF_PORT"`
	AllowedOrigin []string      `mapstructure:"ALLOWED_ORIGIN"`
}

type Locales struct {
	BasePath           string `mapstructure:"LOCALES_BASE_PATH"`
	SupportedLanguages string `mapstructure:"LOCALES_SUPPORTED_LANGUAGES"`
}
