// Code generated by goa v3.14.6, DO NOT EDIT.
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
	"strings"

	server "github.com/4x2Hach1/learning/next-go/api/gen/server"
	serverviews "github.com/4x2Hach1/learning/next-go/api/gen/server/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
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

// BuildLoginRequest instantiates a HTTP request object with method and path
// set to call the "server" service "login" endpoint
func (c *Client) BuildLoginRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: LoginServerPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("server", "login", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeLoginRequest returns an encoder for requests sent to the server login
// server.
func EncodeLoginRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*server.LoginPayload)
		if !ok {
			return goahttp.ErrInvalidType("server", "login", "*server.LoginPayload", v)
		}
		body := NewLoginRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("server", "login", err)
		}
		return nil
	}
}

// DecodeLoginResponse returns a decoder for responses returned by the server
// login endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeLoginResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				return nil, goahttp.ErrDecodingError("server", "login", err)
			}
			return body, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("server", "login", resp.StatusCode, string(body))
		}
	}
}

// BuildAuthUserRequest instantiates a HTTP request object with method and path
// set to call the "server" service "authUser" endpoint
func (c *Client) BuildAuthUserRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AuthUserServerPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("server", "authUser", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAuthUserRequest returns an encoder for requests sent to the server
// authUser server.
func EncodeAuthUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*server.AuthUserPayload)
		if !ok {
			return goahttp.ErrInvalidType("server", "authUser", "*server.AuthUserPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeAuthUserResponse returns a decoder for responses returned by the
// server authUser endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeAuthUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body AuthUserResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("server", "authUser", err)
			}
			p := NewAuthUserUserOK(&body)
			view := "default"
			vres := &serverviews.User{Projected: p, View: view}
			if err = serverviews.ValidateUser(vres); err != nil {
				return nil, goahttp.ErrValidationError("server", "authUser", err)
			}
			res := server.NewUser(vres)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("server", "authUser", resp.StatusCode, string(body))
		}
	}
}

// BuildUsersRequest instantiates a HTTP request object with method and path
// set to call the "server" service "users" endpoint
func (c *Client) BuildUsersRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UsersServerPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("server", "users", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUsersRequest returns an encoder for requests sent to the server users
// server.
func EncodeUsersRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*server.UsersPayload)
		if !ok {
			return goahttp.ErrInvalidType("server", "users", "*server.UsersPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeUsersResponse returns a decoder for responses returned by the server
// users endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeUsersResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body UsersResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("server", "users", err)
			}
			res := NewUsersUserOK(body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("server", "users", resp.StatusCode, string(body))
		}
	}
}

// BuildUserByIDRequest instantiates a HTTP request object with method and path
// set to call the "server" service "userById" endpoint
func (c *Client) BuildUserByIDRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id int
	)
	{
		p, ok := v.(*server.UserByIDPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("server", "userById", "*server.UserByIDPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UserByIDServerPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("server", "userById", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUserByIDRequest returns an encoder for requests sent to the server
// userById server.
func EncodeUserByIDRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*server.UserByIDPayload)
		if !ok {
			return goahttp.ErrInvalidType("server", "userById", "*server.UserByIDPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeUserByIDResponse returns a decoder for responses returned by the
// server userById endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeUserByIDResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body UserByIDResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("server", "userById", err)
			}
			p := NewUserByIDUserOK(&body)
			view := "default"
			vres := &serverviews.User{Projected: p, View: view}
			if err = serverviews.ValidateUser(vres); err != nil {
				return nil, goahttp.ErrValidationError("server", "userById", err)
			}
			res := server.NewUser(vres)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("server", "userById", resp.StatusCode, string(body))
		}
	}
}

// BuildNewUserRequest instantiates a HTTP request object with method and path
// set to call the "server" service "newUser" endpoint
func (c *Client) BuildNewUserRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: NewUserServerPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("server", "newUser", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeNewUserRequest returns an encoder for requests sent to the server
// newUser server.
func EncodeNewUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*server.NewUserPayload)
		if !ok {
			return goahttp.ErrInvalidType("server", "newUser", "*server.NewUserPayload", v)
		}
		body := NewNewUserRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("server", "newUser", err)
		}
		return nil
	}
}

