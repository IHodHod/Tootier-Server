package model

type View struct {
	ViewID uint64 `json:"view_id" gorm:"primaryKey"`
	TootiID uint64 `json:"tooti_id"` // TODO relation
	UserID uint64 `json:"user_id"` // TODO relation
}
