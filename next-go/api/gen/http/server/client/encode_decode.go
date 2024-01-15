// Code generated by goa v3.14.5, DO NOT EDIT.
//
// server HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/4x2Hach1/learning/next-go/api/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	server "github.com/4x2Hach1/learning/next-go/api/gen/server"
	goahttp "goa.design/goa/v3/http"
)

// BuildHelloRequest instantiates a HTTP request object with method and path
// set to call the "server" service "hello" endpoint
func (c *Client) BuildHelloRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		name string
	)
	{
		p, ok := v.(*server.HelloPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("server", "hello", "*server.HelloPayload", v)
		}
		name = p.Name
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: HelloServerPath(name)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("server", "hello", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeHelloResponse returns a decoder for responses returned by the server
// hello endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeHelloResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("server", "hello", err)
			}
			return body, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("server", "hello", resp.StatusCode, string(body))
		}
	}
}
