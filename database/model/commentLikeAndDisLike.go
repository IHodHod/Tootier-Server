package model



type CommentLike struct {
	CommentLikeID uint64 `json:"comment_like_id" gorm:"pimaryKey"`
	CommentID uint64 `json:"comment_id"` // TODO relation
	UserID uint64 `json:"user_id"` // TODO relation
}

type CommentDisLike struct {
	CommentDisLikeID uint64 `json:"comment_dis_like_id" gorm:"pimaryKey"`
	CommentID uint64 `json:"comment_id"` // TODO relation
	UserID uint64 `json:"user_id"` // TODO relation
}

