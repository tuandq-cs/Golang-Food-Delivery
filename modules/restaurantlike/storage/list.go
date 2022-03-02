package restaurantlikestorage

import (
	"Golang_Edu/common"
	restaurantlikemodel "Golang_Edu/modules/restaurantlike/model"
	"context"
	"github.com/btcsuite/btcutil/base58"
	"time"
)

func (store *sqlStore) GetUserLikes(
	context context.Context,
	ids []int,
) (map[int]int, error) {
	db := store.db
	userLikes := make(map[int]int)
	type sqlData struct {
		RestaurantId int `gorm:"column:restaurant_id;"`
		LikeCount    int `gorm:"column:count;"`
	}
	var result []sqlData
	if err := db.Table(restaurantlikemodel.Like{}.TableName()).
		Select("restaurant_id, count(user_id) as count").
		Where("restaurant_id IN ?", ids).
		Group("restaurant_id").
		Find(&result).Error; err != nil {
		return userLikes, common.ErrDB(err)
	}
	for _, item := range result {
		userLikes[item.RestaurantId] = item.LikeCount
	}
	return userLikes, nil
}

func (store *sqlStore) CheckUserHasLiked(context context.Context, userId int, resIds []int) (map[int]bool, error) {
	db := store.db
	userHasLiked := make(map[int]bool)
	type sqlData struct {
		RestaurantId int `gorm:"column:restaurant_id;"`
	}
	var result []sqlData
	if err := db.Table(restaurantlikemodel.Like{}.TableName()).
		Where("user_id = ? AND restaurant_id IN ?", userId, resIds).
		Find(&result).Error; err != nil {
		return userHasLiked, err
	}
	for _, item := range result {
		userHasLiked[item.RestaurantId] = true
	}
	return userHasLiked, nil
}

func (store *sqlStore) GetUsersLikeRestaurant(
	context context.Context,
	conditions map[string]interface{},
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
) ([]common.SimpleUser, error) {
	db := store.db
	var result []restaurantlikemodel.Like
	db = db.Table(restaurantlikemodel.Like{}.TableName()).Where(conditions)
	if v := filter; filter != nil {
		if v.RestaurantId > 0 {
			db = db.Where("restaurant_id = ?", v.RestaurantId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	if v := paging.FakeCursor; v != "" {
		createdAt, err := time.Parse(time.RFC3339Nano, string(base58.Decode(v)))
		if err != nil {
			return nil, common.ErrDB(err)
		}
		db = db.Where("created_at < ?", createdAt.Format("2006-01-02 15:04:05"))
	} else {
		db = db.Offset(paging.Offset())
	}

	db = db.Preload("User")
	if err := db.Order("created_at desc").Limit(paging.Limit).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	listUsers := make([]common.SimpleUser, len(result))
	for i, _ := range result {
		listUsers[i] = *result[i].User
		if i == len(result)-1 {
			paging.NextCursor = base58.Encode([]byte(result[i].CreatedAt.Format(time.RFC3339Nano)))
		}
	}
	return listUsers, nil
}
