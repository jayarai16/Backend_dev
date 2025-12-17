package routes

import (
	"github.com/gofiber/fiber/v2"
	"backend_dev_task/internal/handler"
)

func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	api := app.Group("/users")
	api.Post("/", userHandler.CreateUser)
	api.Get("/:id", userHandler.GetUser)
	api.Put("/:id", userHandler.UpdateUser)
	api.Delete("/:id", userHandler.DeleteUser)
	api.Get("/", userHandler.ListUsers)
}