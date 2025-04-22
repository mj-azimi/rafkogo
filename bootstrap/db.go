package bootstrap

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // باید Public باشه (با D بزرگ)

func GetEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}

func DbConnect() {
    host := GetEnv("DB_HOST", "localhost")
    port := GetEnv("DB_PORT", "3306")
    dbname := GetEnv("DB_DATABASE", "rafkogo")
    username := GetEnv("DB_USERNAME", "root")
    password := GetEnv("DB_PASSWORD", "")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        username, password, host, port, dbname)

    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("❌ Error connecting to database:", err)
    }

    DB = database
    fmt.Println("✅ Connection to the database was successful.")
}
