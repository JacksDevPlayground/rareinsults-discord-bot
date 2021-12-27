package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		token := os.Getenv("DISCORD_TOKEN")
		ConnectToDiscord(token)
	} else {
		token := os.Getenv("DISCORD_TOKEN")
		ConnectToDiscord(token)
	}

}
