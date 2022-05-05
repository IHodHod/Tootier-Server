package model

import "time"

type Notify struct {
	NotifyID         uint64    `json:"notify_id" gorm:"primaryKey" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	IsSend           bool      `json:"is_send" faker:"oneof:true,false"`
	Description      string    `json:"description" faker:"sentence"`
	CreatedAt        time.Time `json:"created_at" faker:"time"`
	SendAt           time.Time `json:"send_at" faker:"time"`
	ToUserID         uint64    `json:"to_user_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	MentionID        uint64    `json:"mention_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	FollowingID      uint64    `json:"following_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	CommentID        uint64    `json:"comment_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	ReTootiID        uint64    `json:"re_tooti_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	TootiID          uint64    `json:"tooti_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	QouteID          uint64    `json:"qoute_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	LikeID           uint64    `json:"like_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	DisLikeID        uint64    `json:"dis_like_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	CommentLikeID    uint64    `json:"comment_like_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	CommentDisLikeID uint64    `json:"comment_dis_like_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
}
