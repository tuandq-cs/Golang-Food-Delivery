package usermodel

import (
	"Golang_Edu/common"
	"errors"
)

const (
	EntityName = "User"
)

type User struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"-" gorm:"column:password;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	Phone           string        `json:"phone" gorm:"column:phone;"`
	Role            string        `json:"role" gorm:"column:role"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar"`
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

func (User) TableName() string {
	return "users"
}

var (
	ErrEmailOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)

	ErrEmailCannotBeBlank = common.NewCustomError(
		errors.New("email can not be blank"),
		"email has already existed",
		"ErrEmailExisted",
	)

	ErrPasswordCannotBeBlank = common.NewCustomError(
		errors.New("password can not be blank"),
		"email has already existed",
		"ErrEmailExisted",
	)
)
