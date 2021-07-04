package user

//Repository interface to user repository
type Repository interface {
	Get(uuid UUID) (error, *User)
}
