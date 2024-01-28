// Code generated by goa v3.14.6, DO NOT EDIT.
//
// server HTTP client CLI support package
//
// Command:
// $ goa gen github.com/4x2Hach1/learning/next-go/api/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	server "github.com/4x2Hach1/learning/next-go/api/gen/server"
	goa "goa.design/goa/v3/pkg"
)

// BuildHelloPayload builds the payload for the server hello endpoint from CLI
// flags.
func BuildHelloPayload(serverHelloName string) (*server.HelloPayload, error) {
	var name string
	{
		name = serverHelloName
	}
	v := &server.HelloPayload{}
	v.Name = name

	return v, nil
}

// BuildLoginPayload builds the payload for the server login endpoint from CLI
// flags.
func BuildLoginPayload(serverLoginBody string) (*server.LoginPayload, error) {
	var err error
	var body LoginRequestBody
	{
		err = json.Unmarshal([]byte(serverLoginBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"Ea qui.\",\n      \"password\": \"Et saepe non.\"\n   }'")
		}
	}
	v := &server.LoginPayload{
		Email:    body.Email,
		Password: body.Password,
	}

	return v, nil
}

// BuildAuthUserPayload builds the payload for the server authUser endpoint
// from CLI flags.
func BuildAuthUserPayload(serverAuthUserToken string) (*server.AuthUserPayload, error) {
	var token string
	{
		token = serverAuthUserToken
	}
	v := &server.AuthUserPayload{}
	v.Token = token

	return v, nil
}

// BuildUsersPayload builds the payload for the server users endpoint from CLI
// flags.
func BuildUsersPayload(serverUsersToken string) (*server.UsersPayload, error) {
	var token string
	{
		token = serverUsersToken
	}
	v := &server.UsersPayload{}
	v.Token = token

	return v, nil
}

// BuildUserByIDPayload builds the payload for the server userById endpoint
// from CLI flags.
func BuildUserByIDPayload(serverUserByIDID string, serverUserByIDToken string) (*server.UserByIDPayload, error) {
	var err error
	var id int
	{
		var v int64
		v, err = strconv.ParseInt(serverUserByIDID, 10, strconv.IntSize)
		id = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT")
		}
	}
	var token string
	{
		token = serverUserByIDToken
	}
	v := &server.UserByIDPayload{}
	v.ID = id
	v.Token = token

	return v, nil
}

// BuildNewUserPayload builds the payload for the server newUser endpoint from
// CLI flags.
func BuildNewUserPayload(serverNewUserBody string) (*server.NewUserPayload, error) {
	var err error
	var body NewUserRequestBody
	{
		err = json.Unmarshal([]byte(serverNewUserBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"A enim praesentium modi dicta cumque.\",\n      \"name\": \"Autem tempore fugit.\",\n      \"password\": \"At molestiae aliquam pariatur.\"\n   }'")
		}
	}
	v := &server.NewUserPayload{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	}

	return v, nil
}

// BuildUpdateUserPayload builds the payload for the server updateUser endpoint
// from CLI flags.
func BuildUpdateUserPayload(serverUpdateUserBody string, serverUpdateUserToken string) (*server.UpdateUserPayload, error) {
	var err error
	var body UpdateUserRequestBody
	{
		err = json.Unmarshal([]byte(serverUpdateUserBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"Iure aliquid deserunt enim.\",\n      \"name\": \"Molestiae nostrum.\"\n   }'")
		}
	}
	var token string
	{
		token = serverUpdateUserToken
	}
	v := &server.UpdateUserPayload{
		Name:  body.Name,
		Email: body.Email,
	}
	v.Token = token

	return v, nil
}

// BuildMemoriesPayload builds the payload for the server memories endpoint
// from CLI flags.
func BuildMemoriesPayload(serverMemoriesToken string) (*server.MemoriesPayload, error) {
	var token string
	{
		token = serverMemoriesToken
	}
	v := &server.MemoriesPayload{}
	v.Token = token

	return v, nil
}

// BuildMemoryByIDPayload builds the payload for the server memoryById endpoint
// from CLI flags.
func BuildMemoryByIDPayload(serverMemoryByIDID string, serverMemoryByIDToken string) (*server.MemoryByIDPayload, error) {
	var err error
	var id int
	{
		var v int64
		v, err = strconv.ParseInt(serverMemoryByIDID, 10, strconv.IntSize)
		id = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT")
		}
	}
	var token string
	{
		token = serverMemoryByIDToken
	}
	v := &server.MemoryByIDPayload{}
	v.ID = id
	v.Token = token

	return v, nil
}

