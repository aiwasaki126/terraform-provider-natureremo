package infra

import (
	"context"
	"fmt"
	"terraform-provider-natureremo/internal/client/domain/entity"
	"terraform-provider-natureremo/internal/client/domain/repository"
	"terraform-provider-natureremo/internal/client/domain/valueobject"
	"terraform-provider-natureremo/internal/client/infra/gen"

	"github.com/go-playground/validator/v10"
)

var (
	_ repository.Profile = &ProfileRepository{}
)

// ProfileRepository is the repository that handles profile entity.
type ProfileRepository struct {
	client   *gen.Client
	validate *validator.Validate
}

func NewProfileRepository(c *gen.Client) *ProfileRepository {
	return &ProfileRepository{client: c, validate: validator.New()}
}

func (r *ProfileRepository) GetProfile(ctx context.Context) (*entity.Profile, error) {
	resp, err := r.client.Get1UsersMe(ctx)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := extractResponseBody[gen.UserResponse](resp)
	if err != nil {
		return nil, err
	}
	if err := r.validate.Struct(respBody); err != nil {
		return nil, err
	}
	user, err := entity.ReconstructUser(respBody.Id, respBody.Nickname)
	if err != nil {
		return nil, err
	}
	return entity.ReconstructProfile(*user, valueobject.Preference{})
}

func (r *ProfileRepository) UpdateProfile(ctx context.Context, p *entity.Profile) (*entity.Profile, error) {
	param, err := r.extractUpdateProfileParam(p)
	if err != nil {
		return nil, err
	}
	resp, err := r.client.Post1UsersMeWithFormdataBody(ctx, *param)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := extractResponseBody[gen.UserResponse](resp)
	if err != nil {
		return nil, err
	}
	if err := r.validate.Struct(respBody); err != nil {
		return nil, err
	}
	user, err := entity.ReconstructUser(respBody.Id, respBody.Nickname)
	if err != nil {
		return nil, err
	}
	return entity.ReconstructProfile(*user, p.GetPreference())
}

func (r *ProfileRepository) extractUpdateProfileParam(p *entity.Profile) (*gen.UpdateProfileParam, error) {
	param := new(gen.UpdateProfileParam)

	if err := setNickname(param, p); err != nil {
		return nil, err
	}
	if err := setCountryUpdateParam(param, p); err != nil {
		return nil, err
	}
	if err := setDistanceUnitParam(param, p); err != nil {
		return nil, err
	}
	if err := setTempUnit(param, p); err != nil {
		return nil, err
	}
	return param, nil
}

func setNickname(param *gen.UpdateProfileParam, p *entity.Profile) error {
	switch p.GetNickname() {
	case "":
		return fmt.Errorf("invalid empty nickname")
	default:
		param.Nickname = toPtr(p.GetNickname())
	}
	return nil
}

func setCountryUpdateParam(param *gen.UpdateProfileParam, p *entity.Profile) error {
	switch p.GetCountry() {
	case "":
		break
	case string(gen.JP):
		param.Country = toPtr(gen.JP)
	case string(gen.US):
		param.Country = toPtr(gen.US)
	case string(gen.CA):
		param.Country = toPtr(gen.CA)
	case string(gen.SG):
		param.Country = toPtr(gen.SG)
	case string(gen.AU):
		param.Country = toPtr(gen.AU)
	case string(gen.NZ):
		param.Country = toPtr(gen.NZ)
	case string(gen.OTHERS):
		param.Country = toPtr(gen.OTHERS)
	default:
		return fmt.Errorf("invalid country %s in profile update params", p.GetCountry())
	}
	return nil
}

func setDistanceUnitParam(param *gen.UpdateProfileParam, p *entity.Profile) error {
	switch p.GetDistanceUnit() {
	case "":
		break
	case string(gen.Metric):
		param.DistanceUnit = toPtr(gen.Metric)
	case string(gen.Imperial):
		param.DistanceUnit = toPtr(gen.Imperial)
	default:
		return fmt.Errorf("invalid distance_unit %s in profile update params", p.GetDistanceUnit())
	}
	return nil
}

func setTempUnit(param *gen.UpdateProfileParam, p *entity.Profile) error {
	switch p.GetTempUnit() {
	case "":
		break
	case string(gen.C):
		param.TempUnit = toPtr(gen.C)
	case string(gen.F):
		param.TempUnit = toPtr(gen.F)
	default:
		return fmt.Errorf("invalid temp_unit %s in profile update params", p.GetTempUnit())
	}
	return nil
}
