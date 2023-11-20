package infra

import (
	"reflect"
	"terraform-provider-natureremo/internal/client/domain/entity"
	"terraform-provider-natureremo/internal/client/domain/valueobject"
	"terraform-provider-natureremo/internal/client/infra/gen"
	"testing"
)

func TestProfileRepository_extractUpdateProfileParam(t *testing.T) {
	type fields struct {
		client *gen.Client
	}
	type args struct {
		nickname     string
		country      string
		distanceUnit string
		tempUnit     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *gen.UpdateProfileParam
		wantErr bool
	}{
		{
			"normal", fields{}, args{"jack", "JP", "metric", "c"},
			&gen.UpdateProfileParam{Nickname: toPtr("jack"), Country: toPtr(gen.JP), DistanceUnit: toPtr(gen.Metric), TempUnit: toPtr(gen.C)},
			false,
		},
		{
			"undefined_country", fields{}, args{"jack", "XX", "metric", "c"},
			nil,
			true,
		},
		{
			"undefined_distance_unit", fields{}, args{"jack", "JP", "feet", "c"},
			nil,
			true,
		},
		{
			"undefined_temp_unit", fields{}, args{"jack", "JP", "imperial", "u"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ProfileRepository{
				client: tt.fields.client,
			}
			u, _ := entity.ReconstructUser("hoge", tt.args.nickname)
			p, err := valueobject.NewPreference(tt.args.country, tt.args.distanceUnit, tt.args.tempUnit)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProfileRepository.extractUpdateProfileParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			profile, err := entity.ReconstructProfile(*u, *p)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProfileRepository.extractUpdateProfileParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got, err := r.extractUpdateProfileParam(profile)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProfileRepository.extractUpdateProfileParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProfileRepository.extractUpdateProfileParam() = %v, want %v", got, tt.want)
			}
		})
	}
}
