package middleware

import (
	"Golang_Edu/common"
	"Golang_Edu/component/appctx"
	"github.com/gin-gonic/gin"
)

func RequiredRoles(appCtx appctx.AppContext, roles ...string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user := ctx.MustGet(common.CurrentUser).(common.Requester)
		for i := range roles {
			if user.GetRole() == roles[i] {
				ctx.Next()
				return
			}
		}
		panic(common.ErrNoPermission(nil))
	}
}
