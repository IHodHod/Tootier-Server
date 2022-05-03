package model

import "time"

type Tooti struct {
	TootiID   uint64 `json:"tooti_id" gorm:"primaryKey"`
	MediaID   uint64 `json:"media_id"` // TODO relation
	HashtagID uint64 `json:"hashtags"` // TODO relation slice
	Title 	  string `json:"title"` // TODO title nullable
	DeviceID  uint64 `json:"device_id"` // TODO relation
	CreatedAt time.Time `json:"created_at"`
}

type Qoute struct {
	QouteID   	uint64 `json:"qoute_id" gorm:"primaryKey"`
	MediaID   	uint64 `json:"media_id"` // TODO relation
	HashtagID 	uint64 `json:"hashtags"` // TODO relation slice
	Title 	  	string `json:"title"` // TODO title nullable
	DeviceID    uint64 `json:"device_id"` // TODO relation
	CreatedAt   time.Time `json:"created_at"`

	TootiID     uint64 `json:"tooti_id"` // TODO relation
	FromQouteID uint64 `json:"from_qoute_id"` // TODO relation
	ReTootiID   uint64 `json:"re_tooti_id"` // TODO relation
}

type ReTooti struct {
	ReTootiID  uint64 `json:"re_tooti_id" gorm:"primaryKey"`
	DeviceID uint64 `json:"device_id"`// TODO relation

	QouteID uint64 `json:"from_qoute_id"` // TODO relation
	TootiID     uint64 `json:"tooti_id"` // TODO relation
	FromReTooti uint64 `json:"from_re_tooti"` // TODO relation
}


