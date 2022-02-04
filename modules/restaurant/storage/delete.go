package restaurantstorage

import (
	restaurantmodel "Golang_Edu/modules/restaurant/model"
	"context"
)

func (store *sqlStore) Delete(context context.Context, id int) error {
	db := store.db
	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
