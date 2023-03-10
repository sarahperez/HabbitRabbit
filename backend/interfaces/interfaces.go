package interfaces

// code from https://github.com/Duomly/go-bank-backend

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string
	Name     string
	Email    string
	Password string
}

type UserID struct {
	ID uint
}

type ResponseUser struct {
	ID       uint
	Username string
	Name     string
	Email    string
}

type TodoReq struct {
	User        uint
	Description string
}

type TodoItem struct {
	gorm.Model
	User        uint
	Description string
	Completed   bool
}

// Create Validation interface
type Validation struct {
	Value string
	Valid string
}

type ErrResponse struct {
	Message string
}
