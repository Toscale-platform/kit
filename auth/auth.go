package auth

import (
	"errors"
	"github.com/Toscale-platform/kit/env"
	"github.com/Toscale-platform/kit/log"
	"github.com/Toscale-platform/kit/output"
	"github.com/Toscale-platform/kit/validator"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"net/http"
)

type response struct {
	User uint64 `json:"user"`
}

type TotalAdminPermission struct {
	IsAvailableTools          bool `json:"tools,omitempty"`
	IsAvailableTerminals      bool `json:"terminals,omitempty"`
	IsAvailableUsers          bool `json:"users,omitempty"`
	IsAvailableBackendTesting bool `json:"backendTesting,omitempty"`
	IsAvailableDocumentation  bool `json:"documentation,omitempty"`
	IsAvailableInsights       bool `json:"insights,omitempty"`
	IsAvailableBalancer       bool `json:"balancer,omitempty"`
	IsAvailableNews           bool `json:"news,omitempty"`
	IsAvailableTwitter        bool `json:"twitter,omitempty"`
	IsAvailableForex          bool `json:"forex,omitempty"`
	IsAvailableLanguages      bool `json:"languages,omitempty"`
}

var (
	Debug      = env.GetBool("DEBUG", false)
	Host       = env.GetString("AUTH_HOST", "https://auth.toscale.io")
	httpClient = http.Client{}
)

func GetAdminPermissions(ctx *fasthttp.RequestCtx) (permissions TotalAdminPermission, err error) {
	if !Debug {
		permissions, err = FetchAdminPermissions(ctx, Host)
		if err != nil {
			return
		}
	} else {
		permissions = TotalAdminPermission{
			IsAvailableTools:          true,
			IsAvailableTerminals:      true,
			IsAvailableUsers:          true,
			IsAvailableBackendTesting: true,
			IsAvailableDocumentation:  true,
			IsAvailableInsights:       true,
			IsAvailableBalancer:       true,
			IsAvailableNews:           true,
			IsAvailableTwitter:        true,
			IsAvailableForex:          true,
			IsAvailableLanguages:      true,
		}
	}

	return permissions, nil
}

func ValidateAdminPermissions(next fasthttp.RequestHandler, serviceName string) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var permissions TotalAdminPermission
		var err error

		if !Debug {
			permissions, err = FetchAdminPermissions(ctx, Host)
			if err != nil {
				output.JsonMessageResult(ctx, 403, "forbidden")
				return
			}
		} else {
			permissions = TotalAdminPermission{
				IsAvailableTools:          true,
				IsAvailableTerminals:      true,
				IsAvailableUsers:          true,
				IsAvailableBackendTesting: true,
				IsAvailableDocumentation:  true,
				IsAvailableInsights:       true,
				IsAvailableBalancer:       true,
				IsAvailableNews:           true,
				IsAvailableTwitter:        true,
				IsAvailableForex:          true,
			}
		}

		isInvalid := false

		switch serviceName {
		case "terminals":
			{
				if !permissions.IsAvailableTerminals {
					isInvalid = true
				}
			}
		case "tools":
			{
				if !permissions.IsAvailableTools {
					isInvalid = true
				}
			}
		case "users":
			{
				if !permissions.IsAvailableUsers {
					isInvalid = true
				}
			}
		case "backendTesting":
			{
				if !permissions.IsAvailableBackendTesting {
					isInvalid = true
				}
			}
		case "documentation":
			{
				if !permissions.IsAvailableDocumentation {
					isInvalid = true
				}
			}
		case "insights":
			{
				if !permissions.IsAvailableInsights {
					isInvalid = true
				}
			}
		case "balancer":
			{
				if !permissions.IsAvailableBalancer {
					isInvalid = true
				}
			}
		case "news":
			{
				if !permissions.IsAvailableNews {
					isInvalid = true
				}
			}
		case "twitter":
			{
				if !permissions.IsAvailableTwitter {
					isInvalid = true
				}
			}
		case "forex":
			{
				if !permissions.IsAvailableForex {
					isInvalid = true
				}
			}
		case "languages":
			{
				if !permissions.IsAvailableLanguages {
					isInvalid = true
				}
			}
		default:
			{
				output.JsonMessageResult(ctx, 400, "invalid service")
				return
			}
		}

		if isInvalid {
			output.JsonMessageResult(ctx, 403, "forbidden")
			return
		}

		next(ctx)
	}
}

func FetchAdminPermissions(ctx *fasthttp.RequestCtx, host string) (perms TotalAdminPermission, err error) {
	token := string(ctx.Request.Header.Peek("Authorization"))
	if validator.IsEmpty(token) {
		return TotalAdminPermission{}, errors.New("bearer token required")
	}

	return internalGetPermissions(token, host)
}

func IsAdmin(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		if !Debug {
			id, err := VerifyAdmin(ctx, Host)
			if err != nil {
				output.JsonMessageResult(ctx, 403, "forbidden")
				return
			}

			ctx.SetUserValue("user", id)
		} else {
			ctx.SetUserValue("user", uint64(0))
		}

		next(ctx)
	}
}

func IsAdminFiber(c *fiber.Ctx) error {
	if !Debug {
		id, err := VerifyAdminFiber(c, Host)
		if err != nil {
			return fiber.NewError(403, "forbidden")
		}

		c.Locals("user", id)
	} else {
		c.Locals("user", uint64(0))
	}

	return c.Next()
}

func VerifyAdmin(ctx *fasthttp.RequestCtx, host string) (uint64, error) {
	token := string(ctx.Request.Header.Peek("Authorization"))
	if validator.IsEmpty(token) {
		return 0, errors.New("bearer token required")
	}

	return internalVerifyAdmin(token, host)
}

func VerifyUserFiber(c *fiber.Ctx, host string) (uint64, error) {
	token := c.Get("Authorization")
	if validator.IsEmpty(token) {
		return 0, errors.New("bearer token required")
	}

	return internalVerifyUser(token, host)
}

func VerifyAdminFiber(c *fiber.Ctx, host string) (uint64, error) {
	token := c.Get("Authorization")
	if validator.IsEmpty(token) {
		return 0, errors.New("bearer token required")
	}

	return internalVerifyAdmin(token, host)
}

func internalGetPermissions(token, host string) (perms TotalAdminPermission, err error) {
	req, err := http.NewRequest("POST", host+"/adminPermissions", nil)
	if err != nil {
		return TotalAdminPermission{}, errors.New("http request making error: " + err.Error())
	}

	req.Header.Add("Authorization", token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return TotalAdminPermission{}, errors.New("http request error: " + err.Error())
	}

	if resp.StatusCode == 401 || resp.StatusCode == 403 {
		return TotalAdminPermission{}, errors.New("forbidden")
	}

	r := TotalAdminPermission{}

	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return TotalAdminPermission{}, errors.New("unmarshal: " + err.Error())
	}

	if err := resp.Body.Close(); err != nil {
		log.Error().Err(err).Send()
	}

	return r, nil
}

func internalVerifyUser(token, host string) (uint64, error) {
	req, err := http.NewRequest("POST", host+"/verifyUser", nil)
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

	r := response{}

	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return 0, errors.New("unmarshal: " + err.Error())
	}

	if err := resp.Body.Close(); err != nil {
		log.Error().Err(err).Send()
	}

	return r.User, nil
}

func internalVerifyAdmin(token, host string) (uint64, error) {
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

	r := response{}

	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return 0, errors.New("unmarshal: " + err.Error())
	}

	if err := resp.Body.Close(); err != nil {
		log.Error().Err(err).Send()
	}

	return r.User, nil
}
