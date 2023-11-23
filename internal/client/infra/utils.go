package infra

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func toPtr[T any](v T) *T {
	return &v
}

func hasStatusOk(resp *http.Response) error {
	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	default:
		return fmt.Errorf("status=%v, %s", resp.StatusCode, resp.Status)
	}
}

func extractResponseBody[T any](resp *http.Response) (T, error) {
	var t T
	if err := hasStatusOk(resp); err != nil {
		return t, err
	}
	var respBody T
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return t, err
	}
	return respBody, nil
}
