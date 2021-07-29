package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/guil95/go-cleanarch/api"
	db "github.com/guil95/go-cleanarch/pkg/mysql"
)

func main() {
	mysqlDatabase := db.Connect()

	api.Run(mysqlDatabase)
}
