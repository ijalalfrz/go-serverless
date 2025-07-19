package dto

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// CreateDeviceRequest is the request body for the CreateDevice endpoint.
type CreateDeviceRequest struct {
	ID          string `json:"id"          validate:"required"`
	DeviceModel string `json:"deviceModel" validate:"required"` //nolint:tagliatelle
	Name        string `json:"name"        validate:"required"`
	Note        string `json:"note"        validate:"required"`
	Serial      string `json:"serial"      validate:"required"`
}

func (r *CreateDeviceRequest) Bind(_ *http.Request) error {
	if err := validate.Struct(r); err != nil {
		return NewInvalidRequestError(err, InvalidRequestDevicePrefix)
	}

	if !validatePrefixDeviceID(r.ID) {
		return NewInvalidRequestError(errors.New("device id must start with /devices/"),
			InvalidRequestDevicePrefix)
	}

	if !validatePrefixDeviceModel(r.DeviceModel) {
		return NewInvalidRequestError(errors.New("device model must start with /devicemodels/"),
			InvalidRequestDeviceModelPrefix)
	}

	return nil
}

func validatePrefixDeviceID(deviceID string) bool {
	return strings.HasPrefix(deviceID, "/devices/")
}

func validatePrefixDeviceModel(deviceModel string) bool {
	return strings.HasPrefix(deviceModel, "/devicemodels/")
}

// GetDeviceByIDRequest is the url param for the GetDeviceByID endpoint.
type GetDeviceByIDRequest struct {
	ID string `json:"id" validate:"required"`
}

func (r *GetDeviceByIDRequest) Bind(req *http.Request) error {
	id := chi.URLParam(req, "id")
	r.ID = id

	if err := validate.Struct(r); err != nil {
		return NewInvalidRequestError(err, RequiredDeviceID)
	}

	return nil
}

// DeviceResponse is the response body for the GetDeviceByID endpoint.
type DeviceResponse struct {
	ID          string `json:"id"`
	DeviceModel string `json:"deviceModel"` //nolint:tagliatelle
	Name        string `json:"name"`
	Note        string `json:"note"`
	Serial      string `json:"serial"`
}
