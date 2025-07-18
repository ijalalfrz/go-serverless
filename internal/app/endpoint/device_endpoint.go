package endpoint

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	"github.com/ijalalfrz/go-serverless/internal/app/dto"
)

type DeviceService interface {
	CreateDevice(ctx context.Context, req dto.CreateDeviceRequest) error
	GetDeviceByID(ctx context.Context, req dto.GetDeviceByIDRequest) (dto.DeviceResponse, error)
}

func NewDeviceEndpoint(deviceService DeviceService) Device {
	return Device{
		CreateDevice:  makeCreateDeviceEndpoint(deviceService),
		GetDeviceByID: makeGetDeviceByIDEndpoint(deviceService),
	}
}

func makeCreateDeviceEndpoint(deviceService DeviceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*dto.CreateDeviceRequest)
		if !ok {
			return nil, fmt.Errorf("invalid request type: %w", ErrInvalidType)
		}

		err := deviceService.CreateDevice(ctx, *req)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

func makeGetDeviceByIDEndpoint(deviceService DeviceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*dto.GetDeviceByIDRequest)
		if !ok {
			return nil, fmt.Errorf("invalid request type: %w", ErrInvalidType)
		}

		device, err := deviceService.GetDeviceByID(ctx, *req)
		if err != nil {
			return nil, err
		}

		return device, nil
	}
}
