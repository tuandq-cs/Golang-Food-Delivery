package restaurantgin

import (
	"Golang_Edu/common"
	"Golang_Edu/component/appctx"
	restaurantbusiness "Golang_Edu/modules/restaurant/business"
	restaurantstorage "Golang_Edu/modules/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteRestaurant(appCtx appctx.AppContext) func(*gin.Context) {
	return func(context *gin.Context) {
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(err)
		}
		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbusiness.NewDeleteRestaurantBiz(store)
		if err := biz.DeleteRestaurant(context.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}
		context.JSON(http.StatusInternalServerError, common.SimpleSuccessResponse(true))
	}
}
