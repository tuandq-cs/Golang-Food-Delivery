package restaurantlikestorage

import (
	"Golang_Edu/common"
	restaurantlikemodel "Golang_Edu/modules/restaurantlike/model"
	"context"
)

func (store *sqlStore) Insert(context context.Context, newData *restaurantlikemodel.Like) error {
	db := store.db
	if err := db.Create(newData).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
