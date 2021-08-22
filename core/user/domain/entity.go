package user

import (
	"errors"
)

var UserNotFound = errors.New("User not found")
var UserExists = errors.New("User existis")

//User struct
type User struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name" bson:"name"`
	Lastname   string `json:"lastname" bson:"lastname"`
	Age        int64  `json:"age" bson:"age"`
}

func NewUser(firstName string, lastName string, age int64) *User {
	return &User{
		Identifier: NewUUID().String(),
		Name:       firstName,
		Lastname:   lastName,
		Age:        age,
	}
}
