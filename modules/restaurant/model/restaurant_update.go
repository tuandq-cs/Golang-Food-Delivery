package restaurantmodel

import (
	"Golang_Edu/common"
	"strings"
)

type RestaurantUpdate struct {
	common.SQLModel
	Name    *string `json:"name" form:"name" gorm:"column:name;"`
	Address *string `json:"address" form:"address" gorm:"column:addr;"`
	Status  *int    `json:"-" gorm:"column:status;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

func (data *RestaurantUpdate) Validate() error {
	if data.Name != nil {
		str := strings.TrimSpace(*(data.Name))
		if str == "" {
			return ErrNameCannotBeBlank
		}
		data.Name = &str

	}
	if data.Address != nil {
		str := strings.TrimSpace(*(data.Address))
		if str == "" {
			return ErrAddressCannotBeBlank
		}
		data.Address = &str
	}
	return nil
}
