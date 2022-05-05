package model

import "time"

type Profile struct {
	ProfileID       uint64     `json:"profile_id" gorm:"primaryKey" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	UserID          uint64     `josn:"user_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	ImageProfileURL string     `json:"image_profile_url" faker:"url"`
	ImageCoverURL   string     `json:"image_cover_url" faker:"url"`
	Bioghraphy      string     `json:"bioghraphy" faker:"sentence"`
	Birthday        time.Time  `json:"birthday" faker:"time"`
	Gender          uint8      `json:"gender" faker:"oneof:0,1"`
	EnableNotify    bool       `json:"enable_notify" gorm:"default:true" faker:"oneof:true,false"`
	Hashtags        []*Hashtag `json:"hashtags" gorm:"many2many:profile_hashtags"`

	Mentions []Mention `json:"mentions" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
