package restaurantbusiness

import (
	"Golang_Edu/common"
	restaurantmodel "Golang_Edu/modules/restaurant/model"
	"context"
)

type GetListRestaurantsRepository interface {
	GetListRestaurants(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type getListRestaurantsBiz struct {
	repo GetListRestaurantsRepository
}

func NewGetListRestaurantsBiz(repo GetListRestaurantsRepository) *getListRestaurantsBiz {
	return &getListRestaurantsBiz{repo: repo}
}

func (biz *getListRestaurantsBiz) GetListRestaurants(context context.Context,
	paging *common.Paging,
	filter *restaurantmodel.Filter,
) ([]restaurantmodel.Restaurant, error) {
	listData, err := biz.repo.GetListRestaurants(context, filter, paging, "User")
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	return listData, nil
}
