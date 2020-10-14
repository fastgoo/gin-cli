package main

import (
	"log"

	"gin-cli/modules/api_v1"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	api_v1.Init()
}
