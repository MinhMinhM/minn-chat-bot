


package main

import (
	"github.com/labstack/echo/v4"
	"linebot-golang/test"
)


func main() {
	r := echo.New()
	registerAllServices(r)
	r.Logger.Fatal(r.Start(":1323"))
}


func registerAllServices(r *echo.Echo) {
	registerTestService(r)
}

func registerTestService(r *echo.Echo) {
	test.Register(r)
}
