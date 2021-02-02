package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luis16121013/courseSoft/app/midleware"
)

func CreateRoutes(app *fiber.App) {
	RouteApp := app.Group("/app", midleware.Authentication)
	RouteApp.Post("/loguear", LoguearUser())
	RouteApp.Get("/inicio", Home)
	//RouteApp.Get("/users",ListUser)
	//RouteApp.Get("/admins",ListAdmin)
}
