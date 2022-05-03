package model

import "time"

type Follower struct {
	FollowerID uint64 `json:"follower_id" gorm:"primaryKey"`
	UserID     uint64 `json:"user_id"`// TODO relatoin
	FollowerUserID uint64 `json:"follower_user_id"`// TODO relatoin
	CreatedAt time.Time `json:"created_at"`
}
