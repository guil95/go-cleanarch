package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/guil95/go-cleanarch/config"
	userApplication "github.com/guil95/go-cleanarch/user/application"
	userDomain "github.com/guil95/go-cleanarch/user/domain"
	userInfra "github.com/guil95/go-cleanarch/user/infra/repositories"
	"log"
)

func main() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	userService := userApplication.NewService(userInfra.NewMysqlUserRepository(db))
	fmt.Println(userService.GetUser(userDomain.NewUUID()))
}
