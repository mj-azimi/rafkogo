package chats

import "rafkogo/bootstrap"

func Migration() error {
    return bootstrap.DB.AutoMigrate(&Chat{})
}