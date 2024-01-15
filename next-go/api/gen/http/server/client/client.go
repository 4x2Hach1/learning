// Code generated by goa v3.14.5, DO NOT EDIT.
//
// server client HTTP transport
//
// Command:
// $ goa gen github.com/4x2Hach1/learning/next-go/api/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the server service endpoint HTTP clients.
type Client struct {
	// Hello Doer is the HTTP client used to make requests to the hello endpoint.
	HelloDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the server service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		HelloDoer:           doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Hello returns an endpoint that makes HTTP requests to the server service
// hello server.
func (c *Client) Hello() goa.Endpoint {
	var (
		decodeResponse = DecodeHelloResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildHelloRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.HelloDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("server", "hello", err)
		}
		return decodeResponse(resp)
	}
}
