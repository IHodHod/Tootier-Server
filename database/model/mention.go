package model

import "time"

type Mention struct {
	MentionID uint64 `json:"mention_id" gorm:"primaryKey"`

	TootiID uint64 `json:"tooti_id"` // TODO relation
	QouteID uint64 `json:"qoute_id"` 	  // TODO relation
	ReTootiID uint64 `json:"re_tooti_id"` // TODO relation
	ProfileID uint64 `json:"profile_id"` // TODO relation
	CommentID uint64 `json:"comment_id"` // TODO relation

	MentionedUserID uint64 `json:"mentioned_user_id"` // TODO relation
	MentionerUserID uint64 `json:"mentioner_user_id"` // TODO relation
	MentionTime time.Time `json:"mention_time"` // TODO relation
}
