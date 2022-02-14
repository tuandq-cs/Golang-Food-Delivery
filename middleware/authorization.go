package middleware

import (
	"Golang_Edu/common"
	"Golang_Edu/component/appctx"
	"Golang_Edu/component/tokenprovider/jwt"
	usermodel "Golang_Edu/modules/user/model"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type AuthStore interface {
	FindUser(context context.Context,
		conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

// 1. Extract Authentication in header and get token
// 2. Validate token and parse Token payload
// 3. Check if user has been blocked ??

func RequiredAuth(appCtx appctx.AppContext, authStore AuthStore) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 1. Extract Authentication in header and get token
		token, err := extractTokenFromHeaderString(ctx.GetHeader("Authorization"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		// 2. Validate token and parse Token payload
		jwtProvider := jwt.NewJWTTokenProvider(appCtx.SecretKey())
		payload, err := jwtProvider.Validate(token)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		// 3. Check if user has been blocked
		user, err := authStore.FindUser(ctx.Request.Context(), map[string]interface{}{"id": payload.UserId})
		if err != nil {
			panic(common.ErrDB(err))
		}
		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}
		user.Mask(common.DbTypeUser)
		ctx.Set(common.CurrentUser, user)
		ctx.Next()
	}
}

func extractTokenFromHeaderString(header string) (string, error) {
	parts := strings.Split(header, " ")
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}
	return parts[1], nil
}
