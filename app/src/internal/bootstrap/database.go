package bootstrap

import "app/src/internal/db"

func InitDatabaseClient() {

	db.GetDB()
}
