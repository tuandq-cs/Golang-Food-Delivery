package restaurantstorage

import (
	"Golang_Edu/common"
	restaurantmodel "Golang_Edu/modules/restaurant/model"
	"context"
	"gorm.io/gorm"
)

func (store *sqlStore) FindDataWithConditions(
	context context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	var data restaurantmodel.Restaurant
	if err := store.db.First(&data, conditions).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrDataNotFound
		}
		return nil, err
	}
	return &data, nil
}
