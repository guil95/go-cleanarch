package user

import (
	"bufio"
	"encoding/csv"
	"github.com/guil95/go-cleanarch/core/user/domain"
	"io"
	"log"
	"mime/multipart"
	"strconv"
)

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

func (s Service) ListUser() (error, *[]user.User) {
	err, u := s.repo.List()

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

func (s Service) SaveUserBatch(file multipart.File) error {
	reader := csv.NewReader(bufio.NewReader(file))
	var userSlice []*user.User
	userLength := 0
	userMaxSimultaneous := 5000
	counterRoutines := 0
	//errorc := make(chan error)

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
			return err
		}

		age, _ := strconv.ParseInt(line[2], 10, 64)

		userSlice = append(userSlice, user.NewUser(line[0], line[1], age))

		userLength++

		if userLength >= userMaxSimultaneous {
			counterRoutines++
			go func(users []*user.User) {
				//errorc <- s.repo.CreateBatch(users, userMaxSimultaneous)
				_ = s.repo.CreateBatch(users, userMaxSimultaneous)
			}(userSlice)

			userLength = 0
			userSlice = []*user.User{}
		}

	}

	counterRoutines++

	go func(users []*user.User) {
		_ = s.repo.CreateBatch(users, userMaxSimultaneous)
		//errorc <- s.repo.CreateBatch(users, userMaxSimultaneous)
	}(userSlice)

	//for i := 0; i < counterRoutines; i++ {
	//	if err := <-errorc; err != nil {
	//		log.Println(err)
	//	}
	//}

	return nil
}
