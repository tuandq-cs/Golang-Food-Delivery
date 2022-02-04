package main

import (
	"Golang_Edu/component/appctx"
	restaurantgin "Golang_Edu/modules/restaurant/transport/gin"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//dbDriver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	dbOptions := os.Getenv("DB_OPTIONS")
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, password, host, port, dbName, dbOptions)
	//dsn := "food_delivery:12345@tcp(127.0.0.1:3307)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db = db.Debug()

	appCtx := appctx.NewAppContext(db)

	router := gin.Default()
	// Version 1
	v1 := router.Group("/v1")
	{
		restaurants := v1.Group("restaurants")
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
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
