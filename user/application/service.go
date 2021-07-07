package user

import "github.com/guil95/go-cleanarch/user/domain"

type Service struct {
	repo user.Repository
}

func NewService(repo user.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s Service) GetUser(userID user.UUID) (error, *user.User){
	err, u := s.repo.Get(userID)

	if err != nil {
		return err, nil
	}

	return nil, u
}