//go:build unit

package service

import (
	"context"
	"database/sql"
	"testing"

	"github.com/ijalalfrz/go-serverless/internal/app/dto"
	"github.com/stretchr/testify/assert"
)

func TestDeviceService_CreateDevice(t *testing.T) {
	createDeviceRequest := func(name string, req dto.CreateDeviceRequest, mockRepo *MockDeviceRepository, wantErr error) func(t *testing.T) {
		return func(t *testing.T) {
			svc := NewDeviceService(mockRepo)
			err := svc.CreateDevice(context.Background(), req)
			if wantErr != nil {
				assert.ErrorIs(t, err, wantErr)
				return
			}
			assert.NoError(t, err)
		}
	}

	t.Run("success", createDeviceRequest(
		"success",
		dto.CreateDeviceRequest{
			ID:          "device-3",
			DeviceModel: "Model Z",
			Name:        "Test Device 3",
			Note:        "Test Note 3",
			Serial:      "SN003",
		},
		&MockDeviceRepository{devices: mockDevices},
		nil,
	))

	t.Run("db_error", createDeviceRequest(
		"db_error",
		dto.CreateDeviceRequest{
			ID:          "device-3",
			DeviceModel: "Model Z",
			Name:        "Test Device 3",
			Note:        "Test Note 3",
			Serial:      "SN003",
		},
		&MockDeviceRepository{err: ErrMockDB},
		ErrMockDB,
	))
}

func TestDeviceService_GetDeviceByID(t *testing.T) {
	getDeviceByIDRequest := func(name string, req dto.GetDeviceByIDRequest, mockRepo *MockDeviceRepository, want dto.DeviceResponse, wantErr error) func(t *testing.T) {
		return func(t *testing.T) {
			svc := NewDeviceService(mockRepo)
			got, err := svc.GetDeviceByID(context.Background(), req)
			if wantErr != nil {
				assert.ErrorIs(t, err, wantErr)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, want, got)
		}
	}

	t.Run("success", getDeviceByIDRequest(
		"success",
		dto.GetDeviceByIDRequest{ID: "device-1"},
		&MockDeviceRepository{devices: mockDevices},
		dto.DeviceResponse{
			ID:          mockDevices[0].ID,
			DeviceModel: mockDevices[0].DeviceModel,
			Name:        mockDevices[0].Name,
			Note:        mockDevices[0].Note,
			Serial:      mockDevices[0].Serial,
		},
		nil,
	))

	t.Run("not_found", getDeviceByIDRequest(
		"not_found",
		dto.GetDeviceByIDRequest{ID: "non-existent"},
		&MockDeviceRepository{devices: mockDevices},
		dto.DeviceResponse{},
		sql.ErrNoRows,
	))

	t.Run("db_error", getDeviceByIDRequest(
		"db_error",
		dto.GetDeviceByIDRequest{ID: "device-1"},
		&MockDeviceRepository{err: ErrMockDB},
		dto.DeviceResponse{},
		ErrMockDB,
	))
}
