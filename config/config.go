package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	PORT        string
	DB_USER     string
	DB_PASSWORD string
	DB_DATABASE string
	DB_HOST     string
	DB_PORT     string
}

var ENV Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Env failed to read")
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		log.Fatal("Failed connect to ENV")
	}

	log.Println("Successfully load server")
}
