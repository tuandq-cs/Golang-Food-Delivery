package restaurantlikebusiness

import (
	"Golang_Edu/common"
	restaurantlikemodel "Golang_Edu/modules/restaurantlike/model"
	"context"
)

type userLikeRestaurantStore interface {
	FindDataWithConditions(
		context context.Context,
		condition map[string]interface{},
	) (*restaurantlikemodel.Like, error)
	Insert(context context.Context, newData *restaurantlikemodel.Like) error
}

type userLikeRestaurantBiz struct {
	store userLikeRestaurantStore
}

func NewUserLikeRestaurantBiz(store userLikeRestaurantStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store: store,
	}
}

// LikeRestaurant
// 1. Check if user has been like this restaurant
// 2. If not, create record in restaurant_likes table
func (biz *userLikeRestaurantBiz) LikeRestaurant(
	context context.Context,
	newData *restaurantlikemodel.Like,
) error {
	// 1. Check if user has been like this restaurant
	oldData, err := biz.store.FindDataWithConditions(
		context,
		map[string]interface{}{"user_id": newData.UserId, "restaurant_id": newData.RestaurantId},
	)
	// If err == ErrDataNotFound -> Continue
	if err != nil && err != common.ErrDataNotFound {
		return err
	}
	if oldData != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(common.ErrEntityExisted(
			restaurantlikemodel.EntityName,
			nil,
		))
	}
	// 2. If not, create record in restaurant_likes table
	if err := biz.store.Insert(context, newData); err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}
	return nil
}
