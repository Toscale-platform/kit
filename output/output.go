package output

import (
	"fmt"
	"github.com/Toscale-platform/toscale-kit/log"
	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
)

type out struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func JsonMessageResult(ctx *fasthttp.RequestCtx, code int, r string) {
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

func CORSOptions(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-Type", "text/html")
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "*")
	ctx.Response.Header.SetStatusCode(200)
	ctx.Response.Header.Set("Connection", "close")
}

func OutputJson(ctx *fasthttp.RequestCtx, code int, result interface{}) {
	// Marshal provided interface into JSON structure
	jsonResult, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Error().Err(err).Send()
		JsonMessageResult(ctx, 500, "errors.common.internalError")
		return
	}
	// Write content-type, statuscode, payload
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "Authorization")
	ctx.Response.SetStatusCode(code)
	if _, err = fmt.Fprint(ctx, string(jsonResult)); err != nil {
		log.Error().Err(err).Send()
	}
	ctx.Response.Header.Set("Connection", "close")
}

func JsonNoIndent(ctx *fasthttp.RequestCtx, code int, result interface{}) {
	// Marshal provided interface into JSON structure
	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Error().Err(err).Send()
		JsonMessageResult(ctx, 500, "errors.common.internalError")
		return
	}
	// Write content-type, statuscode, payload
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "Authorization")
	ctx.Response.SetStatusCode(code)
	if _, err = fmt.Fprint(ctx, string(jsonResult)); err != nil {
		log.Error().Err(err).Send()
	}
	ctx.Response.Header.Set("Connection", "close")
}
