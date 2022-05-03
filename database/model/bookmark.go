package model

type Bookmark struct {
	BookmarkID uint64 `json:"bookmark_id" gorm:"primaryKey"`
	TootiID uint64 `json:"tooti_id" gorm:"default:0"`
	ReTootiID uint64 `json:"re_tooti_id" gorm:"default:0"`
	QouteID uint64 `json:"qoute_id" gorm:"default:0"`
	UserID uint64 `json:"user_id"`
}
