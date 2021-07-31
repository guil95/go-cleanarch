package user

//Repository interface to user repository
type Repository interface {
	Get(uuid UUID) (error, *User)
	List() (error, *[]User)
	Create(user *User) (error, *User)
	CreateBatch(user []*User, userMaxSimultaneous int) error
	SearchByName(userName string) (error, *User)
}
