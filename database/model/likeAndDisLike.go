package model

import "time"

type Like struct {
	LikeID        uint64    `json:"like_id" gorm:"primaryKey"`
	UserID        uint64    `json:"user_id"`
	LikeTime time.Time 	    `json:"like_time"`
	TootiID       uint64    `json:"tooti_id"`
	QouteID       uint64    `json:"qoute_id"`
	ReTootiID     uint64    `json:"re_tooti_id"`
	Notifes       []Notify  `json:"notifes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type DisLike struct {
	DisLikeID     uint64    `json:"dis_like_id" gorm:"primaryKey"`
	UserID        uint64    `json:"user_id"`
	DisLikeTime   time.Time `json:"dis_like_time"`
	TootiID       uint64    `json:"tooti_id"`
	QouteID       uint64    `json:"qoute_id"`
	ReTootiID     uint64    `json:"re_tooti_id"`
	Notifes       []Notify  `json:"notifes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
