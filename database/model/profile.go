package model

import "time"

type Profile struct {
	ProfileID uint64 `json:"profile_id" gorm:"primaryKey"`
	UserID uint64 `josn:"user_id"` // TODO relation
	ImageProfileURL string `json:"image_profile_url"`
	ImageCoverURL string `json:"image_cover_url"`
	Bioghraphy string `json:"bioghraphy"`
	Birthday time.Time `json:"birthday"`
	Gender uint8 `json:"gender"`
	EnableNotify bool `json:"enable_notify" default:"true"` // TODO
}

