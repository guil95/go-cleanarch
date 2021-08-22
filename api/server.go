package api

import (
	"github.com/gofiber/fiber/v2"
	userInfrastructure "github.com/guil95/go-cleanarch/core/user/infra/http"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func Run(db *mongo.Database) {
	log.Println("Listen server on :8000")

	app := fiber.New(fiber.Config{
		BodyLimit: 1024 * 1024 * 1024, // this is the default limit of 4MB
	})

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
