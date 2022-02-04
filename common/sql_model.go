package common

import "time"

type SQLModel struct {
	Id        int       `json:"id" gorm:"column:id;<-:false;"`
	Status    int       `json:"status" gorm:"column:status;default:1;"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;"`
}
