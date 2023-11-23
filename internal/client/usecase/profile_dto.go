package usecase

import "terraform-provider-natureremo/internal/client/domain/entity"

type ProfileDto struct {
	Id       string
	Nickname string
}

func newProfileDto(p *entity.Profile) (*ProfileDto, error) {
	return &ProfileDto{
		Id:       p.User.GetId(),
		Nickname: p.User.GetNickname(),
	}, nil
}
