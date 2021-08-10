package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	userInfrastructure "github.com/guil95/go-cleanarch/core/user/infra/http"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func Run(db *gorm.DB) {
	log.Println("Listen server on :8000")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Poc golang clean arch"})
	})

	userInfrastructure.CreateApi(router, db)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("API_PORT")), router))
}
