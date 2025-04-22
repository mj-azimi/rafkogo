package userModel

import (
	"rafkogo/bootstrap"
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Username  string      `gorm:"column:username"`
	Password  string       `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func Migration() error {
    return bootstrap.DB.AutoMigrate(&User{})
}