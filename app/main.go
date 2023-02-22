package main

import (
	"financials/app/cmd"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Couldn't read .env file")
	}

	viper.SetDefault("DB_HOST", os.Getenv("DB_HOST"))
	viper.SetDefault("DB_USER", os.Getenv("DB_USER"))
	viper.SetDefault("DB_PASSWORD", os.Getenv("DB_PASSWORD"))

	cmd.Execute()
}
