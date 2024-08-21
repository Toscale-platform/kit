package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestCors(t *testing.T) {
	app := fiber.New()
	app.Use(Cors)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"ok": true})
	})

	getReq := httptest.NewRequest("GET", "/", nil)
	optionsReq := httptest.NewRequest("OPTIONS", "/", nil)

	resp, err := app.Test(getReq, 1)
	assert.Nil(t, err)

	assert.Equal(t, resp.Header.Get(fiber.HeaderAccessControlAllowOrigin), "*")
	assert.Equal(t, resp.Header.Get(fiber.HeaderAccessControlAllowMethods), "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS")
	assert.Equal(t, resp.Header.Get(fiber.HeaderAccessControlAllowHeaders), "Authorization")
	assert.Equal(t, resp.Header.Get(fiber.HeaderContentType), "application/json")

	resp, err = app.Test(optionsReq, 1)
	assert.Nil(t, err)

	assert.Equal(t, resp.Header.Get(fiber.HeaderAccessControlAllowHeaders), "*")
	assert.Equal(t, resp.Header.Get(fiber.HeaderContentType), "text/html")
}

func TestVerifyUser(t *testing.T) {
	app := fiber.New()
	app.Use(VerifyUser)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"ok": true})
	})

	getReq := httptest.NewRequest("GET", "/", nil)

	resp, err := app.Test(getReq, 1)
	assert.Nil(t, err)

	assert.Equal(t, resp.StatusCode, 403)
}
