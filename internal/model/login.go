package model

import (
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l Login) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Email, validation.Required, is.Email),
		validation.Field(&l.Password, validation.Required, validation.Length(6, 50)),
	)
}
