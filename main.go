package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		token := os.Getenv("DISCORD_TOKEN")
		ConnectToDiscord(token)
	} else {

		token := os.Getenv("DISCORD_TOKEN")
		ConnectToDiscord(token)
	}

}
