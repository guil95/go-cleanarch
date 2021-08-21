package api

import (
	"github.com/gofiber/fiber/v2"
	userInfrastructure "github.com/guil95/go-cleanarch/core/user/infra/http"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Run(db *gorm.DB) {
	log.Println("Listen server on :8000")

	app := fiber.New()

	app.Get("/", func(context *fiber.Ctx) error {
		err := context.SendStatus(http.StatusOK)
		if err != nil {
			return nil
		}

		return context.JSON(fiber.Map{"message": "ok"})
	})

	userInfrastructure.CreateApi(app, db)

	log.Fatal(app.Listen(":8000"))
}
