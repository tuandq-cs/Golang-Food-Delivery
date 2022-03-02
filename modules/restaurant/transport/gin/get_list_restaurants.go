package restaurantgin

import (
	"Golang_Edu/common"
	"Golang_Edu/component/appctx"
	restaurantbusiness "Golang_Edu/modules/restaurant/business"
	restaurantmodel "Golang_Edu/modules/restaurant/model"
	restaurantrepository "Golang_Edu/modules/restaurant/repository"
	restaurantstorage "Golang_Edu/modules/restaurant/storage"
	restaurantlikestorage "Golang_Edu/modules/restaurantlike/storage"
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
		requester := context.MustGet(common.CurrentUser).(common.Requester)
		listRestaurantsStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		likedUsersStore := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		repo := restaurantrepository.NewGetListRestaurantsRepository(listRestaurantsStore, likedUsersStore, requester)
		biz := restaurantbusiness.NewGetListRestaurantsBiz(repo)
		listData, err := biz.GetListRestaurants(context.Request.Context(), &paging, &filter)
		if err != nil {
			panic(err)
		}
		for i := range listData {
			listData[i].Mask(common.DbTypeRestaurant)
			if i == len(listData)-1 {
				paging.NextCursor = listData[i].FakeId.String()
			}
		}
		context.JSON(http.StatusOK, common.NewSuccessResponse(listData, paging, filter))
	}
}
