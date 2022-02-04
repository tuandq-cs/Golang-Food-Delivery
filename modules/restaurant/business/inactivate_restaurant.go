package restaurantbusiness

import (
	"Golang_Edu/common"
	restaurantmodel "Golang_Edu/modules/restaurant/model"
	"context"
)

type InactivateRestaurantStore interface {
	FindDataWithConditions(
		context context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	UpdateDataWithConditions(
		context context.Context,
		updateData *restaurantmodel.RestaurantUpdate,
		conditions map[string]interface{},
	) error
}

type inactivateRestaurantBiz struct {
	store InactivateRestaurantStore
}

func NewInactivateRestaurantBiz(store InactivateRestaurantStore) *inactivateRestaurantBiz {
	return &inactivateRestaurantBiz{store: store}
}

func (biz *inactivateRestaurantBiz) InactivateRestaurant(context context.Context, id int) error {
	// Check if restaurant is existed
	oldData, err := biz.store.FindDataWithConditions(context, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	//
	if oldData.Status == common.InactiveStatus {
		return restaurantmodel.ErrResHasBeenInactivated
	}
	inactiveStatus := common.InactiveStatus
	if err := biz.store.UpdateDataWithConditions(
		context,
		&restaurantmodel.RestaurantUpdate{Status: &inactiveStatus},
		map[string]interface{}{"id": id},
	); err != nil {
		return err
	}
	return nil
}
