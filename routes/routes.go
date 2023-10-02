package routes

import (
	"jobboard/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutesApi(app *fiber.App) {

}

func SetupRoutesWeb(app *fiber.App) {

}

func SetupRoutes(app *fiber.App) {
	indexHandler := handler.NewIndexHandler()
	userHandler := handler.NewUserHandler()
	jobboardHandler := handler.NewJobBoardHandlerHandler()

	app.Get("/", indexHandler.GetIndex)
	app.Get("/aboutus", indexHandler.GetAboutUs)

	app.Get("/counter", indexHandler.GetCounter)
	app.Post("/counter", indexHandler.PostCounter)

	app.Get("/users/register", userHandler.GetRegister)
	app.Post("/users/register", userHandler.PostRegister)
	app.Get("/users/login", userHandler.GetLogin)
	app.Post("/users/login", userHandler.PostLogin)

	app.Get("/jobboards", jobboardHandler.List)
	app.Get("/jobboards/:id", jobboardHandler.GetById)
	app.Get("/jobboards/:id/edit", jobboardHandler.EditById)
	app.Post("/jobboards/:id/delete", jobboardHandler.DeleteById)
	app.Post("/jobboards", jobboardHandler.Create)
}
