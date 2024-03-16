package auth

import (
	"errors"
	"fmt"
	"github.com/Toscale-platform/toscale-kit/log"
	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
	"io"
	"net/http"
)

type out struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

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

func (a *Auth) isAdmin(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		if !a.isDebug {
			id, err := verifyAdmin(ctx, a.host)
			if err != nil {
				outputJsonMessageResult(ctx, 403, "forbidden")
				return
			}
			ctx.SetUserValue("user", id)
		} else {
			ctx.SetUserValue("user", 0)
		}
		next(ctx)
	}
}

func verifyAdmin(ctx *fasthttp.RequestCtx, host string) (int, error) {
	token := string(ctx.Request.Header.Peek("Authorization"))
	if token == "" {
		return 0, errors.New("bearer token required")
	}

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

func outputJsonMessageResult(ctx *fasthttp.RequestCtx, code int, r string) {
	// Write content-type, statuscode, payload
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "Authorization")
	ctx.Response.Header.SetStatusCode(code)
	out := out{code, r}
	jsonResult, _ := json.Marshal(out)
	if _, err := fmt.Fprint(ctx, string(jsonResult)); err != nil {
		log.Error().Err(err).Send()
	}
	ctx.Response.Header.Set("Connection", "close")
}
