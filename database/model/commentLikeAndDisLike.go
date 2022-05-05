package model

type CommentLike struct {
	CommentLikeID uint64   `json:"comment_like_id" gorm:"primaryKey" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	CommentID     uint64   `json:"comment_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	UserID        uint64   `json:"user_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	Notifes       []Notify `json:"notifes"`
}

type CommentDisLike struct {
	CommentDisLikeID uint64   `json:"comment_dis_like_id" gorm:"primaryKey" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	CommentID        uint64   `json:"comment_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	UserID           uint64   `json:"user_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	Notifes          []Notify `json:"notifes"`
}
