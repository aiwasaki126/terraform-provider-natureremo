package repository

import (
	"context"
	"terraform-provider-natureremo/internal/client/domain/entity"
)

// Device provides interface of repository related to device entity.
type Device interface {
	GetAllDevices(ctx context.Context) ([]*entity.Device, error)
	GetDevice(ctx context.Context, id string) (*entity.Device, error)
	UpdateDevice(ctx context.Context, d *entity.Device) (*entity.Device, error)
	DeleteDevice(ctx context.Context, d *entity.Device) error
}
