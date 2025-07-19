package dto

import (
	"net/http"

	"github.com/ijalalfrz/go-serverless/internal/pkg/exception"
	"github.com/ijalalfrz/go-serverless/internal/pkg/lang"
)

const (
	InvalidRequestDevicePrefix      = "INVALID_REQUEST_DEVICE_PREFIX"
	InvalidRequestDeviceModelPrefix = "INVALID_REQUEST_DEVICE_MODEL_PREFIX"
	RequiredDeviceID                = "REQUIRED_DEVICE_ID_PARAM"
)

func NewInvalidRequestError(err error, uiCode string) exception.ApplicationError {
	return exception.ApplicationError{
		StatusCode: http.StatusBadRequest,
		Localizable: lang.Localizable{
			MessageID: "errors.invalid_request",
			MessageVars: map[string]interface{}{
				"message": err.Error(),
			},
		},
		UICode: uiCode,
		Cause:  err,
	}
}
