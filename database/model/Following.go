package model

import "time"

type Following struct {
	FollowingID uint64 `json:"following_id" gorm:"primaryKey"`
	UserID     uint64 `json:"user_id"`
	FollowingUserID uint64 `json:"following_user_id"`

	CreatedAt time.Time `json:"created_at"`
}
