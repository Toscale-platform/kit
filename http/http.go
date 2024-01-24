package http

import (
	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
)

var (
	client                = &fasthttp.Client{}
	headerContentTypeJson = []byte("application/json")
)

func Get(url string, v interface{}) error {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodGet)

	resp := fasthttp.AcquireResponse()

	if err := client.Do(req, resp); err != nil {
		return err
	}

	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if v != nil {
		if err := json.Unmarshal(resp.Body(), v); err != nil {
			return err
		}
	}

	return nil
}

func Post(url string, data interface{}, v interface{}) error {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentTypeBytes(headerContentTypeJson)

	if data != nil {
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return err
		}

		req.SetBodyRaw(dataBytes)
	}

	resp := fasthttp.AcquireResponse()

	if err := client.Do(req, resp); err != nil {
		return err
	}

	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if v != nil {
		if err := json.Unmarshal(resp.Body(), v); err != nil {
			return err
		}
	}

	return nil
}
