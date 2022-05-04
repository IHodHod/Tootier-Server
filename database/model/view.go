package model

type View struct {
	ViewID uint64 `json:"view_id" gorm:"primaryKey"`
	UserID uint64 `json:"user_id"`
	TootiID uint64 `json:"tooti_id"`
	QouteID uint64 `json:"qoute_id"`
	ReTootiID uint64 `json:"re_tooti_id"`
}
