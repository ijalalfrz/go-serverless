package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/ijalalfrz/go-serverless/internal/pkg/exception"
	"github.com/ijalalfrz/go-serverless/internal/pkg/lang"
)

// ErrInvalidType invalid type of request.
var ErrInvalidType = exception.ApplicationError{
	Localizable: lang.Localizable{
		Message: "invalid type",
	},
	StatusCode: exception.CodeBadRequest,
}

type Device struct {
	CreateDevice  endpoint.Endpoint
	GetDeviceByID endpoint.Endpoint
}

type Endpoint struct {
	Device
}
