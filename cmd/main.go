package main

import (
	"jobboard/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")
	config := fiber.Config{
		Views: engine,
		// ViewsLayout: config.ViewsLayout,
	}

	app := fiber.New(config)
	app.Static("/", "./public")

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":7050"))
}
