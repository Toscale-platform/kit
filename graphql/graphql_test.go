package graphql

import (
	"github.com/Toscale-platform/kit/tests"
	"github.com/valyala/fasthttp"
	"net"
	"net/http"
	"testing"
	"time"
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
	var calls int
	url, cl := makeServer(func(ctx *fasthttp.RequestCtx) {
		calls++
		tests.Equal(t, string(ctx.Request.Header.Method()), http.MethodPost)

		b := ctx.Request.Body()
		tests.Equal(t, string(b), `{"query":"query {}","variables":null}`)

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
	tests.Err(t, err)

	tests.Equal(t, calls, 1) // calls
	tests.Equal(t, responseData["something"], "yes")
}

func TestDoJSONBadRequestErr(t *testing.T) {
	var calls int
	url, cl := makeServer(func(ctx *fasthttp.RequestCtx) {
		calls++
		tests.Equal(t, string(ctx.Request.Header.Method()), http.MethodPost)

		b := ctx.Request.Body()
		tests.Equal(t, string(b), `{"query":"query {}","variables":null}`)

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
	tests.HasErr(t, err)
	tests.Equal(t, calls, 1)
	tests.Equal(t, err.Error(), "graphql: miscellaneous message as to why the the request was bad")
}

func TestQueryJSON(t *testing.T) {
	var calls int
	url, cl := makeServer(func(ctx *fasthttp.RequestCtx) {
		calls++
		b := ctx.Request.Body()
		tests.Equal(t, string(b), `{"query":"query {}","variables":{"username":"matryer"}}`)

		_, err := ctx.WriteString(`{"data":{"value":"some data"}}`)
		tests.Err(t, err)
	})
	defer cl()

	client := NewClient(url)

	req := NewRequest("query {}")
	req.Var("username", "matryer")

	// check variables
	tests.Nil(t, req)
	tests.Equal(t, req.vars["username"], "matryer")

	var resp struct {
		Value string
	}

	err := client.Run(req, &resp, time.Second)
	tests.Err(t, err)

	tests.Equal(t, calls, 1)
	tests.Equal(t, resp.Value, "some data")
}

func TestHeader(t *testing.T) {
	var calls int
	url, cl := makeServer(func(ctx *fasthttp.RequestCtx) {
		calls++
		tests.Equal(t, string(ctx.Request.Header.Peek("X-Custom-Header")), "123")

		_, err := ctx.WriteString(`{"data":{"value":"some data"}}`)
		tests.Err(t, err)
	})
	defer cl()

	client := NewClient(url)

	req := NewRequest("query {}")
	req.Header.Set("X-Custom-Header", "123")

	var resp struct {
		Value string
	}

	err := client.Run(req, &resp, time.Second)
	tests.Err(t, err)

	tests.Equal(t, calls, 1)
	tests.Equal(t, resp.Value, "some data")
}

func TestRealAPI(t *testing.T) {
	client := NewClient("https://graphqlzero.almansi.me/api")

	req := NewRequest("query {post(id: 1) {id title body}}")

	var responseData map[string]interface{}

	err := client.Run(req, &responseData, 30*time.Second)
	tests.Err(t, err)

	tests.Equal(t, responseData["post"].(map[string]interface{})["id"], "1")
}
