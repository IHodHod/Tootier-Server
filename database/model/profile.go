package model

import "time"

type Profile struct {
	ProfileID       uint64     `json:"profile_id" gorm:"primaryKey"`
	UserID          uint64     `josn:"user_id"`
	ImageProfileURL string     `json:"image_profile_url"`
	ImageCoverURL   string     `json:"image_cover_url"`
	Bioghraphy      string     `json:"bioghraphy"`
	Birthday        time.Time  `json:"birthday"`
	Gender          uint8      `json:"gender"`
	EnableNotify    bool       `json:"enable_notify" gorm:"default:true"`
	Hashtags        []*Hashtag `json:"hashtags" gorm:"many2many:profile_hashtags"`

	Mentions []Mention `json:"mentions" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
