package restaurantgin

import (
	"Golang_Edu/component/appctx"
	restaurantbusiness "Golang_Edu/modules/restaurant/business"
	restaurantstorage "Golang_Edu/modules/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteRestaurant(appCtx appctx.AppContext) func(*gin.Context) {
	return func(context *gin.Context) {
		id, convertError := strconv.Atoi(context.Param("id"))
		if convertError != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": convertError.Error()})
			return
		}
		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbusiness.NewDeleteRestaurantBiz(store)
		if err := biz.DeleteRestaurant(context.Request.Context(), id); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"data": true})
	}
}
