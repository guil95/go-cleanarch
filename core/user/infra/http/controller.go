package user

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	userApplication "github.com/guil95/go-cleanarch/core/user/application"
	userDomain "github.com/guil95/go-cleanarch/core/user/domain"
	userInfrastructure "github.com/guil95/go-cleanarch/core/user/infra/repositories"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func CreateApi(app *fiber.App, db *gorm.DB) {
	app.Get("/users", func(context *fiber.Ctx) error {
		list(context, userApplication.NewService(userInfrastructure.NewMysqlUserRepository(db)))
		return nil
	})

	app.Get("/users/:id", func(context *fiber.Ctx) error {
		findById(context, userApplication.NewService(userInfrastructure.NewMysqlUserRepository(db)))
		return nil
	})

	app.Post("/users", func(context *fiber.Ctx) error {
		save(context, userApplication.NewService(userInfrastructure.NewMysqlUserRepository(db)))
		return nil
	})

	app.Post("/users-batch", func(context *fiber.Ctx) error {
		saveBatch(context, userApplication.NewService(userInfrastructure.NewMysqlUserRepository(db)))
		return nil
	})

	app.Post("/users-async", func(context *fiber.Ctx) error {
		saveAsync(context, userApplication.NewService(userInfrastructure.NewMysqlUserRepository(db)))
		return nil
	})
}

func list(c *fiber.Ctx, service *userApplication.Service) {
	err, users := service.ListUser()

	if err != nil {
		if err == userDomain.UserNotFound {
			log.Println(fmt.Sprintf("Users not found"))
			err := c.Status(fiber.StatusNotFound).JSON(NewResponseError("Users not found"))

			if err != nil {
				return
			}

			return
		}

		err := c.Status(fiber.StatusInternalServerError).JSON(NewResponseError("Internal Server Error"))

		if err != nil {
			return
		}

		return
	}

	err = c.Status(fiber.StatusOK).JSON(users)

	if err != nil {
		return
	}

	return
}

func findById(c *fiber.Ctx, service *userApplication.Service) {
	idParam := c.Params("id")

	id, err := userDomain.StringToUUID(idParam)

	if err != nil {
		log.Println(err.Error())
		err := c.Status(fiber.StatusUnprocessableEntity).JSON(NewResponseError("Unprocessable entity"))

		if err != nil {
			return
		}

		return
	}

	err, user := service.GetUser(id)

	if err != nil {
		if err == userDomain.UserNotFound {
			log.Println(fmt.Sprintf("User %s not found", id))

			err := c.Status(http.StatusNotFound).JSON(NewResponseError("User not found"))

			if err != nil {
				return
			}

			return
		}

		log.Println(err.Error())

		err = c.Status(http.StatusInternalServerError).JSON(NewResponseError("Internal Server Error"))

		if err != nil {
			return
		}

		return
	}

	err = c.Status(http.StatusOK).JSON(user)

	if err != nil {
		return
	}

	return
}

type CreateUserPayload struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Age       int64  `json:"age" binding:"required,min=1"`
}

func save(c *fiber.Ctx, service *userApplication.Service) {
	var userPayload CreateUserPayload

	if err := c.BodyParser(userPayload); err != nil {
		log.Println(err)

		err := c.Status(http.StatusUnprocessableEntity).JSON(NewResponseError("Unprocessable entity"))

		if err != nil {
			return
		}

		return
	}

	err, userCreated := service.SaveUser(
		userDomain.NewUser(
			userPayload.Firstname,
			userPayload.Lastname,
			userPayload.Age,
		),
	)

	if err != nil {
		log.Println(err)

		if err == userDomain.UserExists {
			err := c.Status(http.StatusUnprocessableEntity).JSON(NewResponseError(
				fmt.Sprintf("User %s exists", userPayload.Firstname)),
			)

			if err != nil {
				return
			}

			return
		}

		err := c.Status(http.StatusInternalServerError).JSON(NewResponseError("Internal error"))

		if err != nil {
			return
		}

		return
	}

	err = c.Status(http.StatusCreated).JSON(userCreated)

	if err != nil {
		return
	}

	return
}

func saveBatch(c *fiber.Ctx, service *userApplication.Service) {
	file, err := c.FormFile("file")

	if err != nil {
		err := c.Status(http.StatusInternalServerError).JSON(NewResponseError("Internal Server Error"))

		if err != nil {
			return
		}

		return
	}

	err = service.SaveUserBatch(file)

	if err != nil {
		return
	}
}

func saveAsync(c *fiber.Ctx, service *userApplication.Service) {
	file, err := c.FormFile("file")

	if err != nil {
		err := c.Status(http.StatusInternalServerError).JSON(NewResponseError("Internal Server Error"))

		if err != nil {
			return
		}

		return
	}

	err = service.SaveAsync(file)

	if err != nil {
		return
	}
}
