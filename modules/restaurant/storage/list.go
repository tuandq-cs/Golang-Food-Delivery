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
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	db := store.db
	var listData []restaurantmodel.Restaurant
	// Conditions
	if filter.UserId > 0 {
		db = db.Where("owner_id = ?", filter.UserId)
	}
	db = db.Where("status not in (0)")
	// Total records
	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	// Preload fields
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}
	// Get list records
	// Speed up query if cursor is provided
	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)
		if err != nil {
			return nil, common.ErrDB(err)
		}
		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		db = db.Offset(paging.Offset())
	}
	if err := db.Limit(paging.Limit).
		Order("id desc").
		Find(&listData).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return listData, nil
}
