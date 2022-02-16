package restaurantlikestorage

import (
	"Golang_Edu/common"
	restaurantlikemodel "Golang_Edu/modules/restaurantlike/model"
	"context"
	"gorm.io/gorm"
)

func (store *sqlStore) FindDataWithConditions(
	context context.Context,
	condition map[string]interface{},
) (*restaurantlikemodel.Like, error) {
	db := store.db
	var data restaurantlikemodel.Like
	if err := db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrDataNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
