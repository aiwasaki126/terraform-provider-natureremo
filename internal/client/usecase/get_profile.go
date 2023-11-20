package usecase

import (
	"context"
	"terraform-provider-natureremo/internal/client/domain/repository"
)

type GetProfile struct {
	repositoryProfile repository.Profile
}

func NewGetProfile(r repository.Profile) *GetProfile {
	return &GetProfile{
		repositoryProfile: r,
	}
}

func (u *GetProfile) Get(ctx context.Context) (*ProfileDto, error) {
	profile, err := u.repositoryProfile.GetProfile(ctx)
	if err != nil {
		return nil, err
	}
	return newProfileDto(profile)
}
