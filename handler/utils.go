package handler

type Response struct {
	Code    int           `json:"code"`
	Message string        `json:"message,omitempty"`
	Data    []interface{} `json:"data,omitempty"`
}