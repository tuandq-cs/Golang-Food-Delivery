package usermodel

import "strings"

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password;"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

func (user *UserLogin) Validate() error {
	if strings.TrimSpace(user.Email) == "" {
		return ErrEmailCannotBeBlank
	}

	if strings.TrimSpace(user.Password) == "" {
		return ErrPasswordCannotBeBlank
	}
	return nil
}
