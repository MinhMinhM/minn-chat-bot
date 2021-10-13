


package main

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"minh-chat-bot/core"
	"minh-chat-bot/db"
	"minh-chat-bot/test"
	"os"
	"time"
)


func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://5f3652fafbed40f0a8c3e3fe8bffecaa@o994759.ingest.sentry.io/5953410",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	//sentry.CaptureMessage("It works!")

	//load /env file
	err=godotenv.Load(".env")
	if err != nil {
		fmt.Println("Cannot load .env file")
	}

	//connect to DB
	mySql:= db.NewMySQLGorm()

	//add mysql to base context before giving it to all service
	ctx:=&core.BaseContext{
		Mysql: mySql,
	}
	//initiate echo and register service
	r := echo.New()
	registerAllServices(r,ctx)

	//get and open port
	port:=os.Getenv("port")
	r.Logger.Fatal(r.Start(":"+port))
}


func registerAllServices(r *echo.Echo,ctx *core.BaseContext) {
	registerTestService(r,ctx)
}

func registerTestService(r *echo.Echo,ctx *core.BaseContext) {
	test.Register(r,ctx)
}
