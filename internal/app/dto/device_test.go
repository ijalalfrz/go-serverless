//go:build unit

package dto

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestCreateDeviceRequest_Bind(t *testing.T) {
	bindCreateDeviceRequest := func(name string, req CreateDeviceRequest, wantErr bool, errMsg string) func(t *testing.T) {
		return func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/api/devices", nil)
			err := req.Bind(request)

			if wantErr {
				assert.Error(t, err)
				if errMsg != "" {
					assert.Contains(t, err.Error(), errMsg)
				}
				return
			}
			assert.NoError(t, err)
		}
	}

	t.Run("success", bindCreateDeviceRequest(
		"success",
		CreateDeviceRequest{
			ID:          "/devices/123",
			DeviceModel: "/devicemodels/model-x",
			Name:        "Test Device",
			Note:        "Test Note",
			Serial:      "SN123",
		},
		false,
		"",
	))

	t.Run("invalid_device_id_prefix", bindCreateDeviceRequest(
		"invalid_device_id_prefix",
		CreateDeviceRequest{
			ID:          "123",
			DeviceModel: "/devicemodels/model-x",
			Name:        "Test Device",
			Note:        "Test Note",
			Serial:      "SN123",
		},
		true,
		"device id must start with /devices/",
	))

	t.Run("invalid_device_model_prefix", bindCreateDeviceRequest(
		"invalid_device_model_prefix",
		CreateDeviceRequest{
			ID:          "/devices/123",
			DeviceModel: "model-x",
			Name:        "Test Device",
			Note:        "Test Note",
			Serial:      "SN123",
		},
		true,
		"device model must start with /devicemodels/",
	))

	t.Run("missing_required_fields", bindCreateDeviceRequest(
		"missing_required_fields",
		CreateDeviceRequest{
			ID:          "/devices/123",
			DeviceModel: "/devicemodels/model-x",
		},
		true,
		"validation",
	))
}

func TestGetDeviceByIDRequest_Bind(t *testing.T) {
	bindGetDeviceByIDRequest := func(name string, urlParam string, wantID string, wantErr bool) func(t *testing.T) {
		return func(t *testing.T) {
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", urlParam)

			req := httptest.NewRequest(http.MethodGet, "/api/devices/"+urlParam, nil)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			var deviceReq GetDeviceByIDRequest
			err := deviceReq.Bind(req)

			if wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, wantID, deviceReq.ID)
		}
	}

	t.Run("success", bindGetDeviceByIDRequest(
		"success",
		"device-123",
		"device-123",
		false,
	))

	t.Run("empty_id", bindGetDeviceByIDRequest(
		"empty_id",
		"",
		"",
		true,
	))
}

func TestValidatePrefixDeviceID(t *testing.T) {
	validateDeviceID := func(name string, deviceID string, want bool) func(t *testing.T) {
		return func(t *testing.T) {
			got := validatePrefixDeviceID(deviceID)
			assert.Equal(t, want, got)
		}
	}

	t.Run("valid_prefix", validateDeviceID(
		"valid_prefix",
		"/devices/123",
		true,
	))

	t.Run("invalid_prefix", validateDeviceID(
		"invalid_prefix",
		"123",
		false,
	))

	t.Run("empty_string", validateDeviceID(
		"empty_string",
		"",
		false,
	))
}

func TestValidatePrefixDeviceModel(t *testing.T) {
	validateDeviceModel := func(name string, deviceModel string, want bool) func(t *testing.T) {
		return func(t *testing.T) {
			got := validatePrefixDeviceModel(deviceModel)
			assert.Equal(t, want, got)
		}
	}

	t.Run("valid_prefix", validateDeviceModel(
		"valid_prefix",
		"/devicemodels/model-x",
		true,
	))

	t.Run("invalid_prefix", validateDeviceModel(
		"invalid_prefix",
		"model-x",
		false,
	))

	t.Run("empty_string", validateDeviceModel(
		"empty_string",
		"",
		false,
	))
}
