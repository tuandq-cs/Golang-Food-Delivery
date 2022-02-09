package restaurantgin

import (
	"Golang_Edu/common"
	"Golang_Edu/component/appctx"
	restaurantbusiness "Golang_Edu/modules/restaurant/business"
	restaurantmodel "Golang_Edu/modules/restaurant/model"
	restaurantstorage "Golang_Edu/modules/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateRestaurant(appCtx appctx.AppContext) func(*gin.Context) {
	return func(context *gin.Context) {
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var updateData restaurantmodel.RestaurantUpdate
		if err := context.ShouldBind(&updateData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		updateRestaurantBiz := restaurantbusiness.NewUpdateRestaurantBiz(store)
		if err := updateRestaurantBiz.UpdateRestaurant(context.Request.Context(), &updateData, int(uid.GetLocalID())); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
