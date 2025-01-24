package models

import (
	"strings"
	"testing"

	"git.mediamagic.ai/styley-sdks/styley-golang-sdk/internal/http"
)

const (
	BasePath = "https://api-qa.mediamagic.ai"
)

func TestGetModelByName(t *testing.T) {
	tests := []struct {
		name        string
		apikey      string
		modelName   string
		expectErr   bool
		expectModel string // The expected model name
	}{
		{
			name:        "Success: Valid API Key and Model Name",
			apikey:      "xxx-xxxx-xxx-xxxx", // Replace with valid API key
			modelName:   "/Omni Zero",
			expectErr:   false,
			expectModel: "/Omni Zero", // Expected model name
		},
		{
			name:        "Error: Missing API Key",
			apikey:      "", // No API key
			modelName:   "/Omni Zero",
			expectErr:   true,
			expectModel: "",
		},
		{
			name:        "Error: Invalid API Key",
			apikey:      "xxx-xxxx-xxx-xxxx", // Invalid API key
			modelName:   "/Omni Zero",
			expectErr:   true,
			expectModel: "",
		},
		{
			name:        "Error: Model Not Found",
			apikey:      "xxx-xxxx-xxx-xxxx", //Replace with valid API key
			modelName:   "/NonExistentModel",// Model that doesn't exist
			expectErr:   true,
			expectModel: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(http.New(tt.apikey, BasePath))

			model, err := client.GetByName(tt.modelName)

			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}
			expectedModelName := strings.TrimPrefix(tt.expectModel, "/")
			if !tt.expectErr && model.Name != expectedModelName {
				t.Errorf("expected model name: %s, got: %s", tt.expectModel, model.Name)
			}
		})
	}
}

func TestGetModelsById(t *testing.T) {
	tests := []struct {
		name      string
		modelID   string
		apikey    string
		expectErr bool
	}{
		{
			name:      "Success API response",
			modelID:   "844218fa-c5d0-4cee-90ce-0b42d226ac8d",
			expectErr: false,
			apikey:    "xxx-xxxx-xxx-xxxx",//Replace with the valid APIkey
		},
		{
			name:      "Error: Invalid Model ID",
			modelID:   "invalid-id",
			expectErr: true,
			apikey:    "xxx-xxxx-xxx-xxxx",//Replace with the valid APIkey
		},
		{
			name:      "Error: UnAuthorized",
			modelID:   "844218fa-c5d0-4cee-90ce-0b42d226ac8d",
			expectErr: true,
			apikey:    "9b7d9a38-5fad-11ef-ac07-30d042e69440",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			client := NewClient(http.New(tt.apikey, BasePath))

			model, err := client.GetByID(tt.modelID)
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}
			if !tt.expectErr && model.ID != tt.modelID {
				t.Errorf("expected model ID: %s, got: %s", tt.modelID, model.ID)
			}
		})
	}
}

func TestModelsList(t *testing.T) {
	tests := []struct {
		name         string
		expectErr    bool
		apikey       string
		expectModels int
	}{
		{
			name:         "sucess API response",
			apikey:       "xxx-xxxx-xxx-xxxx",//Replace with the valid API key
			expectErr:    false,
			expectModels: 243,
		},
		{
			name:         "Error: Missing API Key",
			expectErr:    true,
			expectModels: 0,
		},
		{
			name:         "Error: Invalid API Key",
			apikey:       "xxx-xxxx-xxx-xxxx",
			expectErr:    true,
			expectModels: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			client := NewClient(http.New(tt.apikey, BasePath))

			models, err := client.List()
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}

			if len(models) < tt.expectModels {
				t.Errorf("expected %d models, got: %d", tt.expectModels, len(models))
			}

		})
	}

}
