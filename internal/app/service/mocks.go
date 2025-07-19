//go:build unit

package service

import (
	"context"
	"errors"

	"github.com/ijalalfrz/go-serverless/internal/app/model"
	"github.com/ijalalfrz/go-serverless/internal/pkg/exception"
)

// Mock errors.
var (
	ErrMockDB      = errors.New("mock db error")
	ErrMockPublish = errors.New("mock publish error")
)

// MockDeviceRepository implements DeviceRepository interface.
type MockDeviceRepository struct {
	devices    []model.Device
	err        error
	getByIDErr error
	createErr  error
}

func (m *MockDeviceRepository) Create(ctx context.Context, device model.Device) error {
	if m.createErr != nil {
		return m.createErr
	}
	if m.err != nil {
		return m.err
	}
	m.devices = append(m.devices, device)
	return nil
}

func (m *MockDeviceRepository) GetByID(ctx context.Context, id string) (model.Device, error) {
	if m.getByIDErr != nil {
		return model.Device{}, m.getByIDErr
	}
	if m.err != nil {
		return model.Device{}, m.err
	}

	// Check if device exists in the mock data
	for _, device := range m.devices {
		if device.ID == id {
			return device, nil
		}
	}

	// Return the error that the service expects
	return model.Device{}, exception.ErrRecordNotFound
}

// Test data.
var mockDevices = []model.Device{
	{
		ID:          "device-1",
		DeviceModel: "Model X",
		Name:        "Test Device 1",
		Note:        "Test Note 1",
		Serial:      "SN001",
	},
	{
		ID:          "device-2",
		DeviceModel: "Model Y",
		Name:        "Test Device 2",
		Note:        "Test Note 2",
		Serial:      "SN002",
	},
}
