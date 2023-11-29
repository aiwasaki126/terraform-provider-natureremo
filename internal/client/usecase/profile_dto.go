package usecase

import "terraform-provider-natureremo/internal/client/domain/entity"

// ProfileDto is the data transfer object of the profile.
type ProfileDto struct {
	Id       string
	Nickname string
}

// newProfileDto is the constructor for ProfileDto from entity.
// This constructor is assumed to be used only in usevase layer.
func newProfileDto(p *entity.Profile) (*ProfileDto, error) {
	return &ProfileDto{
		Id:       p.User.GetId(),
		Nickname: p.User.GetNickname(),
	}, nil
}
