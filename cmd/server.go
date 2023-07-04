package main

import (
	"fmt"
	"os"

	"template-go/config"
	"template-go/src/api"
	echohttp "template-go/tool/echo_http"
	"template-go/tool/logger"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_ENV") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			logger.LogFatalError("Error read env file", err)
		}
	}

	db := config.ConnectDB()

	svc := echohttp.NewService(db)

	e := api.Routes(svc)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
