package repository

import (
	"context"
	"terraform-provider-natureremo/internal/client/domain/entity"
)

// Profile provides interface of repository related to profile entity.
type Profile interface {
	GetProfile(ctx context.Context) (*entity.Profile, error)
	UpdateProfile(ctx context.Context, m *entity.Profile) (*entity.Profile, error)
}
