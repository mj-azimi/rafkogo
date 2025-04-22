package chats

import (
	"rafkogo/bootstrap"
	"time" // اضافه کردن پکیج time
)

type Chat struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	ME        bool      `gorm:"column:me"`
	ClientID  int       `gorm:"column:client_id"`
	Text      string    `gorm:"column:text"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func Create(chat Chat) error {
	return bootstrap.DB.Create(&chat).Error
}

func FindById(id int) (*Chat, error) {
	var chat Chat
	err := bootstrap.DB.Where("id = ?", id).First(&chat).Error
	if err != nil {
		return nil, err
	}
	return &chat, nil
}

func All(clientID int) ([]Chat, error) {
	var chats []Chat
	err := bootstrap.DB.Where("client_id = ?", clientID).Find(&chats).Error
	return chats, err
}

func Delete(id int) error {
	return bootstrap.DB.Delete(&Chat{}, id).Error
}

func Update(id int, newText string) error {
	return bootstrap.DB.Model(&Chat{}).Where("id = ?", id).Update("text", newText).Error
}