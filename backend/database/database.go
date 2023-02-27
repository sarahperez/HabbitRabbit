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