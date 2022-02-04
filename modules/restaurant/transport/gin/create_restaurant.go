package restaurantgin

import (
	"Golang_Edu/component/appctx"
	restaurantbusiness "Golang_Edu/modules/restaurant/business"
	restaurantmodel "Golang_Edu/modules/restaurant/model"
	restaurantstorage "Golang_Edu/modules/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(appCtx appctx.AppContext) func(*gin.Context) {
	return func(context *gin.Context) {
		var data restaurantmodel.RestaurantCreate
		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		createRestaurantBiz := restaurantbusiness.NewCreateRestaurantBiz(store)
		if err := createRestaurantBiz.CreateNewRestaurant(context.Request.Context(), &data); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": data.Id})
	}
}