// BuildNewMemoryPayload builds the payload for the server newMemory endpoint
// from CLI flags.
func BuildNewMemoryPayload(serverNewMemoryBody string, serverNewMemoryToken string) (*server.NewMemoryPayload, error) {
	var err error
	var body NewMemoryRequestBody
	{
		err = json.Unmarshal([]byte(serverNewMemoryBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"date\": \"2004-11-09\",\n      \"detail\": \"Reiciendis suscipit possimus eligendi.\",\n      \"location\": \"Aspernatur repellendus veritatis.\",\n      \"title\": \"Et et.\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.date", body.Date, goa.FormatDate))
		if err != nil {
			return nil, err
		}
	}
	var token string
	{
		token = serverNewMemoryToken
	}
	v := &server.NewMemoryPayload{
		Title:    body.Title,
		Date:     body.Date,
		Location: body.Location,
		Detail:   body.Detail,
	}
	v.Token = token

	return v, nil
}

// BuildDeleteMemoryPayload builds the payload for the server deleteMemory
// endpoint from CLI flags.
func BuildDeleteMemoryPayload(serverDeleteMemoryID string, serverDeleteMemoryToken string) (*server.DeleteMemoryPayload, error) {
	var err error
	var id int
	{
		var v int64
		v, err = strconv.ParseInt(serverDeleteMemoryID, 10, strconv.IntSize)
		id = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT")
		}
	}
	var token string
	{
		token = serverDeleteMemoryToken
	}
	v := &server.DeleteMemoryPayload{}
	v.ID = id
	v.Token = token

	return v, nil
}

// BuildUpdateMemoryPayload builds the payload for the server updateMemory
// endpoint from CLI flags.
func BuildUpdateMemoryPayload(serverUpdateMemoryBody string, serverUpdateMemoryID string, serverUpdateMemoryToken string) (*server.UpdateMemoryPayload, error) {
	var err error
	var body UpdateMemoryRequestBody
	{
		err = json.Unmarshal([]byte(serverUpdateMemoryBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"date\": \"1974-01-04\",\n      \"detail\": \"Nihil officia officiis dignissimos magni ea sed.\",\n      \"location\": \"Harum sapiente dolor ut provident pariatur non.\",\n      \"title\": \"Veniam id neque nam sunt eos vero.\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.date", body.Date, goa.FormatDate))
		if err != nil {
			return nil, err
		}
	}
	var id int
	{
		var v int64
		v, err = strconv.ParseInt(serverUpdateMemoryID, 10, strconv.IntSize)
		id = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT")
		}
	}
	var token string
	{
		token = serverUpdateMemoryToken
	}
	v := &server.UpdateMemoryPayload{
		Title:    body.Title,
		Date:     body.Date,
		Location: body.Location,
		Detail:   body.Detail,
	}
	v.ID = id
	v.Token = token

	return v, nil
}

// BuildNewHeavyPayload builds the payload for the server newHeavy endpoint
// from CLI flags.
func BuildNewHeavyPayload(serverNewHeavyToken string) (*server.NewHeavyPayload, error) {
	var token string
	{
		token = serverNewHeavyToken
	}
	v := &server.NewHeavyPayload{}
	v.Token = token

	return v, nil
}

// BuildCheckHeavyPayload builds the payload for the server checkHeavy endpoint
// from CLI flags.
func BuildCheckHeavyPayload(serverCheckHeavyKey string, serverCheckHeavyToken string) (*server.CheckHeavyPayload, error) {
	var key string
	{
		key = serverCheckHeavyKey
	}
	var token string
	{
		token = serverCheckHeavyToken
	}
	v := &server.CheckHeavyPayload{}
	v.Key = key
	v.Token = token

	return v, nil
}

// BuildDeleteHeavyPayload builds the payload for the server deleteHeavy
// endpoint from CLI flags.
func BuildDeleteHeavyPayload(serverDeleteHeavyKey string, serverDeleteHeavyToken string) (*server.DeleteHeavyPayload, error) {
	var key string
	{
		key = serverDeleteHeavyKey
	}
	var token string
	{
		token = serverDeleteHeavyToken
	}
	v := &server.DeleteHeavyPayload{}
	v.Key = key
	v.Token = token

	return v, nil
}
