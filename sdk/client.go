package sdk

import (
	"os"

	"github.com/pingponglabs-backend/styley-golang-sdk/internal/http"
	"github.com/pingponglabs-backend/styley-golang-sdk/sdk/deployments"
	"github.com/pingponglabs-backend/styley-golang-sdk/sdk/models"
)

const (
	BasePath = "https://api-qa.mediamagic.ai"
)

type Client struct {
	models      *models.Client
	deployments *deployments.Client
}

type Options struct {
	Key string
	Url string
}

type Option func(*Options)

func WithKey(key string) Option {
	return func(o *Options) {
		o.Key = key
	}
}

func NewClient(opts ...Option) *Client {
	url := os.Getenv("MM_HOST_URL")
	if url == "" {
		url = BasePath
	}

	defaultOptions := Options{
		Key: os.Getenv("X_STYLEY_KEY"),
		Url: url,
	}

	for _, opt := range opts {
		opt(&defaultOptions)
	}

	// Default http transport
	httpTransport := http.New(defaultOptions.Key, defaultOptions.Url)

	// Clients
	modelsClient := models.NewClient(httpTransport)
	deploymentsClient := deployments.NewClient(httpTransport, modelsClient)

	return &Client{
		models:      modelsClient,
		deployments: deploymentsClient,
	}
}

func (c *Client) Models() *models.Client {
	return c.models
}

func (c *Client) Deployments() *deployments.Client {
	return c.deployments
}
