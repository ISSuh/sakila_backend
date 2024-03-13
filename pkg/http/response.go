package http

import "net/http"

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func NewResponseOK() Response {
	return Response{
		Status:  http.StatusOK,
		Message: "",
	}
}

func NewResponseOKWithBody(message string) Response {
	return Response{
		Status:  http.StatusOK,
		Message: message,
	}
}
