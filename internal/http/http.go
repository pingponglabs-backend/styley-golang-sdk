package http

import (
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type Client struct {
	transport *http.Client
	apiKey    string
	url       string
}

func New(apiKey,url string) *Client {
	return &Client{
		transport: http.DefaultClient,
		apiKey:    apiKey,
		url:       url,
	}
}

func (c *Client) Get(path string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s",c.url, path), nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create GET request")
	}

	req.Header.Set("x-mediamagic-key", c.apiKey)

	response, err := c.transport.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make GET request")
	}
	return c.handleResponse(response)
}

func (c *Client) Post(path string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.url, path), body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create POST request")
	}

	req.Header.Set("x-mediamagic-key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	response, err := c.transport.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make POST request")
	}
	return c.handleResponse(response)
}

func (c *Client) handleResponse(response *http.Response) ([]byte, error) {
	defer response.Body.Close()

	if response.StatusCode > 399 {
		return nil, errors.Errorf("request failed: %s", response.Status)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	return responseBody, nil
}
