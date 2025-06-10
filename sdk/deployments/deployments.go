package deployments

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/pingponglabs-backend/styley-golang-sdk/internal/http"
	"github.com/pingponglabs-backend/styley-golang-sdk/sdk/models"
	"github.com/pkg/errors"
)

const (
	COMPLETE = "complete"
	FAILED   = "failed"
	ETA      = 180
)

type Client struct {
	httpClient   *http.Client
	modelsClient *models.Client
}

func NewClient(httpClient *http.Client, modelsClient *models.Client) *Client {
	return &Client{
		httpClient:   httpClient,
		modelsClient: modelsClient,
	}
}

func (c *Client) List() (*DeploymentsResponse, error) {
	response, err := c.httpClient.Get("/api/v1/deployments")
	if err != nil {
		return nil, errors.Wrap(err, "failed to list deployments")
	}

	var deployments *DeploymentsResponse
	if err := json.Unmarshal(response, &deployments); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal deployments")
	}

	return deployments, nil
}

func (c *Client) Create(deployment CreateDeployment) (*Deployment, error) {
	model, err := c.modelsClient.GetByID(deployment.ModelId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get model by ModelId: "+deployment.ModelId)
	}

	deployment.ModelId = model.ID

	body, err := json.Marshal(deployment)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal deployment body")
	}
	response, err := c.httpClient.Post("/api/v1/deployments", bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create deployment")
	}

	var deploymentResponse *Deployment
	if err := json.Unmarshal(response, &deploymentResponse); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal deployment response")
	}
	status := deploymentResponse.Job.Status
	eta := deploymentResponse.Job.Eta
	if eta == 0 {
		eta = ETA
	}
	if deployment.Synchronous {
		for status != COMPLETE && status != FAILED && eta > 0 {
			time.Sleep(10 * time.Second)
			job, err := c.GetJob(deploymentResponse.Job.Id)
			if err != nil {
				return nil, errors.Wrap(err, "failed to check deployment status")
			}
			status = job.Status
			deploymentResponse.Job = *job
			deploymentResponse.Status = job.Status
			deploymentResponse.Logs = job.Logs
			eta -= 5
		}
	}

	return deploymentResponse, nil
}

func (c *Client) GetJob(jobID string) (*Job, error) {
	response, err := c.httpClient.Get("/api/v1/jobs/" + jobID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Job by JobID: "+jobID)
	}
	var jobResponse *Job
	if err := json.Unmarshal(response, &jobResponse); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal GetJob response")
	}

	return jobResponse, nil
}
