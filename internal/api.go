package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/tabo-syu/parl/env"
)

type API struct {
	client   *http.Client
	password string
}

func NewAPI(password string) *API {
	return &API{http.DefaultClient, password}
}

type serverInfo struct {
	Version     string
	ServerName  string
	Description string
	WorldGUID   string
}

func (api *API) ServerInfo(ctx context.Context) (*serverInfo, error) {
	res, err := api.request(ctx, http.MethodGet, "/v1/api/info", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get server info: %w", err)
	}

	defer res.Body.Close()

	info, err := decode[serverInfo](res)
	if err != nil {
		return nil, fmt.Errorf("failed to decode server info: %w", err)
	}

	return info, nil
}

func (api *API) SaveWorld(ctx context.Context) error {
	res, err := api.request(ctx, http.MethodPost, "/v1/api/save", nil)
	if err != nil {
		return fmt.Errorf("failed to save world: %w", err)
	}

	defer res.Body.Close()

	return nil
}

func (api *API) ShutdownServer(ctx context.Context, waittime int, message string) error {
	res, err := api.request(ctx, http.MethodPost, "/v1/api/shutdown", map[string]any{
		"waittime": waittime,
		"message":  message,
	})
	if err != nil {
		return fmt.Errorf("failed to save world: %w", err)
	}

	defer res.Body.Close()

	return nil
}

func (api *API) request(ctx context.Context, method, path string, body any) (*http.Response, error) {
	var (
		reader *bytes.Reader
		bs     = []byte{}
	)

	if body != nil {
		var err error

		bs, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal json: %w", err)
		}
	}

	reader = bytes.NewReader(bs)

	u, err := url.JoinPath("http://"+env.Host+":"+env.Port, path)
	if err != nil {
		return nil, fmt.Errorf("failed to join a path: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, method, u, reader)
	if err != nil {
		return nil, fmt.Errorf("failed to build to request: %w", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth("admin", env.Password)

	res, err := api.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("faild to request: %w", err)
	}

	return res, nil
}

func decode[Res any](res *http.Response) (*Res, error) {
	response := json.NewDecoder(res.Body)

	if res.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("response code is over 400: %s", res.Status)
	}

	typed := new(Res)
	if err := response.Decode(typed); err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}

	return typed, nil
}
