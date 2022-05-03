package model

type Bookmark struct {
	BookmarkID uint64 `json:"bookmark_id" gorm:"primaryKey"`
	TootiID uint64 `json:"tooti_id" ` // TODO default null - relation
	UserID uint64 `json:"user_id"` // TODO relation
}
