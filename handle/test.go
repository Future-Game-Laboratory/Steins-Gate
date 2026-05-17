package handle

import (
	"github.com/gofiber/fiber/v3"
)

func HelloWorld(c fiber.Ctx) error {
	return c.SendString("hello world")
}
