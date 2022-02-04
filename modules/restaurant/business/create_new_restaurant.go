package restaurantbusiness

import (
	restaurantmodel "Golang_Edu/modules/restaurant/model"
	"context"
)

type CreateRestaurantStore interface {
	Insert(context context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createNewRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createNewRestaurantBiz {
	return &createNewRestaurantBiz{store: store}
}

func (biz *createNewRestaurantBiz) CreateNewRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}
	if insertError := biz.store.Insert(context, data); insertError != nil {
		return insertError
	}
	return nil
}
