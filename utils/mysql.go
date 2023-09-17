package utils

import (
	"fmt"
	"pokemon-fight/configs"
	"strings"
	"time"

	"pokemon-fight/models"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func InitDB(config *configs.ServerConfig) *gorm.DB {

	set := config.Database.MySQL
	dsnString := []string{
		set.Username, ":", set.Password, "@tcp(", set.Host, ":", set.Port, ")/", set.Name, "?parseTime=true&loc=Asia%2FJakarta&charset=utf8mb4&collation=utf8mb4_unicode_ci"}
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

func PostgreSQL(config *configs.ServerConfig) *gorm.DB {

	set := config.Database.PostgreSQL

	sslmode := "disable"
	timeZone := "Asia/Jakarta"
	user := set.Name
	password := set.Password
	host := set.Host
	port := set.Port
	dbName := set.Name
	// dsn := fmt.Sprintf("host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai")
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v", host, user, password, dbName, port, sslmode, timeZone)

	fmt.Println("--DNS CONNECTION--")
	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // open connection

	if err != nil {
		panic(err)
	}

	db.Use(dbresolver.Register(dbresolver.Config{}).SetMaxIdleConns(10).SetMaxOpenConns(100).SetConnMaxLifetime(time.Hour))

	return db
}

func InitialMigrate(config *configs.ServerConfig, db *gorm.DB) {
	if config.Mode == "DEV" {
		db.Migrator().DropTable(&models.Level{})
		db.Migrator().DropTable(&models.User{})
		db.Migrator().DropTable(&models.Score{})
		db.Migrator().DropTable(&models.Competition{})
		db.Migrator().DropTable(&models.Season{})
		db.Migrator().DropTable(&models.Blacklist{})

		// db.Migrator().DropTable(&models.Pokemon{})

		db.AutoMigrate(&models.Level{})
		db.AutoMigrate(&models.User{})
		db.AutoMigrate(&models.Score{})
		db.AutoMigrate(&models.Competition{})
		db.AutoMigrate(&models.Season{})
		db.AutoMigrate(&models.Blacklist{})

		// db.AutoMigrate(&models.Pokemon{})

		db.Migrator().DropColumn(&models.Score{}, "TotalPoints")
		db.Migrator().DropColumn(&models.Score{}, "SeasonId")

	} else {
		db.AutoMigrate(&models.Level{})
		db.AutoMigrate(&models.User{})
		db.AutoMigrate(&models.Score{})
		db.AutoMigrate(&models.Competition{})
		db.AutoMigrate(&models.Season{})
		db.AutoMigrate(&models.Blacklist{})

		// db.AutoMigrate(&models.Pokemon{})
	}

}
