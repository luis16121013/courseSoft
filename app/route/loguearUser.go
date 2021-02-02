package route

import "github.com/gofiber/fiber/v2"

func LoguearUser() fiber.Handler{
	return func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		if username == "luis" && password == "1234" {
			return c.Redirect("/app/inicio", 302)
		}
		return c.Redirect("/", 302)
	}
}
