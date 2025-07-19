//go:build unit

package service

import (
	"context"
	"testing"

	"github.com/ijalalfrz/go-serverless/internal/app/dto"
	"github.com/ijalalfrz/go-serverless/internal/app/model"
	"github.com/ijalalfrz/go-serverless/internal/pkg/exception"
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

	t.Run("device_already_exists", createDeviceRequest(
		"device_already_exists",
		dto.CreateDeviceRequest{
			ID:          "device-1", // This ID already exists in mockDevices
			DeviceModel: "Model Z",
			Name:        "Test Device 3",
			Note:        "Test Note 3",
			Serial:      "SN003",
		},
		&MockDeviceRepository{devices: mockDevices},
		exception.ErrConflict,
	))

	t.Run("get_by_id_error", createDeviceRequest(
		"get_by_id_error",
		dto.CreateDeviceRequest{
			ID:          "device-3",
			DeviceModel: "Model Z",
			Name:        "Test Device 3",
			Note:        "Test Note 3",
			Serial:      "SN003",
		},
		&MockDeviceRepository{devices: mockDevices, getByIDErr: ErrMockDB},
		ErrMockDB,
	))

	t.Run("create_error_after_check", createDeviceRequest(
		"create_error_after_check",
		dto.CreateDeviceRequest{
			ID:          "device-3",
			DeviceModel: "Model Z",
			Name:        "Test Device 3",
			Note:        "Test Note 3",
			Serial:      "SN003",
		},
		&MockDeviceRepository{devices: mockDevices, createErr: ErrMockDB},
		ErrMockDB,
	))

	t.Run("empty_request_fields", createDeviceRequest(
		"empty_request_fields",
		dto.CreateDeviceRequest{
			ID:          "",
			DeviceModel: "",
			Name:        "",
			Note:        "",
			Serial:      "",
		},
		&MockDeviceRepository{devices: mockDevices},
		nil, // Should succeed since device doesn't exist
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
		exception.ErrRecordNotFound,
	))

	t.Run("db_error", getDeviceByIDRequest(
		"db_error",
		dto.GetDeviceByIDRequest{ID: "device-1"},
		&MockDeviceRepository{err: ErrMockDB},
		dto.DeviceResponse{},
		ErrMockDB,
	))

	t.Run("empty_id", getDeviceByIDRequest(
		"empty_id",
		dto.GetDeviceByIDRequest{ID: ""},
		&MockDeviceRepository{devices: mockDevices},
		dto.DeviceResponse{},
		exception.ErrRecordNotFound,
	))

	t.Run("special_characters_id", getDeviceByIDRequest(
		"special_characters_id",
		dto.GetDeviceByIDRequest{ID: "device-1@#$%"},
		&MockDeviceRepository{devices: mockDevices},
		dto.DeviceResponse{},
		exception.ErrRecordNotFound,
	))
}

func TestDeviceService_NewDeviceService(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := &MockDeviceRepository{}
		svc := NewDeviceService(mockRepo)

		assert.NotNil(t, svc)
		assert.Equal(t, mockRepo, svc.deviceRepo)
	})

	t.Run("nil_repository", func(t *testing.T) {
		svc := NewDeviceService(nil)

		assert.NotNil(t, svc)
		assert.Nil(t, svc.deviceRepo)
	})
}

