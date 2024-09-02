# Toscale Kit

[![Go Report Card](https://goreportcard.com/badge/github.com/Toscale-platform/kit)](https://goreportcard.com/report/github.com/Toscale-platform/kit)
[![Go Reference](https://pkg.go.dev/badge/github.com/Toscale-platform/kit.svg)](https://pkg.go.dev/github.com/Toscale-platform/kit)

A small toolkit for creating microservices.

## Packages:

### Env

Thanks: [joho/godotenv](https://github.com/joho/godotenv)

Create `.env` file in the root of your project or run app with env variables.
```dotenv
TOKEN=ABC
PORT=8080
EXCHANGES=binance,bitfinex
DEBUG=true
```

Then in your Go app you can do something like:
```go
import "github.com/Toscale-platform/kit/env"

token := env.GetString("TOKEN")
port := env.GetInt("PORT")
exchanges := env.GetSlice("EXCHANGES")
debug := env.GetBool("DEBUG")
```

### Log

Thanks: [rs/zerolog](https://github.com/rs/zerolog)

```go
import "github.com/Toscale-platform/kit/log"

log.Error().Msg("Error message")
log.Info().Str("key", "value").Msg("Info message")
```

### HTTP

Thanks: [valyala/fasthttp](https://github.com/valyala/fasthttp)

```go
import "github.com/Toscale-platform/kit/http"

http.Get("https://example.com", nil)

body := Body{}
http.Post("https://example.com", &body, nil)
```

### GraphQL

Thanks: [machinebox/graph](https://github.com/machinebox/graphql)

```go
import (
    "time"
    "github.com/Toscale-platform/kit/http"
)

client := graphql.NewClient("https://machinebox.io/graphql")

req := graphql.NewRequest(`
    query ($key: String!) {
        items (id: $key) {
            field1
            field2
            field3
        }
    }
`)

req.Var("key", "value")
req.Header.Set("Cache-Control", "no-cache")

var respData ResponseStruct
err := client.Run(req, &respData, time.Minute)
```

### Validator

Thanks: [go-playground/validator](https://github.com/go-playground/validator)

```go
import "github.com/Toscale-platform/kit/validator"

if validator.IsExchange("binance") {
    //
}
```

### Output

```go
import "github.com/Toscale-platform/kit/output"

r := router.New()

r.GET("/path", handler)
r.OPTIONS("/path", output.CORSOptions)

func handler(ctx *fasthttp.RequestCtx){
    res := map[string]string{"foo": "bar"}
	
    output.JsonNoIndent(ctx, 200, res)
    output.JsonMessageResult(ctx, 200, "message")
}
```

### Auth

```go
import "github.com/Toscale-platform/kit/auth"

isDebug := env.GetBool("DEBUG")
host := env.GetString("AUTH_HOST")
authManager := auth.Init(host, isDebug)

r := router.New()

r.GET("/path", authManager.IsAdmin(handler))
```
