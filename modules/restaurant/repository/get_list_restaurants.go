package restaurantrepository

import (
	"Golang_Edu/common"
	"Golang_Edu/modules/restaurant/model"
	"context"
	"log"
)

type GetListRestaurantsStore interface {
	ListDataWithConditions(
		context context.Context,
		paging *common.Paging,
		filter *restaurantmodel.Filter,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type LikedUsersStore interface {
	GetUserLikes(context context.Context, ids []int) (map[int]int, error)
	CheckUserHasLiked(context context.Context, userId int, resIds []int) (map[int]bool, error)
}

type getListRestaurantsRepository struct {
	listRestaurantStore GetListRestaurantsStore
	likedUsersStore     LikedUsersStore
	requester           common.Requester
}

func NewGetListRestaurantsRepository(
	listRestaurantStore GetListRestaurantsStore,
	likedUsersStore LikedUsersStore,
	requester common.Requester,
) *getListRestaurantsRepository {
	return &getListRestaurantsRepository{
		listRestaurantStore: listRestaurantStore,
		likedUsersStore:     likedUsersStore,
		requester:           requester,
	}
}

func (repo *getListRestaurantsRepository) GetListRestaurants(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	restaurants, err := repo.listRestaurantStore.ListDataWithConditions(context, paging, filter, moreKeys...)
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}
	// Get list restaurant id to query restaurant_likes table
	resIds := make([]int, len(restaurants))
	for i := range restaurants {
		resIds[i] = restaurants[i].Id
	}

	// Mapping map userLikes, userHasLiked to list restaurants
	//userLikes, err := repo.likedUsersStore.GetUserLikes(context, resIds)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	userHasLiked, err := repo.likedUsersStore.CheckUserHasLiked(context, repo.requester.GetUserId(), resIds)
	if err != nil {
		log.Fatalln(err)
	}
	for i, item := range restaurants {
		//restaurants[i].LikedCount = userLikes[item.Id]
		restaurants[i].HasLiked = userHasLiked[item.Id]
	}

	return restaurants, err

}
