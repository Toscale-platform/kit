package output

import (
	"github.com/Toscale-platform/kit/log"
	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
)

type out struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func withHeaders(ctx *fasthttp.RequestCtx, contentType, allowHeaders string, code int) {
	ctx.Response.Header.Set("Content-Type", contentType)
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", allowHeaders)
	ctx.Response.Header.SetStatusCode(code)
	ctx.Response.Header.Set("Connection", "close")
}

func withDefaultHeaders(ctx *fasthttp.RequestCtx, code int) {
	withHeaders(ctx, "application/json", "Authorization", code)
}

func CORSOptions(ctx *fasthttp.RequestCtx) {
	withHeaders(ctx, "text/html", "*", 200)
}

// OutputJson does the same thing that JsonNoIndent does
//
// Deprecated: In previous versions this method added indentation, this is no longer relevant, use JsonNoIndent now
func OutputJson(ctx *fasthttp.RequestCtx, code int, result interface{}) {
	JsonNoIndent(ctx, code, result)
}

// JsonNoIndent marshaling and writing message without indent
func JsonNoIndent(ctx *fasthttp.RequestCtx, code int, result interface{}) {
	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Error().Err(err).Send()
		JsonMessageResult(ctx, 500, "errors.kit.internalError")
		return
	}

	if _, err := ctx.Write(jsonResult); err != nil {
		log.Error().Err(err).Send()
	}

	withDefaultHeaders(ctx, code)
}

// JsonMessageResult writing text message without indent
func JsonMessageResult(ctx *fasthttp.RequestCtx, code int, r string) {
	jsonResult, err := json.Marshal(out{code, r})
	if err != nil {
		log.Error().Err(err).Send()
		JsonMessageResult(ctx, 500, "errors.kit.internalError")
		return
	}

	if _, err := ctx.Write(jsonResult); err != nil {
		log.Error().Err(err).Send()
	}

	withDefaultHeaders(ctx, code)
}
