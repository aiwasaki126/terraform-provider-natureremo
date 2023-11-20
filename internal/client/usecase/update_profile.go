package usecase

import (
	"context"
	"terraform-provider-natureremo/internal/client/domain/entity"
	"terraform-provider-natureremo/internal/client/domain/repository"
	"terraform-provider-natureremo/internal/client/domain/valueobject"
)

type UpdateProfile struct {
	repositoryProfile repository.Profile
}

func NewUpdateProfile(r repository.Profile) *UpdateProfile {
	return &UpdateProfile{
		repositoryProfile: r,
	}
}

func (u *UpdateProfile) Update(ctx context.Context, id, nickname, country, distanceUnit, tempUnit string) (*ProfileDto, error) {
	user, err := entity.ReconstructUser(id, nickname)
	if err != nil {
		return nil, err
	}
	preference, err := valueobject.NewPreference(country, distanceUnit, tempUnit)
	if err != nil {
		return nil, err
	}
	p, err := entity.ReconstructProfile(*user, *preference)
	if err != nil {
		return nil, err
	}
	profile, err := u.repositoryProfile.UpdateProfile(ctx, p)
	if err != nil {
		return nil, err
	}
	return newProfileDto(profile)
}
