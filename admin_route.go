package main

import (
	"Golang_Edu/common"
	"Golang_Edu/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func adminRoute(g *gin.RouterGroup, appCtx appctx.AppContext) {
	g.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, common.SimpleSuccessResponse("ok"))
	})
}
