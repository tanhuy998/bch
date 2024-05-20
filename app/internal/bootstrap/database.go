package bootstrap

import "app/infrastructure/db"

func InitDatabaseClient() {

	db.GetDB()
}
