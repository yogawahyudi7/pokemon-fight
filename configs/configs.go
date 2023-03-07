package configs

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

type ServerConfig struct {
	Name     string
	Port     string
	Mode     string
	Database struct {
		Driver   string
		Name     string
		Host     string
		Port     string
		Username string
		Password string
	}
}

func Get() *ServerConfig {

	err := godotenv.Load()

	if err != nil {
		log.Info("Error loading .env file")
	}

	var defaultConfig ServerConfig

	defaultConfig.Name = os.Getenv("APP_NAME")
	defaultConfig.Port = os.Getenv("APP_PORT")
	defaultConfig.Mode = os.Getenv("MODE")

	defaultConfig.Database.Driver = os.Getenv("DB_DRIVER")
	defaultConfig.Database.Name = os.Getenv("DB_NAME")
	defaultConfig.Database.Host = os.Getenv("DB_HOST")
	defaultConfig.Database.Port = os.Getenv("DB_PORT")
	defaultConfig.Database.Username = os.Getenv("DB_USERNAME")
	defaultConfig.Database.Password = os.Getenv("DB_PASSWORD")

	return &defaultConfig
}
