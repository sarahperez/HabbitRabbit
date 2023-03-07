// package migrations
package migrations

// code from https://github.com/Duomly/go-bank-backend

import (
	"main/database"
	"main/interfaces"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Refactor Migrate
func Migrate() {
	User := &interfaces.User{}
	database.DB.AutoMigrate(&User)
}
