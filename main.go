package main

import (
	"log"
	"test-case-vhiweb/internal/app"
	"test-case-vhiweb/internal/logger"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error loading .env file", err)
	}

	logger.SetLogrusLogger()

	app.App()
}
