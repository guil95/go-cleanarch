package user

import (
	"errors"
)

var UserNotFound = errors.New("User not found")
var UserExists = errors.New("User existis")

//User struct
type User struct {
	Identifier UUID   `json:"identifier"`
	Name       string `gorm:"type:string;default:null" json:"name"`
	Lastname   string `gorm:"type:string;default:null" json:"lastname"`
	Age        int    `gorm:"type:int;default:null" json:"age"`
}

func NewUser(firstName string, lastName string, age int) *User {
	return &User{
		Identifier: NewUUID(),
		Name:       firstName,
		Lastname:   lastName,
		Age:        age,
	}
}
