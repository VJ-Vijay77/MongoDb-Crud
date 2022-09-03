package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatalln("couldnt load env data...")
	}
}