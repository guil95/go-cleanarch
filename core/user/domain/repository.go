package user

//Repository interface to user repository
type Repository interface {
	Get(id UUID) (error, *User)
	List() (error, *[]User)
	Create(user *User) (error, *User)
	CreateBatch(user []*User) error
	SearchByName(userName string) (error, *User)
}
