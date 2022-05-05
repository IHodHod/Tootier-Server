package model

import "time"

type Mention struct {
	MentionID       uint64    `json:"mention_id" gorm:"primaryKey"`
	TootiID         uint64    `json:"tooti_id"`
	MentionTime     time.Time `json:"mention_time"`
	QouteID         uint64    `json:"qoute_id"`
	ReTootiID       uint64    `json:"re_tooti_id"`
	ProfileID       uint64    `json:"profile_id"`
	CommentID       uint64    `json:"comment_id"`
	MentionerUserID uint64    `json:"mentioner_user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MentionedUserID uint64    `json:"mentioned_user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Notifes         []Notify  `json:"notifes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
