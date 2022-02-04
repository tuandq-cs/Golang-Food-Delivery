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

func GetListRestaurants(appCtx appctx.AppContext) func(*gin.Context) {
	return func(context *gin.Context) {
		// Parse Paging
		var paging common.Paging
		if err := context.ShouldBind(&paging); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// Preprocess Paging
		if err := paging.Preprocess(); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// Parse Filter
		var filter restaurantmodel.Filter
		if err := context.ShouldBind(&filter); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// Main
		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbusiness.NewGetListRestaurantsBiz(store)
		listData, err := biz.GetListRestaurants(context.Request.Context(), &paging, &filter)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": listData, "paging": paging})
	}
}
