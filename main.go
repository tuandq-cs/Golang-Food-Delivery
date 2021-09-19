package main

import (
	"Golang_Edu/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

var db *gorm.DB

type Restaurant struct {
	common.SQLModel
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:addr;"`
}

type RestaurantCreate struct {
	common.SQLModel
	Name    *string `binding:"required" json:"name" form:"name" gorm:"column:name;"`
	Address *string `binding:"required" json:"address" form:"address" gorm:"column:addr;"`
}

type RestaurantUpdate struct {
	common.SQLModel
	Name    *string `json:"name" form:"name" gorm:"column:name;"`
	Address *string `json:"address" form:"address" gorm:"column:addr;"`
}

func (Restaurant) TableName() string       { return "restaurants" }
func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }
func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

type Paging struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

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
	var err error
	db, err = gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db = db.Debug()
	router := gin.Default()
	// Version 1
	v1 := router.Group("/v1")
	{
		restaurants := v1.Group("restaurants")
		{
			restaurants.GET("/:id", getRestaurantById)
			restaurants.GET("", getRestaurants)
			restaurants.POST("", createRestaurant)
			restaurants.PATCH("/:id", updateRestaurant)
			restaurants.PATCH("/:id/block", blockRestaurant)
			restaurants.PATCH("/:id/activate", activateRestaurant)
			restaurants.DELETE("/:id", deleteRestaurant)
		}
	}
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func activateRestaurant(context *gin.Context) {
	id, convertError := strconv.Atoi(context.Param("id"))
	if convertError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": convertError.Error()})
		return
	}
	activateResult := db.Table(Restaurant{}.TableName()).Where("id = ?", id).Update("status", 1)
	if activateResult.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": activateResult.Error.Error()})
		return
	}
	if activateResult.RowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "This restaurant has been activated"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": true})
}

func blockRestaurant(context *gin.Context) {
	id, convertError := strconv.Atoi(context.Param("id"))
	if convertError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": convertError.Error()})
		return
	}
	blockResult := db.Table(Restaurant{}.TableName()).Where("id = ?", id).Update("status", 0)
	if blockResult.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": blockResult.Error.Error()})
		return
	}
	if blockResult.RowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "This restaurant has been blocked"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": true})
}

func deleteRestaurant(context *gin.Context) {
	id, convertError := strconv.Atoi(context.Param("id"))
	if convertError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": convertError.Error()})
		return
	}
	deleteResult := db.Table(Restaurant{}.TableName()).Where("id = ?", id).Delete(nil)
	if deleteResult.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": deleteResult.Error.Error()})
		return
	}
	if deleteResult.RowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "This restaurant has been removed"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": true})
}

func updateRestaurant(context *gin.Context) {
	id, convertError := strconv.Atoi(context.Param("id"))
	if convertError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": convertError.Error()})
		return
	}
	var updateData RestaurantUpdate
	if err := context.ShouldBind(&updateData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Update restaurant
	updateResult := db.Where("id = ? AND status = 1", id).Updates(updateData)
	if updateResult.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": updateResult.Error.Error()})
		return
	}
	if updateResult.RowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "This restaurant has been blocked or removed"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": true})
}

func createRestaurant(context *gin.Context) {
	var data RestaurantCreate
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := db.Create(&data)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": data.Id})
}

func getRestaurants(context *gin.Context) {
	var paging Paging
	if err := context.ShouldBind(&paging); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("Paging object", paging)
	if paging.Page <= 0 {
		paging.Page = 1
	}
	if paging.Limit <= 0 {
		paging.Limit = 10
	}
	var listData []Restaurant
	offset := (paging.Page - 1) * paging.Limit
	queryResult := db.Offset(offset).Limit(paging.Limit).Order("id desc").Find(&listData)
	if queryResult.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": queryResult.Error.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": listData})
}

func getRestaurantById(context *gin.Context) {
	id, convertError := strconv.Atoi(context.Param("id"))
	if convertError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": convertError.Error()})
		return
	}
	var data Restaurant
	result := db.First(&data, id)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": data})
}
