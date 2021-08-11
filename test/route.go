package test

import (

	"github.com/labstack/echo/v4"
)

// Register to register all exposed api routes
func Register(r *echo.Echo) {
	//r.Use(h.validateToken)
	//rp := responser.NewResponser()
	h := NewHandler()
	// userRoute := r.Group("/registration", h.validateAdminRole)
	testRoute := r.Group("/webhook")
	testRoute.GET("", h.webhook)
}
