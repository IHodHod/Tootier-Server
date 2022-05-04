package model

import "time"

type Follower struct {
	FollowerID     uint64    `json:"follower_id" gorm:"primaryKey"`
	UserID         uint64    `json:"user_id"`
	FollowerUserID uint64    `json:"follower_user_id"`
	CreatedAt      time.Time `json:"created_at"`
}
