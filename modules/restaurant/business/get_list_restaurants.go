package restaurantbusiness

import (
	"Golang_Edu/common"
	restaurantmodel "Golang_Edu/modules/restaurant/model"
	"context"
)

type GetListRestaurantStore interface {
	ListDataWithConditions(
		context context.Context,
		paging *common.Paging,
		filter *restaurantmodel.Filter,
	) ([]restaurantmodel.Restaurant, error)
}

type getListRestaurantsBiz struct {
	store GetListRestaurantStore
}

func NewGetListRestaurantsBiz(store GetListRestaurantStore) *getListRestaurantsBiz {
	return &getListRestaurantsBiz{store: store}
}

func (biz *getListRestaurantsBiz) GetListRestaurants(context context.Context,
	paging *common.Paging,
	filter *restaurantmodel.Filter,
) ([]restaurantmodel.Restaurant, error) {
	listData, err := biz.store.ListDataWithConditions(context, paging, filter)
	if err != nil {
		return nil, err
	}
	return listData, nil
}
