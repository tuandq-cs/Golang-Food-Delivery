package restaurantlikebusiness

import (
	"Golang_Edu/common"
	"Golang_Edu/component/asyncjob"
	restaurantlikemodel "Golang_Edu/modules/restaurantlike/model"
	"context"
	"log"
)

type userDislikeRestaurantStore interface {
	FindDataWithConditions(
		context context.Context,
		condition map[string]interface{},
	) (*restaurantlikemodel.Like, error)
	Delete(context context.Context, userId int, restaurantId int) error
}

type desLikedCountStore interface {
	DecreasedLikeCounts(
		context context.Context,
		restaurantId int,
	) error
}

type userDislikeRestaurant struct {
	store        userDislikeRestaurantStore
	desLikeStore desLikedCountStore
}

func NewUserDislikeRestaurantBiz(store userDislikeRestaurantStore, desLikeStore desLikedCountStore) *userDislikeRestaurant {
	return &userDislikeRestaurant{
		store:        store,
		desLikeStore: desLikeStore,
	}
}

// DislikeRestaurant
// 1. Check if user has been like this restaurant
// 2. If yes, delete record in restaurant_likes table
func (biz *userDislikeRestaurant) DislikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	// 1. Check if user has been like this restaurant
	oldData, err := biz.store.FindDataWithConditions(
		ctx,
		map[string]interface{}{"user_id": data.UserId, "restaurant_id": data.RestaurantId},
	)
	if oldData == nil {
		return restaurantlikemodel.ErrCannotDislikeRestaurant(err)
	}
	// 2. If yes, delete record in restaurant_likes table
	if err := biz.store.Delete(ctx, data.UserId, data.RestaurantId); err != nil {
		return restaurantlikemodel.ErrCannotDislikeRestaurant(err)
	}
	// Side effect
	go func() {
		defer common.Recovery()
		job := asyncjob.NewJob(func(ctx context.Context) error {
			if err := biz.desLikeStore.DecreasedLikeCounts(ctx, data.RestaurantId); err != nil {
				log.Println(err)
			}
			return nil
		}, asyncjob.WithName("DecreasedLikeCounts"))
		group := asyncjob.NewGroup(false, job)
		if err := group.Run(ctx); err != nil {
			log.Println(err)
		}
	}()
	return nil
}
