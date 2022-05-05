package model

type Bookmark struct {
	BookmarkID uint64 `json:"bookmark_id" gorm:"primaryKey" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	TootiID    uint64 `json:"tooti_id" gorm:"default:0" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	ReTootiID  uint64 `json:"re_tooti_id" gorm:"default:0" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	QouteID    uint64 `json:"qoute_id" gorm:"default:0" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	UserID     uint64 `json:"user_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
}
