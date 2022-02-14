package userbusiness

import (
	"Golang_Edu/common"
	"Golang_Edu/component/hasher"
	"Golang_Edu/component/tokenprovider"
	usermodel "Golang_Edu/modules/user/model"
	"context"
)

type LoginStore interface {
	FindUser(context context.Context,
		conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginBiz struct {
	store         LoginStore
	loginHasher   hasher.Hasher
	tokenProvider tokenprovider.TokenProvider
	expiry        int
}

func NewLoginBiz(store LoginStore, loginHasher hasher.Hasher,
	provider tokenprovider.TokenProvider, expiry int) *loginBiz {
	return &loginBiz{
		store:         store,
		loginHasher:   loginHasher,
		tokenProvider: provider,
		expiry:        expiry,
	}
}

func (biz *loginBiz) Login(context context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	if err := data.Validate(); err != nil {
		return nil, common.ErrInvalidRequest(err)
	}
	// Find user
	user, err := biz.store.FindUser(context, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}
	// Check password
	if biz.loginHasher.Hash(data.Password+user.Salt) != user.Password {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.GetUserId(),
		Role:   user.GetRole(),
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}
