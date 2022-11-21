package main

import (
	"financials/app/cmd"
	"github.com/spf13/viper"
	"os"
)

func main() {
	viper.SetDefault("DB_HOST", os.Getenv("DB_HOST"))
	viper.SetDefault("DB_USER", os.Getenv("DB_USER"))
	viper.SetDefault("DB_PASSWORD", os.Getenv("DB_PASSWORD"))

	cmd.Execute()
}