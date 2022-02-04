package restaurantstorage

import (
	"Golang_Edu/common"
	restaurantmodel "Golang_Edu/modules/restaurant/model"
	"context"
)

func (store *sqlStore) ListDataWithConditions(
	context context.Context,
	paging *common.Paging,
	filter *restaurantmodel.Filter,
) ([]restaurantmodel.Restaurant, error) {
	db := store.db
	var listData []restaurantmodel.Restaurant
	// Conditions
	if filter.UserId > 0 {
		db = db.Where("owner_id = ?", filter.UserId)
	}
	db = db.Where("status not in (0)")
	// Count records
	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	// Get list records
	if err := db.Limit(paging.Limit).
		Offset(paging.Offset()).
		Order("id desc").
		Find(&listData).Error; err != nil {
		return nil, err
	}
	return listData, nil
}
