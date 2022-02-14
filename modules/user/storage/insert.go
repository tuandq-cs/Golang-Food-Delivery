package userstorage

import (
	"Golang_Edu/common"
	usermodel "Golang_Edu/modules/user/model"
	"context"
)

func (store *sqlStore) InsertUser(context context.Context, data *usermodel.UserCreate) error {
	db := store.db.Begin()

	if err := db.Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
