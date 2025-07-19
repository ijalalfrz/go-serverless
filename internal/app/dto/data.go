package dto

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ErrorResponse response payload.
type ErrorResponse struct {
	Error  string `json:"error"`
	UICode string `json:"uiCode"` //nolint:tagliatelle
}
