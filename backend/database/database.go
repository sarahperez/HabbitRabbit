// package database
package database

// code from https://github.com/Duomly/go-bank-backend

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"main/helpers"
	"main/interfaces"
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

func InitTestDatabase() {
	database, err := gorm.Open(sqlite.Open("forTestingPurposeOnly.db"), &gorm.Config{})
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

func GetCalItems(user uint) []interfaces.CalendarItem {
	var ret []interfaces.CalendarItem
	DB.Table("calendar_items").Where("user = ?", user).Find(&ret)

	return ret
}

// -----------------------------Functions that work in the database - Friends ------------------------------
func GetFriends(name string) []string {
	var ret []string
	DB.Table("friend_statuses").Where("reciever = ? and status = ?", name, "Accepted").Select("requester").Find(&ret)
	var ret1 []string
	DB.Table("friend_statuses").Where("requester = ? and status = ?", name, "Accepted").Select("reciever").Find(&ret1)
	return append(ret, ret1...)
}

func GetRequests(name string) []string {
	var ret []string
	DB.Table("friend_statuses").Where("reciever = ? and status = ?", name, "Requested").Select("requester").Find(&ret)
	return ret
}

func GetBlocked(name string) []string {
	var ret []string
	DB.Table("friend_statuses").Where("reciever = ? and status = ?", name, "Blocked").Select("requester").Find(&ret)
	return ret
}
