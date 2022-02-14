package main

import (
	"Golang_Edu/component/appctx"
	"Golang_Edu/middleware"
	restaurantgin "Golang_Edu/modules/restaurant/transport/gin"
	uploadgin "Golang_Edu/modules/upload/transport/gin"
	usergin "Golang_Edu/modules/user/transport/gin"
	"github.com/gin-gonic/gin"
)

func mainRoute(g *gin.RouterGroup, appCtx appctx.AppContext, authStore middleware.AuthStore) {
	g.POST("/upload", uploadgin.UploadImage(appCtx))
	g.POST("/register", usergin.Register(appCtx))
	g.POST("/authenticate", usergin.Login(appCtx))
	g.GET("/profile", middleware.RequiredAuth(appCtx, authStore), usergin.Profile(appCtx))

	restaurants := g.Group("restaurants", middleware.RequiredAuth(appCtx, authStore))
	{
		restaurants.GET("/:id", restaurantgin.GetRestaurant(appCtx))
		restaurants.GET("", restaurantgin.GetListRestaurants(appCtx))
		restaurants.POST("", restaurantgin.CreateRestaurant(appCtx))
		restaurants.PATCH("/:id", restaurantgin.UpdateRestaurant(appCtx))
		restaurants.PATCH("/:id/inactivate", restaurantgin.InactivateRestaurant(appCtx))
		restaurants.PATCH("/:id/activate", restaurantgin.ActivateRestaurant(appCtx))
		restaurants.DELETE("/:id", restaurantgin.DeleteRestaurant(appCtx))
	}
}
