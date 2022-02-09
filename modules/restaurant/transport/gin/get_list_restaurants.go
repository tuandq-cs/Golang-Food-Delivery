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
			panic(err)
		}
		// Preprocess Paging
		if err := paging.Preprocess(); err != nil {
			panic(err)
		}
		// Parse Filter
		var filter restaurantmodel.Filter
		if err := context.ShouldBind(&filter); err != nil {
			panic(err)
		}
		// Main
		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbusiness.NewGetListRestaurantsBiz(store)
		listData, err := biz.GetListRestaurants(context.Request.Context(), &paging, &filter)
		if err != nil {
			panic(err)
		}
		for i := range listData {
			listData[i].Mask(common.DbTypeRestaurant)
		}
		context.JSON(http.StatusOK, common.NewSuccessResponse(listData, paging, filter))
	}
}
