package restaurantgin

import (
	"Golang_Edu/component/appctx"
	restaurantbusiness "Golang_Edu/modules/restaurant/business"
	restaurantstorage "Golang_Edu/modules/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ActivateRestaurant(appCtx appctx.AppContext) func(*gin.Context) {
	return func(context *gin.Context) {
		// Convert id
		id, convertError := strconv.Atoi(context.Param("id"))
		if convertError != nil {
			context.JSON(http.StatusBadGateway, gin.H{"error": convertError.Error()})
			return
		}
		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbusiness.NewActivateRestaurantBiz(store)
		err := biz.ActivateRestaurant(context, id)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
