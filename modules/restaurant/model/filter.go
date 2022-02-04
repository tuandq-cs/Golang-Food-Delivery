package restaurantmodel

type Filter struct {
	UserId int `json:"-" form:"user_id"`
}
