package types

type APIResponse struct {
	StatusCode    int         `json:"statusCode"`
	StatusMessage string      `json:"statusMessage"`
	Data          interface{} `json:"data,omitempty"`
}
