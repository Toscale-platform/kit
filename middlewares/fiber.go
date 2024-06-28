package middlewares

import "github.com/gofiber/fiber/v2"

func FiberCORS(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS")

	if c.Method() == "OPTIONS" {
		c.Set("Access-Control-Allow-Headers", "*")
	} else {
		c.Set("Access-Control-Allow-Headers", "Authorization")
	}

	return c.Next()
}
