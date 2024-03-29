// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"time"
)

type DeleteMemory struct {
	ID int `json:"id"`
}

type Memory struct {
	ID       int       `json:"id"`
	UserID   int       `json:"userId"`
	Title    string    `json:"title"`
	Date     time.Time `json:"date"`
	Location *string   `json:"location,omitempty"`
	Detail   *string   `json:"detail,omitempty"`
}

type Mutation struct {
}

type NewMemory struct {
	Title    string    `json:"title"`
	Date     time.Time `json:"date"`
	Location *string   `json:"location,omitempty"`
	Detail   *string   `json:"detail,omitempty"`
}

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Query struct {
}

type UpdateMemory struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Date     time.Time `json:"date"`
	Location *string   `json:"location,omitempty"`
	Detail   *string   `json:"detail,omitempty"`
}

type UpdateUser struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
