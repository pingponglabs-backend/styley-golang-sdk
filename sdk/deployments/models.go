package deployments

type CreateDeployment struct {
	Name         string                 `json:"name"`
	Args         map[string]interface{} `json:"args"`
	ModelId      string                 `json:"model_id"`
	Synchronous  bool                   `json:"synchronous"`
	OutputFormat string                 `json:"output_format,omitempty"`
	OutputWidth  int                    `json:"output_width,omitempty"`
	OutputHeight int                    `json:"output_height,omitempty"`
}

type Job struct {
	Id           string   `json:"id"`
	Files        []string `json:"files"`
	CreditsUsed  float32  `json:"credits_used"`
	Status       string   `json:"status"`
	Error        string   `json:"error"`
	DeploymentId string   `json:"deployment_id"`
	Duration     int64    `json:"duration"`
	Logs         string   `json:"logs"`
	Eta          int64    `json:"eta"`
	Etr          int64    `json:"etr"`
}

type Deployment struct {
	ID      string `json:"id"`
	Job     Job    `json:"job"`
	Status  string `json:"status"`
	Name    string `json:"name"`
	ModelId string `json:"model_id"`
	JobId   string `json:"job_id"`
	Logs    string `json:"logs"`
}

type DeploymentsResponse struct {
	Deployments []Deployment `json:"deployments"`
	Total       int          `json:"total"`
}
