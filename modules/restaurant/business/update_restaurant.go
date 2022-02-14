package restaurantbusiness

import (
	"Golang_Edu/common"
	restaurantmodel "Golang_Edu/modules/restaurant/model"
	"context"
	"errors"
)

type UpdateRestaurantStore interface {
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

type UpdateRestaurantBiz struct {
	store     UpdateRestaurantStore
	requester common.Requester
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore, requester common.Requester) *UpdateRestaurantBiz {
	return &UpdateRestaurantBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *UpdateRestaurantBiz) UpdateRestaurant(
	context context.Context,
	updateData *restaurantmodel.RestaurantUpdate,
	id int,
) error {
	// Validation
	if err := updateData.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}
	// Check if restaurant is existed
	oldData, err := biz.store.FindDataWithConditions(context, map[string]interface{}{"id": id})
	if err != nil {
		if err == common.ErrDataNotFound {
			return common.ErrEntityNotFound(restaurantmodel.EntityName, err)
		}
		return err
	}
	if biz.requester.GetRole() != "admin" && biz.requester.GetUserId() != oldData.OwnerId {
		return common.ErrNoPermission(nil)
	}
	// Check if data has been deleted or blocked
	if oldData.Status == 0 {
		return errors.New("data has been deleted or blocked")
	}
	// Update restaurant
	if err := biz.store.UpdateDataWithConditions(context, updateData, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotUpdateEntity(restaurantmodel.EntityName, err)
	}
	return nil
}
