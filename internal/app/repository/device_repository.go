package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ijalalfrz/go-serverless/internal/app/model"
	"github.com/ijalalfrz/go-serverless/internal/pkg/exception"
)

type DeviceRepository struct {
	db        *dynamodb.Client
	tableName string
}

func NewDeviceRepository(db *dynamodb.Client, tableName string) *DeviceRepository {
	return &DeviceRepository{
		db:        db,
		tableName: tableName,
	}
}

func (r *DeviceRepository) Create(ctx context.Context, device model.Device) error {
	id := normalizeID(device.ID)

	device.PK = "DEVICE#" + id

	data, err := attributevalue.MarshalMap(device)
	if err != nil {
		return fmt.Errorf("failed to marshal device: %w", err)
	}

	_, err = r.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &r.tableName,
		Item:      data,
	})
	if err != nil {
		return fmt.Errorf("failed to create device: %w", err)
	}

	return nil
}

func (r *DeviceRepository) GetByID(ctx context.Context, id string) (model.Device, error) {
	nid := normalizeID(id)

	out, err := r.db.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &r.tableName,
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: "DEVICE#" + nid},
		},
	})
	if err != nil {
		return model.Device{}, fmt.Errorf("failed to get device: %w", err)
	}

	if out.Item == nil {
		err := exception.ErrRecordNotFound
		err.MessageVars = map[string]interface{}{
			"name": "device",
		}

		return model.Device{}, err
	}

	device := model.Device{}

	err = attributevalue.UnmarshalMap(out.Item, &device)
	if err != nil {
		return model.Device{}, fmt.Errorf("failed to unmarshal device: %w", err)
	}

	return device, nil
}

func normalizeID(id string) string {
	return strings.TrimPrefix(id, "/devices/")
}
