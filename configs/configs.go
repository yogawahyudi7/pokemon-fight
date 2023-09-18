package configs

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

type ServerConfig struct {
	Name string
	Port string
	Mode string

	ServerStorage string

	GoogleCloudStorage struct {
		Crendential    string
		BuketName      string
		BuketImagePath string
	}
	Database struct {
		MySQL struct {
			Driver   string
			Name     string
			Host     string
			Port     string
			Username string
			Password string
		}
		PostgreSQL struct {
			Driver   string
			Name     string
			Host     string
			Port     string
			Username string
			Password string
		}
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

	defaultConfig.ServerStorage = os.Getenv("SERVER_STORAGE")

	defaultConfig.GoogleCloudStorage.Crendential = os.Getenv("GOOGLE_CLOUD_CREDENTIAL")
	defaultConfig.GoogleCloudStorage.BuketImagePath = os.Getenv("GOOGLE_CLOUD_BUKET_IMAGE_PATH")
	defaultConfig.GoogleCloudStorage.BuketName = os.Getenv("GOOGLE_CLOUD_BUKET_NAME")

	defaultConfig.Database.MySQL.Driver = os.Getenv("MYSQL_DRIVER")
	defaultConfig.Database.MySQL.Name = os.Getenv("MYSQL_NAME")
	defaultConfig.Database.MySQL.Host = os.Getenv("MYSQL_HOST")
	defaultConfig.Database.MySQL.Port = os.Getenv("MYSQL_PORT")
	defaultConfig.Database.MySQL.Username = os.Getenv("MYSQL_USERNAME")
	defaultConfig.Database.MySQL.Password = os.Getenv("MYSQL_PASSWORD")

	defaultConfig.Database.PostgreSQL.Driver = os.Getenv("POSTGRESQL_DRIVER")
	defaultConfig.Database.PostgreSQL.Name = os.Getenv("PostgreSQL_NAME")
	defaultConfig.Database.PostgreSQL.Host = os.Getenv("PostgreSQL_HOST")
	defaultConfig.Database.PostgreSQL.Port = os.Getenv("PostgreSQL_PORT")
	defaultConfig.Database.PostgreSQL.Username = os.Getenv("PostgreSQL_USERNAME")
	defaultConfig.Database.PostgreSQL.Password = os.Getenv("POSTGRESQL_PASSWORD")

	return &defaultConfig
}
