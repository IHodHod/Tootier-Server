package model


type Comment struct {
	CommentID uint64 `json:"comment_id" gorm:"primaryKey"`
	Title string `json:"title"` 	      // TODO  nullable
	MediaID uint64 `json:"media_id"`      // TODO relation

	TootiID uint64 `json:"tooti_id"` 	  // TODO relation
	QouteID uint64 `json:"qoute_id"` 	  // TODO relation
	ReTootiID uint64 `json:"re_tooti_id"` // TODO relation
}
