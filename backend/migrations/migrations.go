// package migrations
package migrations

// code from https://github.com/Duomly/go-bank-backend

import (
	"main/database"
	"main/interfaces"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Refactor Migrate
func MigrateToDo() {
	Todo := &interfaces.TodoItem{}
	database.DB.AutoMigrate(&Todo)
}

// Refactor Migrate
func MigrateCal() {
	calendar := &interfaces.CalendarItem{}
	database.DB.AutoMigrate(&calendar)
}

func MigrateFriends() {
	friend := &interfaces.FriendStatus{}
	database.DB.AutoMigrate(&friend)
}
