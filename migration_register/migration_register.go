package migrationregister

import (
	landingModel "rafkogo/internal/landing/database/landing"
	// chatModel "rafkogo/internal/user/database/chat"
	// userModel "rafkogo/internal/user/database/user"
)

func MigrationRegister(){
	// userModel.Migration()
	// chatModel.Migration()
	landingModel.Migration()
}