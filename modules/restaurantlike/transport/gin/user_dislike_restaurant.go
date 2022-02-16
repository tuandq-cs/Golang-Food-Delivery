package restaurantlikegin

import (
	"Golang_Edu/common"
	"Golang_Edu/component/appctx"
	restaurantlikebusiness "Golang_Edu/modules/restaurantlike/business"
	restaurantlikemodel "Golang_Edu/modules/restaurantlike/model"
	restaurantlikestorage "Golang_Edu/modules/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserDislikeRestaurant(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		u := ctx.MustGet(common.CurrentUser).(common.Requester)
		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       u.GetUserId(),
		}
		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantlikebusiness.NewUserDislikeRestaurantBiz(store)
		if err := biz.DislikeRestaurant(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}