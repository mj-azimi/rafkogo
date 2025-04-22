package main

import (
	"log"
	"rafkogo/bootstrap"
	"rafkogo/config"
	migrationregister "rafkogo/migration_register"
)

func main() {
	app := bootstrap.App()

	migrationregister.MigrationRegister();
	port := config.GetEnv("APP_PORT", "8080")
	if err := app.Run(":" + port); err != nil {
		log.Fatal("خطا در اجرای سرور:", err)
	}
}