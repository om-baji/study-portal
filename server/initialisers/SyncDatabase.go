package initialisers

import (
	"github.com/om-baji/models"
)

func SyncDatabase() {

	clientDb.AutoMigrate(&models.User{})
}
