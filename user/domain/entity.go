package user

//User struct
type User struct {
	UUID UUID
	FirstName string `json:"name"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

func NewUser(firstName string, lastName string, age int) *User {
	return &User{
		UUID: NewUUID(),
		FirstName: firstName,
		LastName: lastName,
		Age: age,
	}
}
