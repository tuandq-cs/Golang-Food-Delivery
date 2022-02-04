package restaurantmodel

import (
	"Golang_Edu/common"
	"strings"
)

type RestaurantCreate struct {
	common.SQLModel
	Name    string `json:"name" form:"name" gorm:"column:name;"`
	Address string `json:"address" form:"address" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return ErrNameCannotBeBlank
	}

	data.Address = strings.TrimSpace(data.Address)
	if data.Address == "" {
		return ErrAddressCannotBeBlank
	}
	return nil
}
