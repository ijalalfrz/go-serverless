package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ijalalfrz/go-serverless/internal/app/dto"
	"github.com/ijalalfrz/go-serverless/internal/pkg/exception"
)

// ResponseWithBody is the common method to encode all response types to the
// client. I chose to do it this way because, since we're using JSON, there's no
// reason to provide anything more specific. It's certainly possible to
// specialize on a per-response (per-method) basis.
func ResponseWithBody(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		return fmt.Errorf("encode response body: %w", err)
	}

	return nil
}

func NoContentResponse(_ context.Context, w http.ResponseWriter, _ interface{}) error {
	w.WriteHeader(http.StatusNoContent)

	return nil
}

func CreatedResponse(_ context.Context, w http.ResponseWriter, _ interface{}) error {
	w.WriteHeader(http.StatusCreated)

	return nil
}

func ErrorResponse(ctx context.Context, err error, respWriter http.ResponseWriter) {
	var (
		appErr  exception.ApplicationError
		uiErr   string
		message string
	)

	respWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	slog.Info("error", "error", err)

	if errors.As(err, &appErr) {
		respWriter.WriteHeader(appErr.StatusCode)

		reqContext, _ := dto.RequestFromContext(ctx)
		// in case of failure to get request context, default language will be used
		message = appErr.Localize(reqContext.Language)

		uiErr = appErr.UICode

		slog.Default().Debug("error", "cause", err.Error())
	} else {
		respWriter.WriteHeader(http.StatusInternalServerError)

		uiErr = exception.InternalServerError

		message = err.Error()
	}

	//nolint:errcheck,errchkjson
	json.NewEncoder(respWriter).Encode(dto.ErrorResponse{
		Error:  message,
		UICode: uiErr,
	})
}
