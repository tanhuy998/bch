package bootstrap

import "app/internal/db"

func InitDatabaseClient() {

	db.GetDB()
}
