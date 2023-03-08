package utils

import (
	"fmt"
	"pokemon-fight/configs"
	"strings"
	"time"

	"pokemon-fight/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func InitDB(config *configs.ServerConfig) *gorm.DB {

	dsnString := []string{
		config.Database.Username, ":", config.Database.Password, "@tcp(", config.Database.Host, ":", config.Database.Port, ")/", config.Database.Name, "?parseTime=true&loc=Asia%2FJakarta&charset=utf8mb4&collation=utf8mb4_unicode_ci"}
	dsn := strings.Join(dsnString, "")

	fmt.Println("--DNS CONNECTION--")
	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // open connection

	if err != nil {
		panic(err)
	}

	db.Use(dbresolver.Register(dbresolver.Config{}).SetMaxIdleConns(10).SetMaxOpenConns(100).SetConnMaxLifetime(time.Hour))

	return db
}

func InitialMigrate(config *configs.ServerConfig, db *gorm.DB) {
	if config.Mode == "DEV" {
		db.Migrator().DropTable(&models.Score{})
		db.Migrator().DropTable(&models.Competition{})
		db.Migrator().DropTable(&models.Season{})
		db.Migrator().DropTable(&models.Blacklist{})
		// db.Migrator().DropTable(&models.Pokemon{})

		db.AutoMigrate(&models.Score{})
		db.AutoMigrate(&models.Competition{})
		db.AutoMigrate(&models.Season{})
		db.AutoMigrate(&models.Blacklist{})
		// db.AutoMigrate(&models.Pokemon{})

		// seeder.AdminSeeder(db)
		// seeder.UserSeeder(db)
		// seeder.PartnerSeeder(db)
		// seeder.ProductSeeder(db)
	} else {
		db.AutoMigrate(&models.Score{})
		db.AutoMigrate(&models.Competition{})
		db.AutoMigrate(&models.Season{})
		db.AutoMigrate(&models.Blacklist{})
		// db.AutoMigrate(&models.Pokemon{})
	}

}
