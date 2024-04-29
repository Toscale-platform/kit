# Toscale Kit
A small toolkit for creating microservices

## Packages:

### Env

Main library: [spf13/viper](https://github.com/spf13/viper)

Create `.env` file in the root of your project or run app with env variables.
```dotenv
TOKEN=ABC
PORT=8080
EXCHANGES=binance,bitfinex
DEBUG=true
```

Then in your Go app you can do something like:
```go
package main

import "github.com/Toscale-platform/toscale-kit/env"

func main() {
    token := env.GetString("TOKEN")
    port := env.GetInt("PORT")
    exchanges := env.GetSlice("EXCHANGES")
    debug := env.GetBool("DEBUG")
}
```

### Log

Main library: [rs/zerolog](https://github.com/rs/zerolog)

```go
package main

import "github.com/Toscale-platform/toscale-kit/log"

func main() {
    log.Error().Msg("Error message")
    log.Info().Str("key", "value").Msg("Info message")
}
```

### HTTP

Main library: [valyala/fasthttp](https://github.com/valyala/fasthttp)

```go
package main

import "github.com/Toscale-platform/toscale-kit/http"

func main() {
    http.Get("https://example.com", nil)

    body := Body{}
    http.Post("https://example.com", &body, nil)
}
```

### Validator

```go
package main

import "github.com/Toscale-platform/toscale-kit/validator"

func main() {
    if validator.IsSymbol("BTC/USDT") {
        //
    }
	
    if validator.IsExchange("binance") {
        //
    }
}
```


### Output

```go
package main

import "github.com/Toscale-platform/toscale-kit/output"

func main() {
    r := router.New()
	
    r.GET("/path", handler)
    r.OPTIONS("/path", output.CORSOptions)
}

func handler(ctx *fasthttp.RequestCtx){
    res := map[string]string{"foo": "bar"}
	
    output.JsonNoIndent(ctx, 200, res)
    output.JsonMessageResult(ctx, 200, "message")
}
```

### Auth

```go
package main

import "github.com/Toscale-platform/toscale-kit/auth"

func main() {
    isDebug := env.GetBool("DEBUG")
    host := env.GetString("AUTH_HOST")
    authManager := auth.Init(host, isDebug)
    
    r := router.New()
    
    r.GET("/path", authManager.IsAdmin(handler))
}
```

### Exchange

```go
package main

import "github.com/Toscale-platform/toscale-kit/exchange"

func main() {
    symbols, err := exchange.GetSymbols("binance")
    if err != nil {
        //
    }
}
```
