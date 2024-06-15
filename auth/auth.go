package auth

import (
	"errors"
	"github.com/Toscale-platform/kit/log"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"

	"github.com/Toscale-platform/kit/output"
	"github.com/valyala/fasthttp"
)

type response struct {
	User int `json:"user"`
}

type Auth struct {
	isDebug bool
	host    string
}

var httpClient = http.Client{}

func Init(host string, isDebug bool) *Auth {
	return &Auth{
		isDebug: isDebug,
		host:    host,
	}
}

func (a *Auth) IsAdmin(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		if !a.isDebug {
			id, err := VerifyAdmin(ctx, a.host)
			if err != nil {
				output.JsonMessageResult(ctx, 403, "forbidden")
				return
			}

			ctx.SetUserValue("user", id)
		} else {
			ctx.SetUserValue("user", 0)
		}

		next(ctx)
	}
}

func (a *Auth) IsAdminFiber(c *fiber.Ctx) error {
	if !a.isDebug {
		id, err := VerifyAdminFiber(c, a.host)
		if err != nil {
			return fiber.NewError(403, "forbidden")
		}

		c.Locals("user", id)
	} else {
		c.Locals("user", 0)
	}

	return c.Next()
}

func VerifyAdmin(ctx *fasthttp.RequestCtx, host string) (int, error) {
	token := string(ctx.Request.Header.Peek("Authorization"))
	if token == "" {
		return 0, errors.New("bearer token required")
	}

	return internalVerifyAdmin(token, host)
}

func VerifyAdminFiber(c *fiber.Ctx, host string) (int, error) {
	token := string(c.Request().Header.Peek("Authorization"))
	if token == "" {
		return 0, errors.New("bearer token required")
	}

	return internalVerifyAdmin(token, host)
}

func internalVerifyAdmin(token, host string) (int, error) {
	req, err := http.NewRequest("POST", host+"/verifyAdmin", nil)
	if err != nil {
		return 0, errors.New("http request making error: " + err.Error())
	}

	req.Header.Add("Authorization", token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return 0, errors.New("http request error: " + err.Error())
	}

	if resp.StatusCode == 401 || resp.StatusCode == 403 {
		return 0, errors.New("forbidden")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, errors.New("body parsing error: " + err.Error())
	}

	r := response{}

	if err := json.Unmarshal(body, &r); err != nil {
		return 0, errors.New("unmarshal: " + err.Error())
	}

	if err := resp.Body.Close(); err != nil {
		log.Error().Err(err).Send()
	}

	return r.User, nil
}
