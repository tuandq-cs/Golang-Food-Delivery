package restaurantlikebusiness

import (
	"Golang_Edu/common"
	"Golang_Edu/component/asyncjob"
	restaurantlikemodel "Golang_Edu/modules/restaurantlike/model"
	"context"
	"log"
)

type userLikeRestaurantStore interface {
	FindDataWithConditions(
		context context.Context,
		condition map[string]interface{},
	) (*restaurantlikemodel.Like, error)
	Insert(context context.Context, newData *restaurantlikemodel.Like) error
}

type incLikedCountResStore interface {
	IncreasedLikeCounts(
		context context.Context,
		restaurantId int,
	) error
}

type userLikeRestaurantBiz struct {
	store        userLikeRestaurantStore
	incLikeStore incLikedCountResStore
}

func NewUserLikeRestaurantBiz(store userLikeRestaurantStore, incLikeStore incLikedCountResStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store:        store,
		incLikeStore: incLikeStore,
	}
}

// LikeRestaurant
// 1. Check if user has been like this restaurant
// 2. If not, create record in restaurant_likes table
func (biz *userLikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	newData *restaurantlikemodel.Like,
) error {
	// 1. Check if user has been like this restaurant
	oldData, err := biz.store.FindDataWithConditions(
		ctx,
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
	if err := biz.store.Insert(ctx, newData); err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	// Side effect
	go func() {
		defer common.Recovery()
		job := asyncjob.NewJob(func(ctx context.Context) error {
			if err := biz.incLikeStore.IncreasedLikeCounts(ctx, newData.RestaurantId); err != nil {
				log.Println(err)
			}
			return nil
		}, asyncjob.WithName("IncreasedLikeCounts"))
		group := asyncjob.NewGroup(false, job)
		if err := group.Run(ctx); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
