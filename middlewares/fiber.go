package middlewares

import (
	"github.com/Toscale-platform/kit/auth"
	"github.com/Toscale-platform/kit/env"
	"github.com/gofiber/fiber/v2"
)

var (
	Debug = env.GetBool("DEBUG")
	Host  = "https://auth.toscale.io"
)

func Cors(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS")

	if c.Method() == fiber.MethodOptions {
		c.Set("Access-Control-Allow-Headers", "*")
		c.Set("Content-Type", "text/html")
		c.Status(fiber.StatusOK)
		return nil
	}

	c.Set("Access-Control-Allow-Headers", "Authorization")
	return c.Next()
}

func VerifyUser(c *fiber.Ctx) error {
	if !Debug {
		id, err := auth.VerifyUserFiber(c, Host)
		if err != nil {
			return fiber.NewError(403, "forbidden")
		}

		c.Locals("user", id)
	} else {
		c.Locals("user", uint64(0))
	}

	return c.Next()
}
