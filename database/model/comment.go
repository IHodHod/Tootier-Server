package model

type Comment struct {
	CommentID       uint64           `json:"comment_id" gorm:"primaryKey" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	Title           string           `json:"title" gorm:"default:null" faker:"sentence"`
	MediaID         uint64           `json:"media_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	UserID          uint64           `json:"user_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	TootiID         uint64           `json:"tootie_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	QouteID         uint64           `json:"qoute_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	ReTootiID       uint64           `json:"re_tooti_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	Hashtags        []*Hashtag       `json:"hashtags"  gorm:"many2many:comment_hashtags"`
	CommentLikes    []CommentLike    `json:"comment_likes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" `
	CommentDisLikes []CommentDisLike `json:"comment_dis_likes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" `
	Mentions        []Mention        `json:"mentions" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" `
	Notifes         []Notify         `json:"notifes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