// DecodeNewUserResponse returns a decoder for responses returned by the server
// newUser endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeNewUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body bool
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("server", "newUser", err)
			}
			return body, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("server", "newUser", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateUserRequest instantiates a HTTP request object with method and
// path set to call the "server" service "updateUser" endpoint
func (c *Client) BuildUpdateUserRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateUserServerPath()}
	req, err := http.NewRequest("PATCH", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("server", "updateUser", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateUserRequest returns an encoder for requests sent to the server
// updateUser server.
func EncodeUpdateUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*server.UpdateUserPayload)
		if !ok {
			return goahttp.ErrInvalidType("server", "updateUser", "*server.UpdateUserPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewUpdateUserRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("server", "updateUser", err)
		}
		return nil
	}
}

// DecodeUpdateUserResponse returns a decoder for responses returned by the
// server updateUser endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeUpdateUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body bool
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("server", "updateUser", err)
			}
			return body, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("server", "updateUser", resp.StatusCode, string(body))
		}
	}
}

// BuildMemoriesRequest instantiates a HTTP request object with method and path
// set to call the "server" service "memories" endpoint
func (c *Client) BuildMemoriesRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: MemoriesServerPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("server", "memories", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeMemoriesRequest returns an encoder for requests sent to the server
// memories server.
func EncodeMemoriesRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*server.MemoriesPayload)
		if !ok {
			return goahttp.ErrInvalidType("server", "memories", "*server.MemoriesPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeMemoriesResponse returns a decoder for responses returned by the
// server memories endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeMemoriesResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body MemoriesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("server", "memories", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateMemoryResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("server", "memories", err)
			}
			res := NewMemoriesMemoryOK(body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("server", "memories", resp.StatusCode, string(body))
		}
	}
}

// BuildMemoryByIDRequest instantiates a HTTP request object with method and
// path set to call the "server" service "memoryById" endpoint
func (c *Client) BuildMemoryByIDRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id int
	)
	{
		p, ok := v.(*server.MemoryByIDPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("server", "memoryById", "*server.MemoryByIDPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: MemoryByIDServerPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("server", "memoryById", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeMemoryByIDRequest returns an encoder for requests sent to the server
// memoryById server.
func EncodeMemoryByIDRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*server.MemoryByIDPayload)
		if !ok {
			return goahttp.ErrInvalidType("server", "memoryById", "*server.MemoryByIDPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeMemoryByIDResponse returns a decoder for responses returned by the
// server memoryById endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeMemoryByIDResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body MemoryByIDResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("server", "memoryById", err)
			}
			p := NewMemoryByIDMemoryOK(&body)
			view := "default"
			vres := &serverviews.Memory{Projected: p, View: view}
			if err = serverviews.ValidateMemory(vres); err != nil {
				return nil, goahttp.ErrValidationError("server", "memoryById", err)
			}
			res := server.NewMemory(vres)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("server", "memoryById", resp.StatusCode, string(body))
		}
	}
}

// BuildNewMemoryRequest instantiates a HTTP request object with method and
// path set to call the "server" service "newMemory" endpoint
func (c *Client) BuildNewMemoryRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: NewMemoryServerPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("server", "newMemory", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeNewMemoryRequest returns an encoder for requests sent to the server
// newMemory server.
func EncodeNewMemoryRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*server.NewMemoryPayload)
		if !ok {
			return goahttp.ErrInvalidType("server", "newMemory", "*server.NewMemoryPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewNewMemoryRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("server", "newMemory", err)
		}
		return nil
	}
}

