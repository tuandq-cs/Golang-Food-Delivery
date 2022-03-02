package restaurantstorage

import (
	"Golang_Edu/common"
	restaurantmodel "Golang_Edu/modules/restaurant/model"
	"context"
	"gorm.io/gorm"
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

func (store *sqlStore) IncreasedLikeCounts(
	context context.Context,
	restaurantId int,
) error {
	db := store.db
	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", restaurantId).
		Update("liked_count", gorm.Expr("liked_count + 1")).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (store *sqlStore) DecreasedLikeCounts(
	context context.Context,
	restaurantId int,
) error {
	db := store.db
	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", restaurantId).
		Update("liked_count", gorm.Expr("liked_count - 1")).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
