package apiclient

import (
	"context"
	"fmt"
	"net/http"
	"terraform-provider-natureremo/internal/client/infra"
	"terraform-provider-natureremo/internal/client/infra/gen"
	"terraform-provider-natureremo/internal/client/usecase"
)

type Client struct {
	accessToken string
	gen.Client
}

func New(accessToken string) (*Client, error) {
	client := &Client{
		accessToken: accessToken,
	}
	c, err := gen.NewClient("https://api.nature.global", client.optAuthorizaionHeader)
	if err != nil {
		return nil, err
	}
	client.Client = *c
	return client, nil
}

func (c *Client) addAuthorizationHeader(ctx context.Context, req *http.Request) error {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	return nil
}

func (c *Client) optAuthorizaionHeader(client *gen.Client) error {
	client.RequestEditors = append(client.RequestEditors, c.addAuthorizationHeader)
	return nil
}

func (c *Client) GetProfile(ctx context.Context) (*usecase.ProfileDto, error) {
	u := usecase.NewGetProfile(infra.NewProfileRepository(&c.Client))
	profile, err := u.Get(ctx)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (c *Client) UpdateProfile(ctx context.Context, id, nickname string, country, distanceUnit, tempUnit string) (*usecase.ProfileDto, error) {
	u := usecase.NewUpdateProfile(infra.NewProfileRepository(&c.Client))
	profile, err := u.Update(ctx, id, nickname, country, distanceUnit, tempUnit)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (c *Client) GetAllDevices(ctx context.Context) ([]*usecase.DeviceDto, error) {
	u := usecase.NewGetAllDevices(infra.NewDeviceRepository(&c.Client))
	devices, err := u.GetAllDevices(ctx)
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func (c *Client) GetDevice(ctx context.Context, id string) (*usecase.DeviceDto, error) {
	u := usecase.NewGetDevice(infra.NewDeviceRepository(&c.Client))
	device, err := u.GetDevice(ctx, id)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (c *Client) UpdateDevice(ctx context.Context, id string, name string, humidityOffset int64, temperatureOffset float64) (*usecase.DeviceDto, error) {
	u := usecase.NewUpdateDevice(infra.NewDeviceRepository(&c.Client))
	deviceDto, err := usecase.NewDeviceDto(id, name, humidityOffset, temperatureOffset)
	if err != nil {
		return nil, err
	}
	deviceDto, err = u.UpdateDevice(ctx, *deviceDto)
	if err != nil {
		return nil, err
	}
	return deviceDto, nil
}

func (c *Client) DeleteDevice(ctx context.Context, id string) error {
	u := usecase.NewDeleteDevice(infra.NewDeviceRepository(&c.Client))
	return u.DeleteDevice(ctx, id)
}
