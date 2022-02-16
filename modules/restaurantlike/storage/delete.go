package restaurantlikestorage

import (
	"Golang_Edu/common"
	restaurantlikemodel "Golang_Edu/modules/restaurantlike/model"
	"context"
)

func (store *sqlStore) Delete(context context.Context, userId int, restaurantId int) error {
	db := store.db
	if err := db.Table(restaurantlikemodel.Like{}.TableName()).
		Where("user_id = ? AND restaurant_id = ?", userId, restaurantId).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
