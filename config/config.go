package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	MongoURI      string
	MongoDatabase string
	JWTSecret     string
}

func LoadConfig() *Config {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Arquivo .env não encontrado, usando váriaveis de ambiente")
	}

	return &Config{
		MongoURI:      viper.GetString("MONGO_URI"),
		MongoDatabase: viper.GetString("MONGO_DATABASE"),
		JWTSecret:     viper.GetString("JWT_SECRET"),
	}
}
