package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	userInfrastructure "github.com/guil95/go-cleanarch/user/infra/http"
	"log"
	"net/http"
)

func Run(db *sql.DB) {
	log.Println("Listen server on :8000")
	r := mux.NewRouter()

	userInfrastructure.CreateApi(r, db)

	log.Fatal(http.ListenAndServe(":8000", r))
}
