package graphql

import (
	"github.com/valyala/fasthttp"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/matryer/is"
)

func makeServer(handler func(ctx *fasthttp.RequestCtx)) (string, func() error) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}

	go func() {
		if err := fasthttp.Serve(ln, handler); err != nil {
			panic(err)
		}
	}()

	return "http://" + ln.Addr().String(), ln.Close
}

func TestDoJSON(t *testing.T) {
	is := is.New(t)

	var calls int
	url, cl := makeServer(func(ctx *fasthttp.RequestCtx) {
		calls++
		is.Equal(string(ctx.Request.Header.Method()), http.MethodPost)
		b := ctx.Request.Body()
		is.Equal(string(b), `{"query":"query {}","variables":null}`)
		ctx.WriteString(`{
			"data": {
				"something": "yes"
			}
		}`)
	})
	defer cl()

	client := NewClient(url)

	var responseData map[string]interface{}
	err := client.Run(&Request{q: "query {}"}, &responseData, 30*time.Second)
	is.NoErr(err)
	is.Equal(calls, 1) // calls
	is.Equal(responseData["something"], "yes")
}

func TestDoJSONBadRequestErr(t *testing.T) {
	is := is.New(t)

	var calls int
	url, cl := makeServer(func(ctx *fasthttp.RequestCtx) {
		calls++
		is.Equal(string(ctx.Request.Header.Method()), http.MethodPost)
		b := ctx.Request.Body()
		is.Equal(string(b), `{"query":"query {}","variables":null}`)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.WriteString(`{
			"errors": [{
				"message": "miscellaneous message as to why the the request was bad"
			}]
		}`)
	})
	defer cl()

	client := NewClient(url)

	var responseData map[string]interface{}
	err := client.Run(&Request{q: "query {}"}, &responseData, 30*time.Second)
	is.Equal(calls, 1) // calls
	is.Equal(err.Error(), "graphql: miscellaneous message as to why the the request was bad")
}

func TestQueryJSON(t *testing.T) {
	is := is.New(t)

	var calls int
	url, cl := makeServer(func(ctx *fasthttp.RequestCtx) {
		calls++
		b := ctx.Request.Body()
		is.Equal(string(b), `{"query":"query {}","variables":{"username":"matryer"}}`)
		_, err := ctx.WriteString(`{"data":{"value":"some data"}}`)
		is.NoErr(err)
	})
	defer cl()

	client := NewClient(url)

	req := NewRequest("query {}")
	req.Var("username", "matryer")

	// check variables
	is.True(req != nil)
	is.Equal(req.vars["username"], "matryer")

	var resp struct {
		Value string
	}
	err := client.Run(req, &resp, 1*time.Second)
	is.NoErr(err)
	is.Equal(calls, 1)

	is.Equal(resp.Value, "some data")
}

func TestHeader(t *testing.T) {
	is := is.New(t)

	var calls int
	url, cl := makeServer(func(ctx *fasthttp.RequestCtx) {
		calls++
		is.Equal(string(ctx.Request.Header.Peek("X-Custom-Header")), "123")

		_, err := ctx.WriteString(`{"data":{"value":"some data"}}`)
		is.NoErr(err)
	})
	defer cl()

	client := NewClient(url)

	req := NewRequest("query {}")
	req.Header.Set("X-Custom-Header", "123")

	var resp struct {
		Value string
	}
	err := client.Run(req, &resp, 1*time.Second)
	is.NoErr(err)
	is.Equal(calls, 1)

	is.Equal(resp.Value, "some data")
}
