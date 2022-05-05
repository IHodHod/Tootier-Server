package model

import "time"

type Following struct {
	FollowingID     uint64    `json:"following_id" gorm:"primaryKey" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	UserID          uint64    `json:"user_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	FollowingUserID uint64    `json:"following_user_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	CreatedAt       time.Time `json:"created_at" faker:"time"`
	Notifes         []Notify  `json:"notifes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
