//package main
//
//import (
//	"github.com/labstack/echo/v4"
//)
//
//func main() {
//
//
//	e := echo.New()
//	//e.GET("/", func(c echo.Context) error {
//	//	return c.String(http.StatusOK, "ok")
//	//})
//	e.POST("/webhook", WebHook)
//	e.Logger.Fatal(e.Start(":1323"))
//
//}
//


package main

import (
	//"os"
	//"MinhChatBot/test"

	//"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"linebot-golang/test"
)


func main() {
	r := echo.New()
	registerAllServices(r)
	r.Logger.Fatal(r.Start(":8808"))
}


func registerAllServices(r *echo.Echo) {
	registerTestService(r)
}

func registerTestService(r *echo.Echo) {
	test.Register(r)
}
