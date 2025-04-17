package test

import (
	"testing"

	"github.com/pingponglabs-backend/styley-golang-sdk/sdk"
	"github.com/pingponglabs-backend/styley-golang-sdk/sdk/deployments"
	"github.com/stretchr/testify/assert"
)

func TestGetJob(t *testing.T) {
	tests := []struct {
		name      string
		jobID     string
		expectErr bool
		apikey    string
	}{
		{
			name:      "Success API response",
			jobID:     "95f94401-0e02-47d6-9c7d-1b28418f11cc",
			expectErr: false,
			apikey:    "xxxx-xxx-xxxx-xxx", //Replace with valid API Key
		},
		{
			name:      "Error: Invalid Job ID",
			jobID:     "invalid-job-id",
			expectErr: true,
			apikey:    "xxxx-xxx-xxxx-xxx",
		},
		{
			name:      "Error: Unauthorized User",
			jobID:     "95f94401-0e02-47d6-9c7d-1b28418f11cc",
			expectErr: true,
			apikey:    "xxxx-xxx-xxxx-xxx",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			client := sdk.NewClient(sdk.WithKey(tt.apikey))

			job, err := client.Deployments().GetJob(tt.jobID)
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}

			if !tt.expectErr {
				assert.NotNil(t, job)
				assert.Equal(t, tt.jobID, job.Id)
			}
		})
	}
}
func TestCreateDeployment(t *testing.T) {
	tests := []struct {
		name          string
		deploymentReq deployments.CreateDeployment
		expectErr     bool
		apikey        string
	}{
		{
			name: "Success API response with valid deployment request",
			deploymentReq: deployments.CreateDeployment{
				Name: "Background Remover2",
				Args: map[string]interface{}{
					"input": "https://cdn.mediamagic.dev/media/c7dbd266-3aa3-11ed-8e27-e679ed67c206.jpeg",
				},
				ModelId:      "844218fa-c5d0-4cee-90ce-0b42d226ac8d",
				Sync:         true,
				OutputFormat: "png",
				OutputWidth:  480,
				OutputHeight: 480,
			},
			expectErr: false,
			apikey:    "xxxx-xxx-xxxx-xxx", //Replace with the valid API key
		},
		{
			name: "Error: Missing Model ID",
			deploymentReq: deployments.CreateDeployment{
				Name: "Background Remover2",
				Args: map[string]interface{}{
					"input": "https://cdn.mediamagic.dev/media/c7dbd266-3aa3-11ed-8e27-e679ed67c206.jpeg",
				},
				Sync: true,
			},
			expectErr: true,
			apikey:    "xxxx-xxx-xxxx-xxx", //Replace with valid API key
		},
		{
			name: "Error: Invalid input args",
			deploymentReq: deployments.CreateDeployment{
				Name: "Background Remover2",
				Args: map[string]interface{}{
					"test": "https://cdn.mediamagic.dev/media/c7dbd266-3aa3-11ed-8e27-e679ed67c206.jpeg",
				},
				ModelId: "844218fa-c5d0-4cee-90ce-0b42d226ac8d",
				Sync:    true,
			},
			expectErr: true,
			apikey:    "xxxx-xxx-xxxx-xxx", //Replace with valid API key
		},
		{
			name: "Error: Unauthorized user",
			deploymentReq: deployments.CreateDeployment{
				Name: "Background Remover2",
				Args: map[string]interface{}{
					"input": "https://cdn.mediamagic.dev/media/c7dbd266-3aa3-11ed-8e27-e679ed67c206.jpeg",
				},
				ModelId: "844218fa-c5d0-4cee-90ce-0b42d226ac8d",
				Sync:    true,
			},
			expectErr: true,
			apikey:    "xxxx-xxx-xxxx-xxx",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			client := sdk.NewClient(sdk.WithKey(tt.apikey))

			resp, err := client.Deployments().Create(tt.deploymentReq)

			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}
			if !tt.expectErr {
				assert.NotNil(t, resp)
				assert.NotEqual(t, len(resp.Job.Files), 0)
			}
		})
	}
}
func TestGetDeployment(t *testing.T) {
	tests := []struct {
		name      string
		expectErr bool
		apikey    string
	}{
		{
			name:      "Success API response",
			expectErr: false,
			apikey:    "xxxx-xxx-xxxx-xxx", //Replace with valid APIkey
		},
		{
			name:      "Error: Invalid API Key",
			expectErr: true,
			apikey:    "xxx-xxxx-xxx-xxxx",
		},
		{
			name:      "Error: Missing API Key",
			expectErr: true,
			apikey:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			client := sdk.NewClient(sdk.WithKey(tt.apikey))
			deployments, err := client.Deployments().List()
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}
			if !tt.expectErr {
				assert.NotNil(t, deployments)
			}
		})
	}
}
