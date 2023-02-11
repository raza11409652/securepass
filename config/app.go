package config

import "github.com/joho/godotenv"

func AppConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Not able to find env file")
	}
}
