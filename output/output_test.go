package output

import (
	"bufio"
	"fmt"
	"github.com/Toscale-platform/kit/tests"
	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
	"testing"
)

func makeRequest(handler func(ctx *fasthttp.RequestCtx)) (resp fasthttp.Response) {
	server := &fasthttp.Server{Handler: handler}

	ln := fasthttputil.NewInmemoryListener()
	defer ln.Close()

	go func() {
		if err := server.Serve(ln); err != nil {
			panic(err)
		}
	}()

	c, err := ln.Dial()
	if err != nil {
		panic(err)
	}
	defer c.Close()

	if _, err = c.Write([]byte("GET / HTTP/1.1\nHost: aa\n\n")); err != nil {
		panic(err)
	}

	r := bufio.NewReader(c)

	if err := resp.Read(r); err != nil {
		panic(err)
	}

	return
}

func TestCORSOptions(t *testing.T) {
	resp := makeRequest(func(ctx *fasthttp.RequestCtx) {
		CORSOptions(ctx)
	})

	tests.Equal(t, string(resp.Header.Peek("Content-Type")), "text/html")
}

func TestOutputJson(t *testing.T) {
	resp := makeRequest(func(ctx *fasthttp.RequestCtx) {
		OutputJson(ctx, 200, out{Code: 123, Message: "test"})
	})

	tests.Equal(t, string(resp.Body()), `{"code":123,"message":"test"}`)
}

func TestJsonNoIndent(t *testing.T) {
	resp := makeRequest(func(ctx *fasthttp.RequestCtx) {
		JsonNoIndent(ctx, 200, out{Code: 123, Message: "test"})
	})

	tests.Equal(t, string(resp.Body()), `{"code":123,"message":"test"}`)
}

func TestJsonMessageResult(t *testing.T) {
	resp := makeRequest(func(ctx *fasthttp.RequestCtx) {
		JsonMessageResult(ctx, 200, "test")
	})

	tests.Equal(t, string(resp.Body()), `{"code":200,"message":"test"}`)
}

func getResult() []byte {
	jsonResult, err := json.Marshal(out{200, "test"})
	if err != nil {
		panic(err)
	}
	return jsonResult
}

func TestFprint(t *testing.T) {
	jsonResult := getResult()

	resp := makeRequest(func(ctx *fasthttp.RequestCtx) {
		if _, err := fmt.Fprint(ctx, string(jsonResult)); err != nil {
			t.Error(err)
		}
	})

	tests.Equal(t, string(resp.Body()), `{"code":200,"message":"test"}`)
}

func TestWrite(t *testing.T) {
	jsonResult := getResult()

	resp := makeRequest(func(ctx *fasthttp.RequestCtx) {
		if _, err := ctx.Write(jsonResult); err != nil {
			t.Error(err)
		}
	})

	tests.Equal(t, string(resp.Body()), `{"code":200,"message":"test"}`)
}

func BenchmarkFprint(b *testing.B) {
	jsonResult := getResult()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		makeRequest(func(ctx *fasthttp.RequestCtx) {
			if _, err := fmt.Fprint(ctx, string(jsonResult)); err != nil {
				b.Error(err)
			}
		})
	}
}

func BenchmarkWrite(b *testing.B) {
	jsonResult := getResult()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		makeRequest(func(ctx *fasthttp.RequestCtx) {
			if _, err := ctx.Write(jsonResult); err != nil {
				b.Error(err)
			}
		})
	}
}
