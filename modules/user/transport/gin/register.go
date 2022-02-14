package usergin

import (
	"Golang_Edu/common"
	"Golang_Edu/component/appctx"
	"Golang_Edu/component/hasher"
	userbusiness "Golang_Edu/modules/user/business"
	usermodel "Golang_Edu/modules/user/model"
	userstorage "Golang_Edu/modules/user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var user *usermodel.UserCreate = &usermodel.UserCreate{}
		if err := ctx.ShouldBind(user); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		md5 := hasher.NewMd5Hasher()
		biz := userbusiness.NewRegisterBiz(store, md5)
		if err := biz.Register(ctx, user); err != nil {
			panic(err)
		}
		user.Mask(common.DbTypeUser)
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(user.FakeId.String()))
	}
}
