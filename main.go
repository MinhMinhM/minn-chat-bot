


package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"linebot-golang/test"
	"os"
)


func main() {
	//load /env file
	err:=godotenv.Load(".env")
	if err != nil {
		println("Cannot load .env file")
	}
	//initiate echo and register service
	r := echo.New()
	registerAllServices(r)

	//get and open port
	port:=os.Getenv("port")
	r.Logger.Fatal(r.Start(":"+port))
}


func registerAllServices(r *echo.Echo) {
	registerTestService(r)
}

func registerTestService(r *echo.Echo) {
	test.Register(r)
}
