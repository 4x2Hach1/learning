// Code generated by goa v3.14.6, DO NOT EDIT.
//
// server HTTP server types
//
// Command:
// $ goa gen github.com/4x2Hach1/learning/next-go/api/design

package server

import (
	server "github.com/4x2Hach1/learning/next-go/api/gen/server"
	serverviews "github.com/4x2Hach1/learning/next-go/api/gen/server/views"
	goa "goa.design/goa/v3/pkg"
)

// LoginRequestBody is the type of the "server" service "login" endpoint HTTP
// request body.
type LoginRequestBody struct {
	// Email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// NewUserRequestBody is the type of the "server" service "newUser" endpoint
// HTTP request body.
type NewUserRequestBody struct {
	// Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// UpdateUserRequestBody is the type of the "server" service "updateUser"
// endpoint HTTP request body.
type UpdateUserRequestBody struct {
	// Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
}

// NewMemoryRequestBody is the type of the "server" service "newMemory"
// endpoint HTTP request body.
type NewMemoryRequestBody struct {
	// Title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Date
	Date *string `form:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// Location
	Location *string `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// Detail
	Detail *string `form:"detail,omitempty" json:"detail,omitempty" xml:"detail,omitempty"`
}

// UpdateMemoryRequestBody is the type of the "server" service "updateMemory"
// endpoint HTTP request body.
type UpdateMemoryRequestBody struct {
	// Title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Date
	Date *string `form:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// Location
	Location *string `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// Detail
	Detail *string `form:"detail,omitempty" json:"detail,omitempty" xml:"detail,omitempty"`
}

// AuthUserResponseBody is the type of the "server" service "authUser" endpoint
// HTTP response body.
type AuthUserResponseBody struct {
	// ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
}

// UsersResponseBody is the type of the "server" service "users" endpoint HTTP
// response body.
type UsersResponseBody []*UserResponse

// UserByIDResponseBody is the type of the "server" service "userById" endpoint
// HTTP response body.
type UserByIDResponseBody struct {
	// ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
}

// MemoriesResponseBody is the type of the "server" service "memories" endpoint
// HTTP response body.
type MemoriesResponseBody []*MemoryResponse

// MemoryByIDResponseBody is the type of the "server" service "memoryById"
// endpoint HTTP response body.
type MemoryByIDResponseBody struct {
	// ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// UsersId
	UsersID *int `form:"users_id,omitempty" json:"users_id,omitempty" xml:"users_id,omitempty"`
	// Title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Date
	Date *string `form:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// Location
	Location *string `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// Detail
	Detail *string `form:"detail,omitempty" json:"detail,omitempty" xml:"detail,omitempty"`
}

// UserResponse is used to define fields on response body types.
type UserResponse struct {
	// ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
}

// MemoryResponse is used to define fields on response body types.
type MemoryResponse struct {
	// ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// UsersId
	UsersID *int `form:"users_id,omitempty" json:"users_id,omitempty" xml:"users_id,omitempty"`
	// Title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Date
	Date *string `form:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// Location
	Location *string `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// Detail
	Detail *string `form:"detail,omitempty" json:"detail,omitempty" xml:"detail,omitempty"`
}

// NewAuthUserResponseBody builds the HTTP response body from the result of the
// "authUser" endpoint of the "server" service.
func NewAuthUserResponseBody(res *serverviews.UserView) *AuthUserResponseBody {
	body := &AuthUserResponseBody{
		ID:    res.ID,
		Name:  res.Name,
		Email: res.Email,
	}
	return body
}

// NewUsersResponseBody builds the HTTP response body from the result of the
// "users" endpoint of the "server" service.
func NewUsersResponseBody(res []*server.User) UsersResponseBody {
	body := make([]*UserResponse, len(res))
	for i, val := range res {
		body[i] = marshalServerUserToUserResponse(val)
	}
	return body
}

// NewUserByIDResponseBody builds the HTTP response body from the result of the
// "userById" endpoint of the "server" service.
func NewUserByIDResponseBody(res *serverviews.UserView) *UserByIDResponseBody {
	body := &UserByIDResponseBody{
		ID:    res.ID,
		Name:  res.Name,
		Email: res.Email,
	}
	return body
}

// NewMemoriesResponseBody builds the HTTP response body from the result of the
// "memories" endpoint of the "server" service.
func NewMemoriesResponseBody(res []*server.Memory) MemoriesResponseBody {
	body := make([]*MemoryResponse, len(res))
	for i, val := range res {
		body[i] = marshalServerMemoryToMemoryResponse(val)
	}
	return body
}

// NewMemoryByIDResponseBody builds the HTTP response body from the result of
// the "memoryById" endpoint of the "server" service.
func NewMemoryByIDResponseBody(res *serverviews.MemoryView) *MemoryByIDResponseBody {
	body := &MemoryByIDResponseBody{
		ID:       res.ID,
		UsersID:  res.UsersID,
		Title:    res.Title,
		Date:     res.Date,
		Location: res.Location,
		Detail:   res.Detail,
	}
	return body
}

// NewHelloPayload builds a server service hello endpoint payload.
func NewHelloPayload(name string) *server.HelloPayload {
	v := &server.HelloPayload{}
	v.Name = name

	return v
}

// NewLoginPayload builds a server service login endpoint payload.
func NewLoginPayload(body *LoginRequestBody) *server.LoginPayload {
	v := &server.LoginPayload{
		Email:    *body.Email,
		Password: *body.Password,
	}

	return v
}

// NewAuthUserPayload builds a server service authUser endpoint payload.
func NewAuthUserPayload(token string) *server.AuthUserPayload {
	v := &server.AuthUserPayload{}
	v.Token = token

	return v
}

// NewUsersPayload builds a server service users endpoint payload.
func NewUsersPayload(token string) *server.UsersPayload {
	v := &server.UsersPayload{}
	v.Token = token

	return v
}

// NewUserByIDPayload builds a server service userById endpoint payload.
func NewUserByIDPayload(id int, token string) *server.UserByIDPayload {
	v := &server.UserByIDPayload{}
	v.ID = id
	v.Token = token

	return v
}

// NewNewUserPayload builds a server service newUser endpoint payload.
func NewNewUserPayload(body *NewUserRequestBody) *server.NewUserPayload {
	v := &server.NewUserPayload{
		Name:     *body.Name,
		Email:    *body.Email,
		Password: *body.Password,
	}

	return v
}

// NewUpdateUserPayload builds a server service updateUser endpoint payload.
func NewUpdateUserPayload(body *UpdateUserRequestBody, token string) *server.UpdateUserPayload {
	v := &server.UpdateUserPayload{
		Name:  *body.Name,
		Email: *body.Email,
	}
	v.Token = token

	return v
}

// NewMemoriesPayload builds a server service memories endpoint payload.
func NewMemoriesPayload(token string) *server.MemoriesPayload {
	v := &server.MemoriesPayload{}
	v.Token = token

	return v
}

// NewMemoryByIDPayload builds a server service memoryById endpoint payload.
func NewMemoryByIDPayload(id int, token string) *server.MemoryByIDPayload {
	v := &server.MemoryByIDPayload{}
	v.ID = id
	v.Token = token

	return v
}

// NewNewMemoryPayload builds a server service newMemory endpoint payload.
func NewNewMemoryPayload(body *NewMemoryRequestBody, token string) *server.NewMemoryPayload {
	v := &server.NewMemoryPayload{
		Title:    *body.Title,
		Date:     *body.Date,
		Location: *body.Location,
		Detail:   *body.Detail,
	}
	v.Token = token

	return v
}

// NewDeleteMemoryPayload builds a server service deleteMemory endpoint payload.
func NewDeleteMemoryPayload(id int, token string) *server.DeleteMemoryPayload {
	v := &server.DeleteMemoryPayload{}
	v.ID = id
	v.Token = token

	return v
}

// NewUpdateMemoryPayload builds a server service updateMemory endpoint payload.
func NewUpdateMemoryPayload(body *UpdateMemoryRequestBody, id int, token string) *server.UpdateMemoryPayload {
	v := &server.UpdateMemoryPayload{
		Title:    *body.Title,
		Date:     *body.Date,
		Location: *body.Location,
		Detail:   *body.Detail,
	}
	v.ID = id
	v.Token = token

	return v
}

// NewNewHeavyPayload builds a server service newHeavy endpoint payload.
func NewNewHeavyPayload(token string) *server.NewHeavyPayload {
	v := &server.NewHeavyPayload{}
	v.Token = token

	return v
}

// NewCheckHeavyPayload builds a server service checkHeavy endpoint payload.
func NewCheckHeavyPayload(key string, token string) *server.CheckHeavyPayload {
	v := &server.CheckHeavyPayload{}
	v.Key = key
	v.Token = token

	return v
}

// NewDeleteHeavyPayload builds a server service deleteHeavy endpoint payload.
func NewDeleteHeavyPayload(key string, token string) *server.DeleteHeavyPayload {
	v := &server.DeleteHeavyPayload{}
	v.Key = key
	v.Token = token

	return v
}

// ValidateLoginRequestBody runs the validations defined on LoginRequestBody
func ValidateLoginRequestBody(body *LoginRequestBody) (err error) {
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Password == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("password", "body"))
	}
	return
}

// ValidateNewUserRequestBody runs the validations defined on NewUserRequestBody
func ValidateNewUserRequestBody(body *NewUserRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Password == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("password", "body"))
	}
	return
}

// ValidateUpdateUserRequestBody runs the validations defined on
// UpdateUserRequestBody
func ValidateUpdateUserRequestBody(body *UpdateUserRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	return
}

// ValidateNewMemoryRequestBody runs the validations defined on
// NewMemoryRequestBody
func ValidateNewMemoryRequestBody(body *NewMemoryRequestBody) (err error) {
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Date == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("date", "body"))
	}
	if body.Location == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("location", "body"))
	}
	if body.Detail == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("detail", "body"))
	}
	if body.Date != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.date", *body.Date, goa.FormatDate))
	}
	return
}

// ValidateUpdateMemoryRequestBody runs the validations defined on
// UpdateMemoryRequestBody
func ValidateUpdateMemoryRequestBody(body *UpdateMemoryRequestBody) (err error) {
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Date == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("date", "body"))
	}
	if body.Location == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("location", "body"))
	}
	if body.Detail == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("detail", "body"))
	}
	if body.Date != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.date", *body.Date, goa.FormatDate))
	}
	return
}
