// this code was taken from a golang banking app tutorial.
// we plan to use this outline to implement our own username and login, and adapt accordingly
// link to the video tutorials that was followed: https://www.youtube.com/playlist?list=PLi21Ag9n6jJJ5bq77cLYpCgOaONcQNqm0
package main

import (
	"example.com/backend_banking_app/api"
	"example.com/backend_banking_app/database"
)

func main() {
	// Do migration
	// migrations.MigrateTransactions()

	// Init database
	database.InitDatabase()
	api.StartApi()
}
