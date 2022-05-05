package model

import "time"

type Notify struct {
	NotifyID         uint64    `json:"notify_id" gorm:"primaryKey"`
	IsSend           bool      `json:"is_send"`
	Description      string    `json:"description"`
	CreatedAt        time.Time `json:"created_at"`
	SendAt           time.Time `json:"send_at"`
	ToUserID         uint64    `json:"to_user_id"`
	MentionID        uint64    `json:"mention_id"`
	FollowingID      uint64    `json:"following_id"`
	CommentID        uint64    `json:"comment_id"`
	ReTootiID        uint64    `json:"re_tooti_id"`
	TootiID          uint64    `json:"tooti_id"`
	QouteID          uint64    `json:"qoute_id"`
	LikeID           uint64    `json:"like_id"`
	DisLikeID        uint64    `json:"dis_like_id"`
	CommentLikeID    uint64    `json:"comment_like_id"`
	CommentDisLikeID uint64    `json:"comment_dis_like_id"`
}
