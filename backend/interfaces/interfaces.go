package interfaces

// code from https://github.com/Duomly/go-bank-backend

import "github.com/jinzhu/gorm"

// most of the structs in here are necessary,
//but some of the types are not
type User struct {
	gorm.Model
	Username string
	Name     string
	Email    string
	Password string
}

type ResponseUser struct {
	ID       uint
	Username string
	Name     string
	Email    string
}

// Create Validation interface
type Validation struct {
	Value string
	Valid string
}

type ErrResponse struct {
	Message string
}
