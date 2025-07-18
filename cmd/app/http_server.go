package app

import (
	"context"
	"log/slog"
	"net/http/pprof"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	chiv5adapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"github.com/go-chi/chi/v5"
	"github.com/ijalalfrz/go-serverless/internal/app/config"
	"github.com/ijalalfrz/go-serverless/internal/app/endpoint"
	"github.com/ijalalfrz/go-serverless/internal/app/repository"
	"github.com/ijalalfrz/go-serverless/internal/app/router"
	"github.com/ijalalfrz/go-serverless/internal/app/service"
	"github.com/ijalalfrz/go-serverless/internal/pkg/db"
	"github.com/ijalalfrz/go-serverless/internal/pkg/lang"
	"github.com/ijalalfrz/go-serverless/internal/pkg/logger"
	"github.com/spf13/cobra"
)

var timeout = 30 * time.Second

type handler func(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error)

var httpServerCmd = &cobra.Command{
	Use:   "http",
	Short: "Serve incoming requests from REST HTTP/JSON API",
	Run: func(_ *cobra.Command, _ []string) {
		slog.Debug("command line flags", slog.String("config_path", cfgFilePath))
		cfg := config.MustInitConfig(cfgFilePath)

		logger.InitStructuredLogger(cfg.LogLevel)

		lambda.Start(getLambdaHandler(cfg))
	},
}

func getLambdaHandler(cfg config.Config) handler {
	lang.SetSupportedLanguages(cfg.Locales.SupportedLanguages)
	lang.SetBasePath(cfg.Locales.BasePath)

	endpts := makeEndpoints(cfg)

	router := router.MakeHTTPRouter(
		endpts,
		cfg,
	)

	// Add pprof routes if enabled
	if cfg.HTTP.PprofEnabled {
		pprofRouter := chi.NewRouter()
		pprofRouter.HandleFunc("/debug/pprof/*", pprof.Index)
		pprofRouter.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		pprofRouter.HandleFunc("/debug/pprof/profile", pprof.Profile)
		pprofRouter.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		pprofRouter.HandleFunc("/debug/pprof/trace", pprof.Trace)
		pprofRouter.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
		pprofRouter.Handle("/debug/pprof/heap", pprof.Handler("heap"))
		pprofRouter.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
		pprofRouter.Handle("/debug/pprof/block", pprof.Handler("block"))
		pprofRouter.Handle("/debug/pprof/mutex", pprof.Handler("mutex"))

		router.Mount("/", pprofRouter)
	}

	chiLambda := chiv5adapter.NewV2(router)

	return func(ctx context.Context,
		req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		return chiLambda.ProxyWithContextV2(ctx, req)
	}
}

func makeEndpoints(cfg config.Config) endpoint.Endpoint {
	dbConn := db.InitDynamoDB(cfg)

	// init all repo
	deviceRepository := repository.NewDeviceRepository(dbConn, cfg.DynamoDB.TableName)

	return endpoint.Endpoint{
		Device: makeDeviceEndpoints(deviceRepository),
	}
}

func makeDeviceEndpoints(deviceRepository *repository.DeviceRepository) endpoint.Device {
	// init device service
	deviceSvc := service.NewDeviceService(deviceRepository)

	return endpoint.NewDeviceEndpoint(deviceSvc)
}
