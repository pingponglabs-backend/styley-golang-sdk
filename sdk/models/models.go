package models

import (
	"encoding/json"
	"fmt"

	"github.com/pingponglabs-backend/styley-golang-sdk/internal/http"
)

type Client struct {
	httpClient *http.Client
}

func NewClient(httpTransport *http.Client) *Client {
	return &Client{
		httpClient: httpTransport,
	}
}

func (c *Client) List() ([]Model, error) {
	response, err := c.httpClient.Get("/api/v1/models")
	if err != nil {
		return nil, err
	}

	var models []Model
	if err := json.Unmarshal(response, &models); err != nil {
		return nil, err
	}

	return models, nil
}

func (c *Client) GetByID(id string) (*Model, error) {
	response, err := c.httpClient.Get(fmt.Sprintf("/api/v1/models/%v", id))
	if err != nil {
		return nil, err
	}

	var model Model
	if err := json.Unmarshal(response, &model); err != nil {
		return nil, err
	}

	return &model, nil
}

func (c *Client) GetByName(name string) (*Model, error) {
	path := generateNamePath(name)
	response, err := c.httpClient.Get(path)
	if err != nil {
		return nil, err
	}

	var model Model
	if err := json.Unmarshal(response, &model); err != nil {
		return nil, err
	}

	return &model, nil
}
