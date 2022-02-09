package restaurantgin

import (
	"Golang_Edu/common"
	"Golang_Edu/component/appctx"
	restaurantbusiness "Golang_Edu/modules/restaurant/business"
	restaurantstorage "Golang_Edu/modules/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ActivateRestaurant(appCtx appctx.AppContext) func(*gin.Context) {
	return func(context *gin.Context) {
		// Convert id
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(err)
		}
		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbusiness.NewActivateRestaurantBiz(store)
		if err := biz.ActivateRestaurant(context, int(uid.GetLocalID())); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
