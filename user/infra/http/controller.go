package user

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	userApplication "github.com/guil95/go-cleanarch/user/application"
	userDomain "github.com/guil95/go-cleanarch/user/domain"
	userInfrastructure "github.com/guil95/go-cleanarch/user/infra/repositories"
	"log"
	"net/http"
)

func CreateApi(r *mux.Router, db *sql.DB) {
	r.Handle("/users/{id}", index(userApplication.NewService(userInfrastructure.NewMysqlUserRepository(db))))
}

func index(service *userApplication.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := userDomain.StringToUUID(vars["id"])

		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			err := json.NewEncoder(w).Encode(NewResponseError("Unprocessable entity"))

			if err != nil {
				return
			}

			return
		}

		err, user := service.GetUser(id)

		if err != nil {
			log.Println(err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			err := json.NewEncoder(w).Encode(NewResponseError("Internal server error"))

			if err != nil {
				log.Println(err.Error())
				return
			}

			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			return
		}
	})
}