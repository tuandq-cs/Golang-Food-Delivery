package middleware

import (
	"Golang_Edu/common"
	"Golang_Edu/component/appctx"
	"github.com/gin-gonic/gin"
)

func Recover(appCxt appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				context.Header("Content-Type", "application/json")
				if appErr, ok := err.(*common.AppError); ok {
					context.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err)
				}
				appErr := common.ErrInternal(err.(error))
				context.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
			}
		}()

		context.Next()
	}
}
