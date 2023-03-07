package utils

import (
	"fmt"
	"pokemon-fight/configs"
	"strings"
	"time"

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

// func InitialMigrate(db *gorm.DB) {
// 	if config.Mode == "development" {
// 		db.Migrator().DropTable(&models.DetailTransaction{})
// 		db.Migrator().DropTable(&models.Transaction{})
// 		db.Migrator().DropTable(&models.Product{})
// 		db.Migrator().DropTable(&models.Rating{})
// 		db.Migrator().DropTable(&models.Cashout{})
// 		db.Migrator().DropTable(&models.Partner{})
// 		db.Migrator().DropTable(&models.User{})

// 		db.AutoMigrate(&models.User{})
// 		db.AutoMigrate(&models.Product{})
// 		db.AutoMigrate(&models.Transaction{})
// 		db.AutoMigrate(&models.Partner{})
// 		db.AutoMigrate(&models.Rating{})
// 		db.AutoMigrate(&models.Cashout{})

// 		seeder.AdminSeeder(db)
// 		seeder.UserSeeder(db)
// 		seeder.PartnerSeeder(db)
// 		seeder.ProductSeeder(db)
// 	} else {
// 		db.AutoMigrate(&models.User{})
// 		db.AutoMigrate(&models.Product{})
// 		db.AutoMigrate(&models.Transaction{})
// 		db.AutoMigrate(&models.Partner{})
// 		db.AutoMigrate(&models.Rating{})
// 		db.AutoMigrate(&models.Cashout{})
// 	}

// }
