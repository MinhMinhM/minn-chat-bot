


package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
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
