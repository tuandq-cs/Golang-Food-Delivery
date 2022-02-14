package usergin

import (
	"Golang_Edu/common"
	"Golang_Edu/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Profile(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("user")
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(user))
	}
}
