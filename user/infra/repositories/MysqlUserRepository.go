package user

import (
	"errors"
	userDomain "github.com/guil95/go-cleanarch/user/domain"
	"gorm.io/gorm"
	"log"
)

type MysqlUserRepository struct {
	db *gorm.DB
}

func NewMysqlUserRepository(db *gorm.DB) *MysqlUserRepository {
	return &MysqlUserRepository{
		db: db,
	}
}

func (repo MysqlUserRepository) Get(uuid userDomain.UUID) (error, *userDomain.User) {
	var u userDomain.User

	tx := repo.db.Model(u).First(&u, "identifier = ?", uuid.String())

	if tx.Error != nil {
		log.Println(tx.Error.Error())

		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return userDomain.UserNotFound, nil
		}

		return tx.Error, nil
	}

	return nil, &u
}

func (repo MysqlUserRepository) Create(user *userDomain.User) (error, *userDomain.User) {
	tx := repo.db.Create(user)

	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return tx.Error, nil
	}

	return nil, user
}

func (repo MysqlUserRepository) SearchByName(userName string) (error, *userDomain.User) {
	var u userDomain.User

	tx := repo.db.Model(u).First(&u, "name = ?", userName)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return userDomain.UserNotFound, nil
		}

		log.Println(tx.Error.Error())
		return tx.Error, nil
	}

	return nil, &u
}