// DecodeNewMemoryResponse returns a decoder for responses returned by the
// server newMemory endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeNewMemoryResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body bool
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("server", "newMemory", err)
			}
			return body, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("server", "newMemory", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteMemoryRequest instantiates a HTTP request object with method and
// path set to call the "server" service "deleteMemory" endpoint
func (c *Client) BuildDeleteMemoryRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id int
	)
	{
		p, ok := v.(*server.DeleteMemoryPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("server", "deleteMemory", "*server.DeleteMemoryPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteMemoryServerPath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("server", "deleteMemory", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteMemoryRequest returns an encoder for requests sent to the server
// deleteMemory server.
func EncodeDeleteMemoryRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*server.DeleteMemoryPayload)
		if !ok {
			return goahttp.ErrInvalidType("server", "deleteMemory", "*server.DeleteMemoryPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeDeleteMemoryResponse returns a decoder for responses returned by the
// server deleteMemory endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeDeleteMemoryResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body bool
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("server", "deleteMemory", err)
			}
			return body, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("server", "deleteMemory", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateMemoryRequest instantiates a HTTP request object with method and
// path set to call the "server" service "updateMemory" endpoint
func (c *Client) BuildUpdateMemoryRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id int
	)
	{
		p, ok := v.(*server.UpdateMemoryPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("server", "updateMemory", "*server.UpdateMemoryPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateMemoryServerPath(id)}
	req, err := http.NewRequest("PATCH", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("server", "updateMemory", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateMemoryRequest returns an encoder for requests sent to the server
// updateMemory server.
func EncodeUpdateMemoryRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*server.UpdateMemoryPayload)
		if !ok {
			return goahttp.ErrInvalidType("server", "updateMemory", "*server.UpdateMemoryPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewUpdateMemoryRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("server", "updateMemory", err)
		}
		return nil
	}
}

// DecodeUpdateMemoryResponse returns a decoder for responses returned by the
// server updateMemory endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeUpdateMemoryResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body bool
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("server", "updateMemory", err)
			}
			return body, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("server", "updateMemory", resp.StatusCode, string(body))
		}
	}
}

// BuildNewHeavyRequest instantiates a HTTP request object with method and path
// set to call the "server" service "newHeavy" endpoint
func (c *Client) BuildNewHeavyRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: NewHeavyServerPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("server", "newHeavy", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeNewHeavyRequest returns an encoder for requests sent to the server
// newHeavy server.
func EncodeNewHeavyRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*server.NewHeavyPayload)
		if !ok {
			return goahttp.ErrInvalidType("server", "newHeavy", "*server.NewHeavyPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeNewHeavyResponse returns a decoder for responses returned by the
// server newHeavy endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeNewHeavyResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				return nil, goahttp.ErrDecodingError("server", "newHeavy", err)
			}
			return body, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("server", "newHeavy", resp.StatusCode, string(body))
		}
	}
}

// BuildCheckHeavyRequest instantiates a HTTP request object with method and
// path set to call the "server" service "checkHeavy" endpoint
func (c *Client) BuildCheckHeavyRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		key string
	)
	{
		p, ok := v.(*server.CheckHeavyPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("server", "checkHeavy", "*server.CheckHeavyPayload", v)
		}
		key = p.Key
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CheckHeavyServerPath(key)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("server", "checkHeavy", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCheckHeavyRequest returns an encoder for requests sent to the server
// checkHeavy server.
func EncodeCheckHeavyRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*server.CheckHeavyPayload)
		if !ok {
			return goahttp.ErrInvalidType("server", "checkHeavy", "*server.CheckHeavyPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeCheckHeavyResponse returns a decoder for responses returned by the
// server checkHeavy endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeCheckHeavyResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body int
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("server", "checkHeavy", err)
			}
			return body, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("server", "checkHeavy", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteHeavyRequest instantiates a HTTP request object with method and
// path set to call the "server" service "deleteHeavy" endpoint
func (c *Client) BuildDeleteHeavyRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		key string
	)
	{
		p, ok := v.(*server.DeleteHeavyPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("server", "deleteHeavy", "*server.DeleteHeavyPayload", v)
		}
		key = p.Key
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteHeavyServerPath(key)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("server", "deleteHeavy", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteHeavyRequest returns an encoder for requests sent to the server
// deleteHeavy server.
func EncodeDeleteHeavyRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*server.DeleteHeavyPayload)
		if !ok {
			return goahttp.ErrInvalidType("server", "deleteHeavy", "*server.DeleteHeavyPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeDeleteHeavyResponse returns a decoder for responses returned by the
// server deleteHeavy endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeDeleteHeavyResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body bool
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("server", "deleteHeavy", err)
			}
			return body, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("server", "deleteHeavy", resp.StatusCode, string(body))
		}
	}
}

// unmarshalUserResponseToServerUser builds a value of type *server.User from a
// value of type *UserResponse.
func unmarshalUserResponseToServerUser(v *UserResponse) *server.User {
	res := &server.User{
		ID:    v.ID,
		Name:  v.Name,
		Email: v.Email,
	}

	return res
}

// unmarshalMemoryResponseToServerMemory builds a value of type *server.Memory
// from a value of type *MemoryResponse.
func unmarshalMemoryResponseToServerMemory(v *MemoryResponse) *server.Memory {
	res := &server.Memory{
		ID:       v.ID,
		UsersID:  v.UsersID,
		Title:    v.Title,
		Date:     v.Date,
		Location: v.Location,
		Detail:   v.Detail,
	}

	return res
}
