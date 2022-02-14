package usermodel

import (
	"Golang_Edu/common"
	"strings"
)

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"password" gorm:"column:password;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	Role            string        `json:"-" gorm:"column:role;`
	Salt            string        `json:"-" gorm:"column:salt;"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (user *UserCreate) Validate() error {
	if strings.TrimSpace(user.Email) == "" {
		return ErrEmailCannotBeBlank
	}

	if strings.TrimSpace(user.Password) == "" {
		return ErrPasswordCannotBeBlank
	}
	return nil
}
