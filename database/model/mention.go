package model

import "time"

type Mention struct {
	MentionID       uint64    `json:"mention_id" gorm:"primaryKey" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	TootiID         uint64    `json:"tooti_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	MentionTime     time.Time `json:"mention_time" faker:"time"`
	QouteID         uint64    `json:"qoute_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	ReTootiID       uint64    `json:"re_tooti_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	ProfileID       uint64    `json:"profile_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	CommentID       uint64    `json:"comment_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	MentionerUserID uint64    `json:"mentioner_user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	MentionedUserID uint64    `json:"mentioned_user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	Notifes         []Notify  `json:"notifes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
