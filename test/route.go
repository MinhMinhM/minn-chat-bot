package test

import (

	"github.com/labstack/echo/v4"
	"minh-chat-bot/core"
)

// Register to register all exposed api routes
func Register(r *echo.Echo,ctx *core.BaseContext) {
	h := NewHandler(ctx)
	testRoute := r.Group("/webhook")
	testRoute.POST("/MinhShop1", h.WebHook)
	testRoute.POST("/MinhShop2", h.WebHook)

	DBroute := r.Group("/DB")
	DBroute.GET("/GetNameByID", h.GetNameByID)
	DBroute.POST("/AddName", h.AddName)
}
