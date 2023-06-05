package model

import (
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Active   int    `json:"active"`
	Email    string `json:"email,omitempty"`
	Tag      string `json:"tag,omitempty"`
	Password string `json:"password,omitempty"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.Required, validation.Length(10, 50)),
		validation.Field(&u.Tag, validation.Required, validation.Length(6, 50)),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 50)),
		validation.Field(&u.Email, validation.Required, is.Email),
	)
}

type Users []*User
