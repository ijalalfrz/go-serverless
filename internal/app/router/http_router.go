package router

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/ijalalfrz/go-serverless/internal/app/config"
	"github.com/ijalalfrz/go-serverless/internal/app/dto"
	"github.com/ijalalfrz/go-serverless/internal/app/endpoint"
	httptransport "github.com/ijalalfrz/go-serverless/internal/pkg/transport/http"
)

// MakeHTTPRouter builds the HTTP router with all the service endpoints.
func MakeHTTPRouter(
	endpts endpoint.Endpoint,
	cfg config.Config,
) *chi.Mux {
	// Initialize Router
	router := chi.NewRouter()

	router.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	router.Route("/api", func(router chi.Router) {
		router.Use(
			httptransport.LoggingMiddleware(slog.Default()),
			httptransport.CORSMiddleware(cfg.HTTP.AllowedOrigin),
			httptransport.Recoverer(slog.Default()),
			render.SetContentType(render.ContentTypeJSON),
		)

		router.Route("/devices", func(router chi.Router) {
			router.Post("/", httptransport.MakeHandlerFunc(
				endpts.Device.CreateDevice,
				httptransport.DecodeRequest[dto.CreateDeviceRequest],
				httptransport.CreatedResponse,
			))

			router.Get("/{id}", httptransport.MakeHandlerFunc(
				endpts.Device.GetDeviceByID,
				httptransport.DecodeRequest[dto.GetDeviceByIDRequest],
				httptransport.ResponseWithBody,
			))
		})
	})

	return router
}
