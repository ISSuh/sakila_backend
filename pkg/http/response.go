package http

import "net/http"

type AppResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func NewRespons() *AppResponse {
	return &AppResponse{}
}

func NewResponseOK() *AppResponse {
	return &AppResponse{
		Status:  http.StatusOK,
		Message: "",
	}
}

func NewResponseOKWithBody(message string) *AppResponse {
	return &AppResponse{
		Status:  http.StatusOK,
		Message: message,
	}
}

func NewResponseErrWithBody(status int, message string) *AppResponse {
	return &AppResponse{
		Status:  status,
		Message: message,
	}
}
