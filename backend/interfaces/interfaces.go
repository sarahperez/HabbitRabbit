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
	User uint
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

//---------------our added functions--------------------

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

type CalendarItem struct {
	gorm.Model
	User     uint
	EventID  int
	StartStr string
	EndStr   string
	Title    string
}

type CalObj struct {
	EventID  int
	StartStr string
	EndStr   string
	Title    string
}

type DeleteCal struct {
	EventID string
}

type FriendStatus struct {
	gorm.Model
	Requester string
	Reciever  string
	Status    string
}

type FriendRequest struct {
	gorm.Model
	Requester string
	Reciever  string
}
