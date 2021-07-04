package user

import (
	"database/sql"
	user "github.com/guil95/go-cleanarch/user/domain"
)

type MysqlUserRepository struct {
	db *sql.DB
}

func NewMysqlUserRepository(db *sql.DB) *MysqlUserRepository {
	return &MysqlUserRepository{
		db: db,
	}
}


func (repo MysqlUserRepository) Get(uuid user.UUID) (error, *user.User) {
	return nil, user.NewUser("Guilherme", "Rodrigues", 26)
}