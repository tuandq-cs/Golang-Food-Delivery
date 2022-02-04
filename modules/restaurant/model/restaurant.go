package restaurantmodel

import (
	"Golang_Edu/common"
	"errors"
)

type Restaurant struct {
	common.SQLModel
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:addr;"`
}

func (Restaurant) TableName() string { return "restaurants" }

var (
	ErrNameCannotBeBlank     = errors.New("restaurant name can not be blank")
	ErrAddressCannotBeBlank  = errors.New("restaurant address can not be blank")
	ErrResHasBeenInactivated = errors.New("restaurant has been inactivated")
	ErrResHasBeenActivated   = errors.New("restaurant has been activated")
)