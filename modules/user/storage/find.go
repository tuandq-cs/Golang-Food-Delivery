package userstorage

import (
	"Golang_Edu/common"
	usermodel "Golang_Edu/modules/user/model"
	"context"
	"gorm.io/gorm"
)

func (store *sqlStore) FindUser(context context.Context,
	conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	db := store.db.Table(usermodel.User{}.TableName())
	var user usermodel.User
	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &user, nil
}
