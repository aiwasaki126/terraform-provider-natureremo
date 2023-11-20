package infra

import (
	"context"
	"fmt"
	"terraform-provider-natureremo/internal/client/domain/entity"
	"terraform-provider-natureremo/internal/client/domain/valueobject"
	"terraform-provider-natureremo/internal/client/infra/gen"

	"github.com/go-playground/validator/v10"
)

type DeviceRepository struct {
	client   *gen.Client
	validate *validator.Validate
}

func NewDeviceRepository(c *gen.Client) *DeviceRepository {
	return &DeviceRepository{client: c, validate: validator.New()}
}

func (r *DeviceRepository) GetAllDevices(ctx context.Context) ([]*entity.Device, error) {
	resp, err := r.client.Get1Devices(ctx)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := extractResponseBody[gen.DeviceResponses](resp)
	if err != nil {
		return nil, err
	}
	for _, dr := range respBody {
		if err := r.validate.Struct(dr); err != nil {
			return nil, err
		}
	}

	devices := make([]*entity.Device, 0, len(respBody))
	for _, d := range respBody {
		device, err := r.buildDeviceEntityFromResponse(&d)
		if err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}
	return devices, nil
}

func (r *DeviceRepository) GetDevice(ctx context.Context, id string) (*entity.Device, error) {
	devices, err := r.GetAllDevices(ctx)
	if err != nil {
		return nil, err
	}
	for _, d := range devices {
		if d.GetId() == id {
			return d, nil
		}
	}
	return nil, fmt.Errorf("not found device id=%s", id)
}

func (r *DeviceRepository) UpdateDevice(ctx context.Context, d *entity.Device) (*entity.Device, error) {
	var err error
	_, err = r.updateName(ctx, d)
	if err != nil {
		return nil, err
	}
	_, err = r.updateHumidityOffset(ctx, d)
	if err != nil {
		return nil, err
	}
	device, err := r.updateTemperatureOffset(ctx, d)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (r *DeviceRepository) updateName(ctx context.Context, d *entity.Device) (*entity.Device, error) {
	updateParam := gen.UpdateDeviceParameters{Name: d.GetName()}
	if err := r.validate.Struct(updateParam); err != nil {
		return nil, err
	}
	resp, err := r.client.Post1DevicesDeviceidWithFormdataBody(ctx, d.GetId(), updateParam)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := extractResponseBody[gen.DeviceResponse](resp)
	if err != nil {
		return nil, err
	}
	if err := r.validate.Struct(respBody); err != nil {
		return nil, err
	}
	return r.buildDeviceEntityFromResponse(&respBody)
}

func (r *DeviceRepository) updateHumidityOffset(ctx context.Context, d *entity.Device) (*entity.Device, error) {
	offsetParam := gen.HumidityOffsetParams{Offset: toPtr(float32(d.GetHumidityOffset()))}
	if err := r.validate.Struct(offsetParam); err != nil {
		return nil, err
	}
	resp, err := r.client.Post1DevicesDeviceidHumidityOffsetWithFormdataBody(ctx, d.GetId(), offsetParam)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := extractResponseBody[gen.DeviceResponse](resp)
	if err != nil {
		return nil, err
	}
	if err := r.validate.Struct(respBody); err != nil {
		return nil, err
	}
	return r.buildDeviceEntityFromResponse(&respBody)
}

func (r *DeviceRepository) updateTemperatureOffset(ctx context.Context, d *entity.Device) (*entity.Device, error) {
	offsetParam := gen.TemperatureOffsetParams{Offset: toPtr(float32(d.GetTemperatureOffset()))}
	if err := r.validate.Struct(offsetParam); err != nil {
		return nil, err
	}
	resp, err := r.client.Post1DevicesDeviceidTemperatureOffsetWithFormdataBody(ctx, d.GetId(), offsetParam)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := extractResponseBody[gen.DeviceResponse](resp)
	if err != nil {
		return nil, err
	}
	if err := r.validate.Struct(respBody); err != nil {
		return nil, err
	}
	return r.buildDeviceEntityFromResponse(&respBody)
}

func (r *DeviceRepository) DeleteDevice(ctx context.Context, d *entity.Device) error {
	resp, err := r.client.Post1DevicesDeviceidDeleteWithFormdataBody(ctx, d.GetId(), gen.Post1DevicesDeviceidDeleteFormdataRequestBody{})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := hasStatusOk(resp); err != nil {
		return err
	}
	return nil
}

func (r *DeviceRepository) buildDeviceEntityFromResponse(d *gen.DeviceResponse) (*entity.Device, error) {
	if err := r.validate.Struct(d); err != nil {
		return nil, err
	}
	setting, err := valueobject.NewDeviceSetting(d.Name, int64(d.HumidityOffset), float64(d.TemperatureOffset))
	if err != nil {
		return nil, err
	}
	spec := valueobject.NewDeviceSpec(d.SerialNumber, d.BtMacAddress, d.FirmwareVersion, d.MacAddress)

	users := make([]*entity.User, 0, len(d.Users))
	for _, u := range d.Users {
		user, err := entity.ReconstructUser(u.Id, u.Nickname)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return entity.ReconstructDevice(d.Id, *setting, *spec, users), nil
}
