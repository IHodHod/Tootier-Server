package model

import "time"

type Hashtag struct {
	HashtagID uint64 `json:"hashtag_id" gorm:"primaryKey"`
	HashMainID uint64 `json:"hash_main_id"` // TODO relation
	ProfileID uint64 `json:"profile_id"`// TODO relation
	TootiID uint64 `json:"tooti_id"`// TODO relation
	QouteID uint64 `json:"qoute_id"`// TODO relation
	ReTootiID uint64 `json:"re_tooti_id"`// TODO relation
	CommentID uint64 `json:"comment_id"`// TODO relation
}


type HashMain struct {
	HashMainID uint64 `json:"hash_main_id" gorm:"primaryKey"`
	Text string `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
