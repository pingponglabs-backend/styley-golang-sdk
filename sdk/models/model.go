package models

type Model struct {
	ID            string                 `json:"id"`
	Name          string                 `json:"name"`
	Args          map[string]interface{} `json:"args"`
	Alias         string                 `json:"alias"`
	ImageName     string                 `json:"image_name,omitempty"`
	InferenceCost float32                `json:"inference_cost,omitempty"`
	IsAsync       bool                   `json:"is_async,omitempty"`
}
