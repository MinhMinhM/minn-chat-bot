package test

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type localResponser struct {
}

type responseBody struct {
	Message      string      `json:"message"`
	MessageLocal string      `json:"message_local"`
	ReturnType   string      `json:"return_type"`
	Data         interface{} `json:"data"`
}

func NewResponser() *localResponser {
	return &localResponser{}
}

// ResponseError will response error back to clients 4xx
func (r *localResponser) ResponseError(c echo.Context, message string, messageLocal string, returnType string, data interface{}) error {
	body := &responseBody{
		Message:      message,
		MessageLocal: messageLocal,
		ReturnType:   returnType,
		Data:         data,
	}
	return c.JSON(http.StatusBadRequest, body)
}

// ResponseSuccess will response 200 back to client with data
func (r *localResponser) ResponseSuccess(c echo.Context, returnType string, message string, messageLocal string, data interface{}) error {
	if message == "" {
		message = "fetched successfully"
	}
	if messageLocal == "" {
		messageLocal = "fetched successfully"
	}
	body := &responseBody{
		Message:      message,
		MessageLocal: messageLocal,
		ReturnType:   returnType,
		Data:         data,
	}
	return c.JSON(http.StatusOK, body)
}
