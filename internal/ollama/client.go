package ollama

import (
	"net/http"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}

// define models
type Model struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Size        int64  `json:"size"`
	CreatedAt   string `json:"created_at"`
}
type ModelInfo struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Size        int64    `json:"size"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	Tags        []string `json:"tags"`
	Author      string   `json:"author"`
	License     string   `json:"license"`
}
type ModelStatus struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}
type Log struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Time    string `json:"time"`
}
type Metrics struct {
	Name   string  `json:"name"`
	CPU    float64 `json:"cpu"`
	Memory int64   `json:"memory"`
	Disk   int64   `json:"disk"`
}
type ModelRun struct {
	Name   string `json:"name"`
	Input  string `json:"input"`
	Output string `json:"output"`
}
type ModelRunResponse struct {
	Name   string `json:"name"`
	Output string `json:"output"`
	Error  string `json:"error"`
}
type ModelRunRequest struct {
	Name   string                 `json:"name"`
	Input  string                 `json:"input"`
	Params map[string]interface{} `json:"params"`
}
type ModelRunResult struct {
	Name   string `json:"name"`
	Output string `json:"output"`
	Error  string `json:"error"`
}
type ModelRunStatus struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Error  string `json:"error"`
}

type GetModelInfo struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Size        int64    `json:"size"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	Tags        []string `json:"tags"`
	Author      string   `json:"author"`
	License     string   `json:"license"`
}

// Add methods to interact with ollama running locally

// Add methods to interact with the Ollama API running on localhost
func (c *Client) GetModel(name string) (*Model, error) {
	// Implement the logic to get a model from the Ollama API
	// using the provided name.
	// Return the model and any error encountered.
	return nil, nil
}
func (c *Client) ListModels() ([]Model, error) {
	// Implement the logic to list all models from the Ollama API.
	// Return the list of models and any error encountered.
	return nil, nil
}
func (c *Client) RunModel(name string, input string) (string, error) {
	// Implement the logic to run a model from the Ollama API
	// using the provided name and input.
	// Return the output and any error encountered.
	return "", nil
}

func (c *Client) GetModelInfo(name string) (*ModelInfo, error) {
	// Implement the logic to get information about a model from the Ollama API
	// using the provided name.
	return nil, nil
}

func (c *Client) GetModelStatus(name string) (*ModelStatus, error) {
	// Implement the logic to get the status of a model from the Ollama API
	// using the provided name.
	// Return the model status and any error encountered.
	return nil, nil
}
func (c *Client) GetModelLogs(name string) ([]Log, error) {
	// Implement the logic to get logs for a model from the Ollama API
	// using the provided name.
	// Return the logs and any error encountered.
	return nil, nil
}
func (c *Client) GetModelMetrics(name string) (*Metrics, error) {
	// Implement the logic to get metrics for a model from the Ollama API
	// using the provided name.
	// Return the metrics and any error encountered.
	return nil, nil
}
