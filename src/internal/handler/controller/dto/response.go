package dto

type Response struct {
	StatusCode    int         `json:"status_code,omitempty"`
	StatusMessage string      `json:"status_message,omitempty"`
	Data          interface{} `json:"data,omitempty"`
	Error         interface{} `json:"error,omitempty"`
}
