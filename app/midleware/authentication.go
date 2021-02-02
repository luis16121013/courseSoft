package midleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Authentication(c *fiber.Ctx) error {
	fmt.Println("midleware: authentication")
	return c.Next()
}
