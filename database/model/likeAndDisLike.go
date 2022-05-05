package model

import "time"

type Like struct {
	LikeID    uint64    `json:"like_id" gorm:"primaryKey" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	UserID    uint64    `json:"user_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	LikeTime  time.Time `json:"like_time" faker:"time"`
	TootiID   uint64    `json:"tooti_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	QouteID   uint64    `json:"qoute_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	ReTootiID uint64    `json:"re_tooti_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	Notifes   []Notify  `json:"notifes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type DisLike struct {
	DisLikeID   uint64    `json:"dis_like_id" gorm:"primaryKey" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	UserID      uint64    `json:"user_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	DisLikeTime time.Time `json:"dis_like_time" faker:"time"`
	TootiID     uint64    `json:"tooti_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	QouteID     uint64    `json:"qoute_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	ReTootiID   uint64    `json:"re_tooti_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	Notifes     []Notify  `json:"notifes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
