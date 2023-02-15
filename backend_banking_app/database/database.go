package database

import (
	"example.com/backend_banking_app/helpers"
	"github.com/jinzhu/gorm"
)

// Create global variable
var DB *gorm.DB

// Create InitDatabase function
//lets us connect with the database and return the connection object
func InitDatabase() {
	database, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=bankapp password=postgres sslmode=disable")
	helpers.HandleErr(err)
	// Set up connection pool
	database.DB().SetMaxIdleConns(20)
	database.DB().SetMaxOpenConns(200)
	DB = database
}