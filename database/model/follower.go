package model

import "time"

type Follower struct {
	FollowerID     uint64    `json:"follower_id" gorm:"primaryKey" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	UserID         uint64    `json:"user_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	FollowerUserID uint64    `json:"follower_user_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	CreatedAt      time.Time `json:"created_at" faker:"time"`
}
