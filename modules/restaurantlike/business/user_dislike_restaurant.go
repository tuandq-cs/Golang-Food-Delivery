package restaurantlikebusiness

import (
	restaurantlikemodel "Golang_Edu/modules/restaurantlike/model"
	"context"
)

type userDislikeRestaurantStore interface {
	FindDataWithConditions(
		context context.Context,
		condition map[string]interface{},
	) (*restaurantlikemodel.Like, error)
	Delete(context context.Context, userId int, restaurantId int) error
}

type userDislikeRestaurant struct {
	store userDislikeRestaurantStore
}

func NewUserDislikeRestaurantBiz(store userDislikeRestaurantStore) *userDislikeRestaurant {
	return &userDislikeRestaurant{
		store: store,
	}
}

// DislikeRestaurant
// 1. Check if user has been like this restaurant
// 2. If yes, delete record in restaurant_likes table
func (biz *userDislikeRestaurant) DislikeRestaurant(
	context context.Context,
	data *restaurantlikemodel.Like,
) error {
	// 1. Check if user has been like this restaurant
	oldData, err := biz.store.FindDataWithConditions(
		context,
		map[string]interface{}{"user_id": data.UserId, "restaurant_id": data.RestaurantId},
	)
	if oldData == nil {
		return restaurantlikemodel.ErrCannotDislikeRestaurant(err)
	}
	// 2. If yes, delete record in restaurant_likes table
	if err := biz.store.Delete(context, data.UserId, data.RestaurantId); err != nil {
		return restaurantlikemodel.ErrCannotDislikeRestaurant(err)
	}
	return nil
}
