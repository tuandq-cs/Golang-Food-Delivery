package restaurantlikegin

import (
	"Golang_Edu/common"
	"Golang_Edu/component/appctx"
	restaurantlikebusiness "Golang_Edu/modules/restaurantlike/business"
	restaurantlikemodel "Golang_Edu/modules/restaurantlike/model"
	restaurantlikestorage "Golang_Edu/modules/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsersLikeRestaurant(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var paging common.Paging
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := paging.Preprocess(); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantlikebusiness.NewListUsersLikeRestaurantBiz(store)
		filter := restaurantlikemodel.Filter{RestaurantId: int(uid.GetLocalID())}
		listUsers, err := biz.ListUsersLikeRestaurant(
			ctx.Request.Context(),
			&filter,
			&paging,
		)
		if err != nil {
			panic(err)
		}
		for i := range listUsers {
			listUsers[i].Mask(common.DbTypeUser)
		}
		ctx.JSON(http.StatusOK, common.NewSuccessResponse(listUsers, paging, filter))
	}
}
