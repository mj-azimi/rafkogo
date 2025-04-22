package database

import "rafkogo/internal/user/database/chats"


func Handle() {
	chats.Migration()
}
