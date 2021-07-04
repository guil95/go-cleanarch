package user

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service{
	return &Service{
		repo: repo,
	}
}

func (s Service) GetUser(userID UUID) (error, *User){
	err, user := s.repo.Get(userID)

	if err != nil {
		panic(err)
	}

	return nil, user
}