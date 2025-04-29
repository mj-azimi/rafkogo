package landingModel

import (
	"rafkogo/utils/mysql"
	"time"
)

type Landing struct {
	ID        int        `gorm:"primaryKey;autoIncrement"`
	Title     *string    `gorm:"column:title"`
	CreatedAt time.Time  `gorm:"column:created_at;autoCreateTime"`
}

func Migration() error {
	return mysql.AutoMigrate(&Landing{})
}
