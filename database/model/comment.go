package model

type Comment struct {
	CommentID       uint64           `json:"comment_id" gorm:"primaryKey"`
	Title           string           `json:"title" gorm:"default:null"`
	MediaID         uint64           `json:"media_id"`
	UserID          uint64           `json:"user_id"`
	TootiID         uint64           `json:"tootie_id"`
	QouteID         uint64           `json:"qoute_id"`
	ReTootiID       uint64           `json:"re_tooti_id"`
	Hashtags        []*Hashtag       `json:"hashtags"  gorm:"many2many:comment_hashtags"`
	CommentLikes    []CommentLike    `json:"comment_likes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CommentDisLikes []CommentDisLike `json:"comment_dis_likes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Mentions        []Mention        `json:"mentions" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Notifes         []Notify         `json:"notifes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
