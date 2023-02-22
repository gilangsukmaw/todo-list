package server

type Response struct {
	Code    int `json:"code,omitempty"`
	Errors  any `json:"errors,omitempty"`
	Data    any `json:"data,omitempty"`
	Message any `json:"message,omitempty"`
}
