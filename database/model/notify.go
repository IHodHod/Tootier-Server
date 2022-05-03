package model

import "time"


type Notify struct {
	NotifyID uint64 `json:"notify_id" gorm:"primaryKey"`
	IsSend bool `json:"is_send"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	SendAt time.Time `json:"send_at"`

	ToUserID uint64 `json:"to_user_id"`// TODO relatoin

	MentionID uint64 `json:"mention_id"`// TODO relatoin
	FollowingID uint64 `json:"following_id"` // TODO relatoin
	CommentID uint64 `json:"comment_id"`// TODO relatoin
	ReTootiID uint64 `json:"re_tooti_id"`// TODO relatoin
	TootiID uint64 `json:"tooti_id"`// TODO relatoin
	QouteID uint64 `json:"qoute_id"`// TODO relatoin
	LikeID uint64 `json:"like_id"`// TODO relatoin
	DisLikeID uint64 `json:"dis_like_id"`// TODO relatoin
	CommentLikeID uint64 `json:"comment_like_id"`// TODO relatoin
	CommentDisLikeID uint64 `json:"comment_dis_like_id"`// TODO relatoin
}
