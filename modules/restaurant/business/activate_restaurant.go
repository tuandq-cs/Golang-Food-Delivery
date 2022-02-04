package restaurantbusiness

import (
	"Golang_Edu/common"
	restaurantmodel "Golang_Edu/modules/restaurant/model"
	"context"
)

type ActivateRestaurantStore interface {
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

type activateRestaurantBiz struct {
	store ActivateRestaurantStore
}

func NewActivateRestaurantBiz(store ActivateRestaurantStore) *activateRestaurantBiz {
	return &activateRestaurantBiz{store: store}
}

func (biz *activateRestaurantBiz) ActivateRestaurant(context context.Context, id int) error {
	// Check if restaurant is existed
	oldData, err := biz.store.FindDataWithConditions(context, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	//
	if oldData.Status == common.ActiveStatus {
		return restaurantmodel.ErrResHasBeenActivated
	}
	activeStatus := common.ActiveStatus
	if err := biz.store.UpdateDataWithConditions(
		context,
		&restaurantmodel.RestaurantUpdate{Status: &activeStatus},
		map[string]interface{}{"id": id},
	); err != nil {
		return err
	}
	return nil
}
