package graphql

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
	"net/http"
	"time"
)

// Client is a client for interacting with a GraphQL API.
type Client struct {
	endpoint   string
	httpClient *fasthttp.Client
}

// NewClient makes a new Client capable of making GraphQL requests.
func NewClient(endpoint string) *Client {
	return &Client{
		endpoint:   endpoint,
		httpClient: &fasthttp.Client{StreamResponseBody: false},
	}
}

// Run executes the query and unmarshals the response from the data field
// into the response object.
// Pass in a nil response object to skip response parsing.
// If the request fails or the server returns an error, the first error
// will be returned.
func (c *Client) Run(req *Request, resp interface{}, timeout time.Duration) error {
	requestBodyObj := internalRequest{
		Query:     req.q,
		Variables: req.vars,
	}

	requestBody, err := json.Marshal(requestBodyObj)
	if err != nil {
		return fmt.Errorf("encode body: %w", err)
	}

	gr := &graphResponse{
		Data: resp,
	}

	r := fasthttp.AcquireRequest()
	r.Header.SetMethod(fasthttp.MethodPost)
	r.SetRequestURI(c.endpoint)
	r.SetBodyRaw(requestBody)

	r.Header.SetContentType("application/json; charset=utf-8")
	r.Header.Set("Accept", "application/json; charset=utf-8")

	for key, values := range req.Header {
		for _, value := range values {
			r.Header.Add(key, value)
		}
	}

	res := fasthttp.AcquireResponse()

	if err := c.httpClient.DoTimeout(r, res, timeout); err != nil {
		return err
	}

	fasthttp.ReleaseRequest(r)
	defer fasthttp.ReleaseResponse(res)

	if err := json.Unmarshal(res.Body(), &gr); err != nil {
		return fmt.Errorf("decoding response: %w", err)
	}

	if len(gr.Errors) > 0 {
		// return first error
		return gr.Errors[0]
	}

	return nil
}

type internalRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

type graphErr struct {
	Message string
}

func (e graphErr) Error() string {
	return "graphql: " + e.Message
}

type graphResponse struct {
	Data   interface{}
	Errors []graphErr
}

// Request is a GraphQL request.
type Request struct {
	q    string
	vars map[string]interface{}

	// Header represent any request headers that will be set
	// when the request is made.
	Header http.Header
}

// NewRequest makes a new Request with the specified string.
func NewRequest(q string) *Request {
	return &Request{
		q:      q,
		Header: make(map[string][]string),
	}
}

// Var sets a variable.
func (req *Request) Var(key string, value interface{}) {
	if req.vars == nil {
		req.vars = make(map[string]interface{})
	}
	req.vars[key] = value
}

// Vars gets the variables for this Request.
func (req *Request) Vars() map[string]interface{} {
	return req.vars
}

// Query gets the query string of this request.
func (req *Request) Query() string {
	return req.q
}
