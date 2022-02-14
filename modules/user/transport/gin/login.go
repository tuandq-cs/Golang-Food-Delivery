package usergin

import (
	"Golang_Edu/common"
	"Golang_Edu/component/appctx"
	"Golang_Edu/component/hasher"
	"Golang_Edu/component/tokenprovider/jwt"
	userbusiness "Golang_Edu/modules/user/business"
	usermodel "Golang_Edu/modules/user/model"
	userstorage "Golang_Edu/modules/user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var user usermodel.UserLogin
		if err := ctx.ShouldBind(&user); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		tokenProvider := jwt.NewJWTTokenProvider(appCtx.SecretKey())
		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		md5 := hasher.NewMd5Hasher()
		biz := userbusiness.NewLoginBiz(store, md5, tokenProvider, 60*60*24*30)
		accessToken, err := biz.Login(ctx, &user)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(accessToken))
	}
}
