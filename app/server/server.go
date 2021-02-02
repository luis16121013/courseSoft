package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/luis16121013/courseSoft/app/route"
)

func StartServer(port string) {
	//declare templates
	engine := html.New("./app/public", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./app/public")

	app.Get("/", LoginView)
	route.CreateRoutes(app)

	log.Fatal(app.Listen(":" + port))
}

func LoginView(c *fiber.Ctx) error {
	return c.Render("login", nil)
}
