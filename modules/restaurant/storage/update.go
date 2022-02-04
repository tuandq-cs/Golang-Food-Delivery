package restaurantstorage

import (
	restaurantmodel "Golang_Edu/modules/restaurant/model"
	"context"
)

func (store *sqlStore) UpdateDataWithConditions(
	context context.Context,
	updateData *restaurantmodel.RestaurantUpdate,
	conditions map[string]interface{},
) error {
	updateResult := store.db.Where(conditions).Updates(updateData)
	if updateResult.Error != nil {
		return updateResult.Error
	}
	return nil
}
