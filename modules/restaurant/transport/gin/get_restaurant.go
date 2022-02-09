package restaurantgin

import (
	"Golang_Edu/common"
	"Golang_Edu/component/appctx"
	restaurantbusiness "Golang_Edu/modules/restaurant/business"
	restaurantstorage "Golang_Edu/modules/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRestaurant(appCtx appctx.AppContext) func(*gin.Context) {
	return func(context *gin.Context) {
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(err)
		}
		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		getRestaurantBiz := restaurantbusiness.NewGetRestaurantBiz(store)
		data, err := getRestaurantBiz.GetRestaurant(context.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}
		data.Mask(common.DbTypeRestaurant)
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
