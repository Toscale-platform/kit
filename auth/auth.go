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

type TotalAdminPermission struct {
	IsAvaliableTools         bool `json:"isAvaliableTools,omitempty"`
	IsAvaliableTerminals     bool `json:"isAvaliableTerminals,omitempty"`
	IsAvaliableUsers         bool `json:"isAvaliableUsers,omitempty"`
	IsAvaliableBalances      bool `json:"isAvaliableBalances,omitempty"`
	IsAvaliableDocumentation bool `json:"isAvaliableDocumentation,omitempty"`
	IsAvaliableInsights      bool `json:"isAvaliableInsights,omitempty"`
	IsAvaliableBalancer      bool `json:"isAvaliableBalancer,omitempty"`
	IsAvaliableNews          bool `json:"isAvaliableNews,omitempty"`
	IsAvaliableTwitter       bool `json:"isAvaliableTwitter,omitempty"`
	IsAvaliableForex         bool `json:"isAvaliableForex,omitempty"`
	IsAvaliableLanguages	 bool `json:"isAvaliableLanguages,omitempty"`
}

var httpClient = http.Client{}

func Init(host string, isDebug bool) *Auth {
	return &Auth{
		isDebug: isDebug,
		host:    host,
	}
}

func (a *Auth) GetAdminPermissions(ctx *fasthttp.RequestCtx, serviceName string) (permissions TotalAdminPermission, err error)	{
		if !a.isDebug {
			permissions, err = FetchAdminPermissions(ctx, a.host)
			if err != nil {
				return
			}
		} else {
			permissions = TotalAdminPermission{
				IsAvaliableTools:         true,
				IsAvaliableTerminals:     true,
				IsAvaliableUsers:         true,
				IsAvaliableBalances:      true,
				IsAvaliableDocumentation: true,
				IsAvaliableInsights:      true,
				IsAvaliableBalancer:      true,
				IsAvaliableNews:          true,
				IsAvaliableTwitter:       true,
				IsAvaliableForex:         true,
				IsAvaliableLanguages:	  true,
			}
		}
		return permissions, nil
}

func (a *Auth) ValidateAdminPermissions(next fasthttp.RequestHandler, serviceName string) fasthttp.RequestHandler	{
	return func(ctx *fasthttp.RequestCtx) {
		var permissions TotalAdminPermission
		var err error
		if !a.isDebug {
			permissions, err = FetchAdminPermissions(ctx, a.host)
			if err != nil {
				output.JsonMessageResult(ctx, 403, "forbidden")
				return
			}
		} else {
			permissions = TotalAdminPermission{
				IsAvaliableTools:         true,
				IsAvaliableTerminals:     true,
				IsAvaliableUsers:         true,
				IsAvaliableBalances:      true,
				IsAvaliableDocumentation: true,
				IsAvaliableInsights:      true,
				IsAvaliableBalancer:      true,
				IsAvaliableNews:          true,
				IsAvaliableTwitter:       true,
				IsAvaliableForex:         true,
			}
		}
		isInvalid := false
		switch serviceName {
			case "terminals" : {
				if !permissions.IsAvaliableTerminals {
					isInvalid = true
				}
			}
			case "tools" : {
				if !permissions.IsAvaliableTools {
					isInvalid = true
				}
			} 
			case "users" : {
				if !permissions.IsAvaliableUsers {
					isInvalid = true
				}
			}
			case "balances" : {
				if !permissions.IsAvaliableBalances {
					isInvalid = true
				}				
			}
			case "documentation" : {
				if !permissions.IsAvaliableDocumentation {
					isInvalid = true
				}				
			}
			case "insights" : {
				if !permissions.IsAvaliableInsights {
					isInvalid = true
				}				
			}
			case "balancer" : {
				if !permissions.IsAvaliableBalancer {
					isInvalid = true
				}				
			}
			case "news" : {
				if !permissions.IsAvaliableNews {
					isInvalid = true
				}				
			}
			case "twitter" : {
				if !permissions.IsAvaliableTwitter {
					isInvalid = true
				}				
			}
			case "forex" : {
				if !permissions.IsAvaliableForex {
					isInvalid = true
				}				
			}
			case "languages" : {
				if !permissions.IsAvaliableLanguages {
					isInvalid = true
				}
			}
			default : {
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
	if token == "" {
		return TotalAdminPermission{}, errors.New("bearer token required")
	}

	return internalGetPermissions(token, host)
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return TotalAdminPermission{}, errors.New("body parsing error: " + err.Error())
	}

	r := TotalAdminPermission{}

	if err := json.Unmarshal(body, &r); err != nil {
		return TotalAdminPermission{}, errors.New("unmarshal: " + err.Error())
	}

	if err := resp.Body.Close(); err != nil {
		log.Error().Err(err).Send()
	}

	return r, nil
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
