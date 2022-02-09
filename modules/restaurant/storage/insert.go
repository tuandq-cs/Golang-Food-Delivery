package restaurantstorage

import (
	"Golang_Edu/common"
	restaurantmodel "Golang_Edu/modules/restaurant/model"
	"context"
)

func (store *sqlStore) Insert(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	result := store.db.Create(&data)
	if result.Error != nil {
		return common.ErrDB(result.Error)
	}
	return nil
}
