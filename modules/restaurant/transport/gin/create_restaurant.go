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

func CreateRestaurant(appCtx appctx.AppContext) func(*gin.Context) {
	return func(context *gin.Context) {
		var data restaurantmodel.RestaurantCreate
		if err := context.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		user := context.MustGet(common.CurrentUser).(common.Requester)
		data.OwnerId = user.GetUserId()
		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		createRestaurantBiz := restaurantbusiness.NewCreateRestaurantBiz(store)
		if err := createRestaurantBiz.CreateNewRestaurant(context.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(common.DbTypeRestaurant)
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId))
	}
}
