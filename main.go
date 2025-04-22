package main

import (
	"log"
	"rafkogo/bootstrap"
	"rafkogo/config"
	"rafkogo/internal/user/database/chats"
)

func main() {
	app := bootstrap.App()


	if err := chats.Migration(); err != nil {
		log.Fatal("❌ خطا در اجرای مایگریشن:", err)
	}

	port := config.GetEnv("APP_PORT", "8080")
	if err := app.Run(":" + port); err != nil {
		log.Fatal("خطا در اجرای سرور:", err)
	}
}