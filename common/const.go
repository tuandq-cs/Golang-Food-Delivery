package common

import "fmt"

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
)

const CurrentUser = "user"

func Recovery() {
	if r := recover(); r != nil {
		fmt.Println("Recovered:", r)
	}
}

type Requester interface {
	GetUserId() int
	GetRole() string
	GetEmail() string
}
