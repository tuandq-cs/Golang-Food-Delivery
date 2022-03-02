package restaurantmodel

import (
	"Golang_Edu/common"
	"errors"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel
	Name       string             `json:"name" gorm:"column:name;"`
	Address    string             `json:"address" gorm:"column:addr;"`
	OwnerId    int                `json:"-" gorm:"owner_id"`
	LikedCount int                `json:"like_count" gorm:"liked_count"`
	HasLiked   bool               `json:"has_liked"`
	Logo       *common.Image      `json:"logo" gorm:"column:logo;"`
	Cover      *common.Images     `json:"cover" gorm:"column:cover;"`
	User       *common.SimpleUser `json:"user" gorm:"PRELOAD:false; foreignKey:OwnerId;"`
}

func (Restaurant) TableName() string { return "restaurants" }

func (data *Restaurant) Mask(dbType int) {
	data.SQLModel.Mask(dbType)

	if u := data.User; u != nil {
		u.Mask(common.DbTypeUser)
	}
}

var (
	ErrNameCannotBeBlank     = errors.New("restaurant name can not be blank")
	ErrAddressCannotBeBlank  = errors.New("restaurant address can not be blank")
	ErrResHasBeenInactivated = errors.New("restaurant has been inactivated")
	ErrResHasBeenActivated   = errors.New("restaurant has been activated")
)
