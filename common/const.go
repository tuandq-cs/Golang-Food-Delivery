package common

import "errors"

var (
	ErrDataNotFound = errors.New("data not found")
)

const (
	DbTypeRestaurant = 1
)
