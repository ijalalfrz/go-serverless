package config

import (
	"fmt"
	"log/slog"

	"github.com/spf13/viper"
)

// MustInitConfig initializes configuration from .env file and returns config structure.
func MustInitConfig(configFile string) Config {
	var (
		vpr = viper.New()
		cfg Config
	)

	// envList := []string{
	// 	"LOG_LEVEL",
	// 	"TRACING_ENABLED",
	// 	"PROFILING_ENABLED",
	// 	"DYNAMODB_ENDPOINT",
	// 	"DYNAMODB_REGION",
	// 	"DYNAMODB_TABLE_NAME",
	// 	"PPROF_ENABLED",
	// 	"ALLOWED_ORIGIN",
	// 	"HTTP_CALLER_TIMEOUT",
	// 	"LOCALES_BASE_PATH",
	// 	"LOCALES_SUPPORTED_LANGUAGES",
	// }
	// for _, env := range envList {
	// 	vpr.BindEnv(env)
	// }

	vpr.AutomaticEnv()

	if configFile != "" {
		vpr.SetConfigFile(configFile)
		vpr.SetConfigType("env")

		if err := vpr.ReadInConfig(); err != nil {
			slog.Error("cannot read config file", slog.String("error", err.Error()))
			panic(err)
		}
	}

	if err := vpr.Unmarshal(&cfg); err != nil {
		slog.Error("cannot unmarshal config file", slog.String("error", err.Error()))

		panic(err)
	}

	vpr.WatchConfig()

	slog.Info("cfg", slog.String("config", fmt.Sprintf("%v", cfg)))

	return cfg
}
