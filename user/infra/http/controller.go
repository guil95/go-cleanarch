package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	userApplication "github.com/guil95/go-cleanarch/user/application"
	userDomain "github.com/guil95/go-cleanarch/user/domain"
	userInfrastructure "github.com/guil95/go-cleanarch/user/infra/repositories"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func CreateApi(r *gin.Engine, db *gorm.DB) {
	r.GET("/users/:id", func(context *gin.Context) {
		findById(context, userApplication.NewService(userInfrastructure.NewMysqlUserRepository(db)))
	})
}

func findById(c *gin.Context,service *userApplication.Service) {
	idParam := c.Param("id")

	id, err := userDomain.StringToUUID(idParam)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnprocessableEntity, NewResponseError("Unprocessable entity"))
	}

	err, user := service.GetUser(id)

	if err != nil {
		if err == userDomain.UserNotFound {
			log.Println(fmt.Sprintf("User %s not found", id))
			c.JSON(http.StatusNotFound, NewResponseError("User not found"))
			return
		}

		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, NewResponseError("Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, user)
	return
}