func TestDeviceService_ErrorMessages(t *testing.T) {
	t.Run("create_device_error_message", func(t *testing.T) {
		mockRepo := &MockDeviceRepository{err: ErrMockDB}
		svc := NewDeviceService(mockRepo)

		req := dto.CreateDeviceRequest{
			ID:          "device-3",
			DeviceModel: "Model Z",
			Name:        "Test Device 3",
			Note:        "Test Note 3",
			Serial:      "SN003",
		}

		err := svc.CreateDevice(context.Background(), req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to get device")
	})

	t.Run("create_device_already_exists_message", func(t *testing.T) {
		mockRepo := &MockDeviceRepository{devices: mockDevices}
		svc := NewDeviceService(mockRepo)

		req := dto.CreateDeviceRequest{
			ID:          "device-1", // Already exists
			DeviceModel: "Model Z",
			Name:        "Test Device 3",
			Note:        "Test Note 3",
			Serial:      "SN003",
		}

		err := svc.CreateDevice(context.Background(), req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "record already exist")
	})

	t.Run("create_device_create_error_message", func(t *testing.T) {
		mockRepo := &MockDeviceRepository{devices: mockDevices, createErr: ErrMockDB}
		svc := NewDeviceService(mockRepo)

		req := dto.CreateDeviceRequest{
			ID:          "device-3",
			DeviceModel: "Model Z",
			Name:        "Test Device 3",
			Note:        "Test Note 3",
			Serial:      "SN003",
		}

		err := svc.CreateDevice(context.Background(), req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create device")
	})

	t.Run("get_device_error_message", func(t *testing.T) {
		mockRepo := &MockDeviceRepository{err: ErrMockDB}
		svc := NewDeviceService(mockRepo)

		req := dto.GetDeviceByIDRequest{ID: "device-1"}

		_, err := svc.GetDeviceByID(context.Background(), req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to get device")
	})
}

func TestDeviceService_EdgeCases(t *testing.T) {
	t.Run("create_device_with_maximum_fields", func(t *testing.T) {
		mockRepo := &MockDeviceRepository{devices: mockDevices}
		svc := NewDeviceService(mockRepo)

		req := dto.CreateDeviceRequest{
			ID:          "device-max-length-id-123456789012345678901234567890",
			DeviceModel: "Model With Very Long Name That Exceeds Normal Limits",
			Name:        "Device Name With Special Characters: @#$%^&*()_+-=[]{}|;':\",./<>?",
			Note:        "Note with multiple lines\nand special characters\tand spaces   ",
			Serial:      "SN-123456789-ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		}

		err := svc.CreateDevice(context.Background(), req)
		assert.NoError(t, err)
	})

	t.Run("get_device_with_special_characters", func(t *testing.T) {
		// Add a device with special characters to mock data
		specialDevice := model.Device{
			ID:          "device-special@#$%",
			DeviceModel: "Model@#$%",
			Name:        "Device@#$%",
			Note:        "Note@#$%",
			Serial:      "SN@#$%",
		}
		mockRepo := &MockDeviceRepository{devices: append(mockDevices, specialDevice)}
		svc := NewDeviceService(mockRepo)

		req := dto.GetDeviceByIDRequest{ID: "device-special@#$%"}

		result, err := svc.GetDeviceByID(context.Background(), req)
		assert.NoError(t, err)
		assert.Equal(t, specialDevice.ID, result.ID)
		assert.Equal(t, specialDevice.DeviceModel, result.DeviceModel)
		assert.Equal(t, specialDevice.Name, result.Name)
		assert.Equal(t, specialDevice.Note, result.Note)
		assert.Equal(t, specialDevice.Serial, result.Serial)
	})

	t.Run("create_device_duplicate_serial", func(t *testing.T) {
		// This test verifies that the service only checks for ID uniqueness
		// Serial uniqueness would be a business rule to implement if needed
		mockRepo := &MockDeviceRepository{devices: mockDevices}
		svc := NewDeviceService(mockRepo)

		req := dto.CreateDeviceRequest{
			ID:          "device-3", // New ID
			DeviceModel: "Model Z",
			Name:        "Test Device 3",
			Note:        "Test Note 3",
			Serial:      "SN001", // Same serial as device-1
		}

		err := svc.CreateDevice(context.Background(), req)
		assert.NoError(t, err) // Should succeed since only ID uniqueness is checked
	})
}

func TestDeviceService_ResponseMapping(t *testing.T) {
	t.Run("complete_device_response_mapping", func(t *testing.T) {
		// Create a device with all fields populated
		completeDevice := model.Device{
			ID:          "complete-device",
			DeviceModel: "Complete Model",
			Name:        "Complete Device",
			Note:        "Complete Note",
			Serial:      "COMPLETE-SN-001",
		}

		mockRepo := &MockDeviceRepository{devices: []model.Device{completeDevice}}
		svc := NewDeviceService(mockRepo)

		req := dto.GetDeviceByIDRequest{ID: "complete-device"}

		result, err := svc.GetDeviceByID(context.Background(), req)
		assert.NoError(t, err)

		// Verify all fields are correctly mapped
		assert.Equal(t, completeDevice.ID, result.ID)
		assert.Equal(t, completeDevice.DeviceModel, result.DeviceModel)
		assert.Equal(t, completeDevice.Name, result.Name)
		assert.Equal(t, completeDevice.Note, result.Note)
		assert.Equal(t, completeDevice.Serial, result.Serial)
	})

	t.Run("partial_device_response_mapping", func(t *testing.T) {
		// Create a device with some empty fields
		partialDevice := model.Device{
			ID:          "partial-device",
			DeviceModel: "",
			Name:        "Partial Device",
			Note:        "",
			Serial:      "PARTIAL-SN-001",
		}

		mockRepo := &MockDeviceRepository{devices: []model.Device{partialDevice}}
		svc := NewDeviceService(mockRepo)

		req := dto.GetDeviceByIDRequest{ID: "partial-device"}

		result, err := svc.GetDeviceByID(context.Background(), req)
		assert.NoError(t, err)

		// Verify fields are correctly mapped, including empty ones
		assert.Equal(t, partialDevice.ID, result.ID)
		assert.Equal(t, partialDevice.DeviceModel, result.DeviceModel)
		assert.Equal(t, partialDevice.Name, result.Name)
		assert.Equal(t, partialDevice.Note, result.Note)
		assert.Equal(t, partialDevice.Serial, result.Serial)
	})
}
