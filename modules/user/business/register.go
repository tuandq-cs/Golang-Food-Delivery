package userbusiness

import (
	"Golang_Edu/common"
	"Golang_Edu/component/hasher"
	usermodel "Golang_Edu/modules/user/model"
	"context"
)

type RegisterStore interface {
	FindUser(context context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	InsertUser(context context.Context, data *usermodel.UserCreate) error
}

type registerBiz struct {
	store          RegisterStore
	registerHasher hasher.Hasher
}

func NewRegisterBiz(store RegisterStore, registerHasher hasher.Hasher) *registerBiz {
	return &registerBiz{store: store, registerHasher: registerHasher}
}

func (biz *registerBiz) Register(context context.Context, data *usermodel.UserCreate) error {
	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}
	user, _ := biz.store.FindUser(context, map[string]interface{}{"email": data.Email})
	if user != nil {
		return usermodel.ErrEmailExisted
	}

	salt := common.GenSalt(50)
	data.Password = biz.registerHasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user" // Hard code

	if err := biz.store.InsertUser(context, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	return nil
}
