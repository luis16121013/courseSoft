package route

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/luis16121013/courseSoft/app/storage"
)

func LoguearUser(r *storage.MysqlReponsitory) fiber.Handler{
	return func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		u,err := storage.LoginUser(username, password,r)
		if err != nil {
			fmt.Println(err)
			return c.Redirect("/", 302)
		}
		fmt.Println(u)
		if u.Id != 0 {
			return c.Redirect("/app/inicio", 302)
		}
		return c.Redirect("/", 302)
	}
}
