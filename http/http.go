package http

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
)

const headerContentTypeJson = "application/json"

var client = &fasthttp.Client{}

func Get(url string, v interface{}) error {
	return GetWithToken(url, "", v)
}

func GetWithToken(url, token string, v interface{}) error {
	return request(fasthttp.MethodGet, url, token, nil, v)
}

func Post(url string, body interface{}, v interface{}) error {
	return PostWithToken(url, "", body, v)
}

func PostWithToken(url, token string, body interface{}, v interface{}) error {
	return request(fasthttp.MethodPost, url, token, body, v)
}

func request(method, url, token string, body interface{}, v interface{}) error {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(method)
	req.Header.SetContentType(headerContentTypeJson)

	if token != "" {
		req.Header.Add(fasthttp.HeaderAuthorization, token)
	}

	if body != nil {
		bytes, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("body parsing error: %s", err.Error())
		}

		req.SetBodyRaw(bytes)
	}

	resp := fasthttp.AcquireResponse()

	if err := client.Do(req, resp); err != nil {
		return err
	}

	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if resp.StatusCode() < 200 || resp.StatusCode() > 299 {
		return fmt.Errorf("status code %d: %s", resp.StatusCode(), string(resp.Body()))
	}

	if v != nil {
		if err := json.Unmarshal(resp.Body(), v); err != nil {
			return fmt.Errorf("body to interface parsing error: %s, body: %s", err.Error(), string(resp.Body()))
		}
	}

	return nil
}
