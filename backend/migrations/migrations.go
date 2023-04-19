// package migrations
package migrations

//code derived from https://github.com/Duomly/go-bank-backend/tree/Golang-course-Lesson-6

import (
	"main/database"
	"main/interfaces"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// create todo_items table
func MigrateToDo() {
	Todo := &interfaces.TodoItem{}
	database.DB.AutoMigrate(&Todo)
}

// create calendar_items table
func MigrateCal() {
	calendar := &interfaces.CalendarItem{}
	database.DB.AutoMigrate(&calendar)
}

// create friend_statuses table
func MigrateFriends() {
	friend := &interfaces.FriendStatus{}
	database.DB.AutoMigrate(&friend)
}
