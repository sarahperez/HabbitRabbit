// package database
package database

// code from https://github.com/Duomly/go-bank-backend

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"main/helpers"
)

// Create global variable
var DB *gorm.DB

// Create InitDatabase function
// lets us connect with the database and return the connection object
func InitDatabase() {
	database, err := gorm.Open(sqlite.Open("backend.db"), &gorm.Config{})
	helpers.HandleErr(err)
	DB = database
}

// -----------------------------Functions that work in the database - TO DO ------------------------------
func GetCompletedItems(user uint) []string {
	completedTodoItems := GetTodoItems(true, user)
	return completedTodoItems
}

func GetIncompleteItems(user uint) []string {
	incompleteTodoItems := GetTodoItems(false, user)
	return incompleteTodoItems
}

func GetTodoItems(completed bool, user uint) []string {
	var ret []string
	DB.Table("todo_items").Where("completed = ? AND user = ?", completed, user).Select("description").Find(&ret)

	return ret
}
