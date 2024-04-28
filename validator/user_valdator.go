package validator

import (
	"go-rest-api/model"

	valdation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValdator interface {
	UserValdate(task model.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValdator {
	return &userValidator{}
}

func (uv *userValidator) UserValdate(user model.User) error {
	return valdation.ValidateStruct(&user,
		valdation.Field(
			&user.Email, 
			valdation.Required.Error("email is required"),
			valdation.RuneLength(1, 30).Error("limited max 30 characters"),
			is.Email.Error("is not email format"),
		),
		valdation.Field(
			&user.Password,
			valdation.Required.Error("password is required"),
			valdation.RuneLength(6, 30).Error("limited min 6 max 30 characters"),
		),
	)
}
