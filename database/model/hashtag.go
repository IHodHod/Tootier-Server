package model

import "time"

type Hashtag struct {
	HashtagID  uint64 `json:"hashtag_id" gorm:"primaryKey"`
	HashMainID   uint64 `json:"hash_main_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Profiles   []*Profile `json:"profiles" gorm:"many2many:profile_hashtags"`
	Tooties    []*Tooti   `json:"tooties" gorm:"many2many:tooti_hashtags"`
	Qoutes     []*Qoute   `json:"qoutes" gorm:"many2many:qoute_hashtags"`
	ReTooties  []*ReTooti `json:"re_tooties" gorm:"many2many:retooti_hashtags"`
	Comments   []*Comment `json:"comments"  gorm:"many2many:comment_hashtags"`
}

type HashMain struct {
	HashMainID uint64 `json:"hash_main_id" gorm:"primaryKey"`
	Text string `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	Hashtags []Hashtag `json:"hashtag" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
