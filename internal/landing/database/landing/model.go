package landingModel

import (
	"rafkogo/bootstrap"
	"time"
)

type Landing struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func Migration() error {
    return bootstrap.DB.AutoMigrate(&Landing{})
}