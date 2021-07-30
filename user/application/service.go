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

func (s Service) GetUser(userID user.UUID) (error, *user.User) {
	err, u := s.repo.Get(userID)

	if err != nil {
		return err, nil
	}

	return nil, u
}

func (s Service) SaveUser(userToSave *user.User) (error, *user.User) {
	err, userExists := s.repo.SearchByName(userToSave.Name)

	if err != nil && err != user.UserNotFound {
		return err, nil
	}

	if userExists != nil {
		return user.UserExists, nil
	}

	err, u := s.repo.Create(userToSave)

	if err != nil {
		return err, nil
	}

	return nil, u
}
