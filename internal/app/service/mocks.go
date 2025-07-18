package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ijalalfrz/go-serverless/internal/app/model"
)

// Mock errors
var (
	ErrMockDB      = errors.New("mock db error")
	ErrMockPublish = errors.New("mock publish error")
)

// MockDeviceRepository implements DeviceRepository interface
type MockDeviceRepository struct {
	devices []model.Device
	err     error
}

func (m *MockDeviceRepository) Create(ctx context.Context, device model.Device) error {
	if m.err != nil {
		return m.err
	}
	m.devices = append(m.devices, device)
	return nil
}

func (m *MockDeviceRepository) GetByID(ctx context.Context, id string) (model.Device, error) {
	if m.err != nil {
		return model.Device{}, m.err
	}
	for _, device := range m.devices {
		if device.ID == id {
			return device, nil
		}
	}
	return model.Device{}, sql.ErrNoRows
}

// Test data
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
