package restaurantlikebusiness

import (
	"Golang_Edu/common"
	restaurantlikemodel "Golang_Edu/modules/restaurantlike/model"
	"context"
)

type listUsersLikeRestaurantStore interface {
	GetUsersLikeRestaurant(
		context context.Context,
		conditions map[string]interface{},
		filter *restaurantlikemodel.Filter,
		paging *common.Paging,
	) ([]common.SimpleUser, error)
}

type listUsersLikeRestaurantBiz struct {
	store listUsersLikeRestaurantStore
}

func NewListUsersLikeRestaurantBiz(store listUsersLikeRestaurantStore) *listUsersLikeRestaurantBiz {
	return &listUsersLikeRestaurantBiz{
		store: store,
	}
}

func (biz *listUsersLikeRestaurantBiz) ListUsersLikeRestaurant(
	context context.Context,
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
) (
	[]common.SimpleUser,
	error,
) {
	listUsers, err := biz.store.GetUsersLikeRestaurant(context, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantlikemodel.EntityName, err)
	}
	return listUsers, nil
}
