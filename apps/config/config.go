package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
)

type AppConfig struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     int
	DBName     string
}

func InitConfig() *AppConfig {
	return ReadENV()
}

func ReadENV() *AppConfig {
	app := AppConfig{}
	isRead := true

	if val, found := os.LookupEnv("DB_USERNAME"); found {
		app.DBUsername = val
		isRead = false
	}

	if val, found := os.LookupEnv("DB_PASSWORD"); found {
		app.DBPassword = val
		isRead = false
	}

	if val, found := os.LookupEnv("DB_HOST"); found {
		app.DBHost = val
		isRead = false
	}

	if val, found := os.LookupEnv("DB_PORT"); found {
		valConv, err := strconv.Atoi(val)
		if err != nil {
			app.DBPort = valConv
			isRead = false
		}
	}

	if val, found := os.LookupEnv("DB_NAME"); found {
		app.DBName = val
		isRead = false
	}

	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Fatal error config file: %s \n", err)
		}
		app.DBUsername = viper.GetString("DB_USERNAME")
		app.DBPassword = viper.GetString("DB_PASSWORD")
		app.DBHost = viper.GetString("DB_HOST")
		app.DBPort = viper.GetInt("DB_PORT")
		app.DBName = viper.GetString("DB_NAME")
	}
	return &app
}
