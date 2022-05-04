package model

type CommentLike struct {
	CommentLikeID uint64   `json:"comment_like_id" gorm:"primaryKey"`
	CommentID     uint64   `json:"comment_id"`
	UserID        uint64   `json:"user_id"`
	Notifes       []Notify `json:"notifes"`
}

type CommentDisLike struct {
	CommentDisLikeID uint64   `json:"comment_dis_like_id" gorm:"primaryKey"`
	CommentID        uint64   `json:"comment_id"`
	UserID           uint64   `json:"user_id"`
	Notifes          []Notify `json:"notifes"`
}
