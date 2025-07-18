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
	TracingEnabled   bool       `mapstructure:"TRACING_ENABLED"`
	ProfilingEnabled bool       `mapstructure:"PROFILING_ENABLED"`
	DB               DB         `mapstructure:",squash"`
	DynamoDB         DynamoDB   `mapstructure:",squash"`
	HTTP             HTTP       `mapstructure:",squash"`
	HTTPCaller       HTTPCaller `mapstructure:",squash"`
	Locales          Locales    `mapstructure:",squash"`
}

type DynamoDB struct {
	Endpoint  string `mapstructure:"DYNAMODB_ENDPOINT"`
	Region    string `mapstructure:"DYNAMODB_REGION"`
	TableName string `mapstructure:"DYNAMODB_TABLE_NAME"`
}

type DB struct {
	DSN                   string        `mapstructure:"DB_DSN"`
	MaxOpenConnections    int           `mapstructure:"DB_MAX_OPEN_CONNECTIONS"`
	MaxIdleConnections    int           `mapstructure:"DB_MAX_IDLE_CONNECTIONS"`
	MaxConnectionLifetime time.Duration `mapstructure:"DB_MAX_CONNECTIONS_LIFETIME"`
	MaxIdleConnectionTime time.Duration `mapstructure:"DB_MAX_IDLE_CONNECTIONS_TIME"`
}

type HTTP struct {
	Port          int           `mapstructure:"HTTP_PORT"`
	Timeout       time.Duration `mapstructure:"HTTP_TIMEOUT"`
	PprofEnabled  bool          `mapstructure:"PPROF_ENABLED"`
	PprofPort     int           `mapstructure:"PPROF_PORT"`
	AllowedOrigin []string      `mapstructure:"ALLOWED_ORIGIN"`
}

type HTTPCaller struct {
	Timeout time.Duration `mapstructure:"HTTP_CALLER_TIMEOUT"`
}

type Locales struct {
	BasePath           string `mapstructure:"LOCALES_BASE_PATH"`
	SupportedLanguages string `mapstructure:"LOCALES_SUPPORTED_LANGUAGES"`
}
