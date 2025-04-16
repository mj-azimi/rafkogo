package bootstrap

import (
	"fmt"
	"log"
	"rafkogo/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConnect() {
	host := config.GetEnv("DB_HOST", "localhost")
	port := config.GetEnv("DB_PORT", "3306")
	dbname := config.GetEnv("DB_DATABASE", "rafkogo")
	username := config.GetEnv("DB_USERNAME", "root")
	password := config.GetEnv("DB_PASSWORD", "")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Error connecting to database:", err)
	}

	DB = database
	fmt.Println("✅ Connection to the database was successful.")
}
