package migrationregister

import (
	chatModel "rafkogo/internal/user/database/chat"
	userModel "rafkogo/internal/user/database/user"
)

func MigrationRegister(){
	userModel.Migration()
	chatModel.Migration()
}