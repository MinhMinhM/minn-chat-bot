package test

import (

	"github.com/labstack/echo/v4"
)

// Register to register all exposed api routes
func Register(r *echo.Echo) {
	h := NewHandler()
	testRoute := r.Group("/webhook")
	testRoute.POST("", h.WebHook)
}
