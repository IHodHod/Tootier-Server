package model


type Comment struct {
	CommentID uint64 `json:"comment_id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"default:null"`
	MediaID   	uint64 `json:"media_id"`
	Tooties   []Tooti `json:"tooties" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Qoutes 	  []Qoute `json:"qoutes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ReTooties []ReTooti `json:"re_tooties" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Hashtags  []*Hashtag `json:"hashtags"  gorm:"many2many:comments_hashtags"`
}
