package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luis16121013/courseSoft/app/midleware"
	"github.com/luis16121013/courseSoft/app/storage"
)

func CreateRoutes(db *storage.MysqlReponsitory, app *fiber.App) {
	RouteApp := app.Group("/app", midleware.Authentication)
	RouteApp.Post("/loguear", LoguearUser(db))
	RouteApp.Get("/inicio", Home)
	//RouteApp.Get("/users",ListUser)
	//RouteApp.Get("/admins",ListAdmin)
}
