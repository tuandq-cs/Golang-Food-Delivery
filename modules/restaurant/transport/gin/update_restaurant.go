package restaurantgin

import (
	"Golang_Edu/component/appctx"
	restaurantbusiness "Golang_Edu/modules/restaurant/business"
	restaurantmodel "Golang_Edu/modules/restaurant/model"
	restaurantstorage "Golang_Edu/modules/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateRestaurant(appCtx appctx.AppContext) func(*gin.Context) {
	return func(context *gin.Context) {
		id, convertError := strconv.Atoi(context.Param("id"))
		if convertError != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": convertError.Error()})
			return
		}
		var updateData restaurantmodel.RestaurantUpdate
		if err := context.ShouldBind(&updateData); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		updateRestaurantBiz := restaurantbusiness.NewUpdateRestaurantBiz(store)
		if err := updateRestaurantBiz.UpdateRestaurant(context.Request.Context(), &updateData, id); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
