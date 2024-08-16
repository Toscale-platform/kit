package middlewares

import "github.com/gofiber/fiber/v2"

func FiberCORS(c *fiber.Ctx) error {
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
