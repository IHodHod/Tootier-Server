package model

import "time"

type Hashtag struct {
	HashtagID  uint64 `json:"hashtag_id" gorm:"primaryKey" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	HashMainID uint64 `json:"hash_main_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`

	Profiles  []*Profile `json:"profiles" gorm:"many2many:profile_hashtags" `
	Tooties   []*Tooti   `json:"tooties" gorm:"many2many:tooti_hashtags" `
	Qoutes    []*Qoute   `json:"qoutes" gorm:"many2many:qoute_hashtags" `
	ReTooties []*ReTooti `json:"re_tooties" gorm:"many2many:retooti_hashtags" `
	Comments  []*Comment `json:"comments"  gorm:"many2many:comment_hashtags" `
}

type HashMain struct {
	HashMainID uint64    `json:"hash_main_id" gorm:"primaryKey" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	Text       string    `json:"text" faker:"sentence"`
	CreatedAt  time.Time `json:"created_at" faker:"time"`
	Hashtags   []Hashtag `json:"hashtag" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
