package service

import (
	"context"
	"fmt"

	"github.com/ijalalfrz/go-serverless/internal/app/dto"
	"github.com/ijalalfrz/go-serverless/internal/app/model"
)

type DeviceRepository interface {
	Create(ctx context.Context, device model.Device) error
	GetByID(ctx context.Context, id string) (model.Device, error)
}

type DeviceService struct {
	deviceRepo DeviceRepository
}

func NewDeviceService(deviceRepo DeviceRepository) *DeviceService {
	return &DeviceService{deviceRepo: deviceRepo}
}

func (s *DeviceService) CreateDevice(ctx context.Context, req dto.CreateDeviceRequest) error {
	device := model.Device{
		ID:          req.ID,
		DeviceModel: req.DeviceModel,
		Name:        req.Name,
		Note:        req.Note,
		Serial:      req.Serial,
	}

	existingDevice, err := s.deviceRepo.GetByID(ctx, req.ID)
	if err != nil {
		return fmt.Errorf("failed to get device: %w", err)
	}

	if existingDevice.ID != "" {
		return fmt.Errorf("device already exists")
	}

	if err := s.deviceRepo.Create(ctx, device); err != nil {
		return fmt.Errorf("failed to create device: %w", err)
	}

	return nil
}

func (s *DeviceService) GetDeviceByID(ctx context.Context, req dto.GetDeviceByIDRequest) (dto.DeviceResponse, error) {
	device, err := s.deviceRepo.GetByID(ctx, req.ID)
	if err != nil {
		return dto.DeviceResponse{}, fmt.Errorf("failed to get device: %w", err)
	}

	return dto.DeviceResponse{
		ID:          device.ID,
		DeviceModel: device.DeviceModel,
		Name:        device.Name,
		Note:        device.Note,
		Serial:      device.Serial,
	}, nil
}
