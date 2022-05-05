package model

import "time"

type Tooti struct {
	TootiID   uint64     `json:"tooti_id" gorm:"primaryKey" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	Title     string     `json:"title" gorm:"default:null"`
	CreatedAt time.Time  `json:"created_at"`
	MediaID   uint64     `json:"media_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	Hashtags  []*Hashtag `json:"hashtags" gorm:"many2many:tooti_hashtags" faker:"slice_len=10"`
	UserID    uint64     `json:"user_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	Bookmark  []Bookmark `json:"bookmarks" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" faker:"slice_len=10"`
	Qoute     []Qoute    `json:"qoute" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" faker:"slice_len=10"`
	ReTooti   []ReTooti  `json:"re_tooti" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" faker:"slice_len=10"`
	Comments  []Comment  `json:"comments" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" faker:"slice_len=10"`
	Likes     []Like     `json:"likes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" faker:"slice_len=10"`
	DisLike   []DisLike  `json:"dis_like" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" faker:"slice_len=10"`
	Mentions  []Mention  `json:"mentions" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" faker:"slice_len=10"`
	Notifes   []Notify   `json:"notifes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" faker:"slice_len=10"`
	Views     []View     `json:"views" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" faker:"slice_len=10"`
}

type Qoute struct {
	QouteID     uint64     `json:"qoute_id" gorm:"primaryKey" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	MediaID     uint64     `json:"media_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	Title       string     `json:"title" gorm:"default:null"`
	CreatedAt   time.Time  `json:"created_at"`
	Hashtags    []*Hashtag `json:"hashtags" gorm:"many2many:qoute_hashtags" faker:"slice_len=10"`
	UserID      uint64     `json:"user_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	TootiID     uint64     `json:"tooti_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`            // one to has many
	FromQouteID *uint64    `json:"from_qoute_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`       // self refrencial has many
	Qoutes      []Qoute    `json:"qoutes" gorm:"foreignkey:FromQouteID" faker:"slice_len=10"` // self refrencial has many
	Comments    []Comment  `json:"comments" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" faker:"slice_len=10"`
	ReTootiID   uint64     `json:"re_tooti_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	Bookmark    []Bookmark `json:"bookmarks" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" faker:"slice_len=10"`
	Likes       []Like     `json:"likes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" faker:"slice_len=10"`
	DisLike     []DisLike  `json:"dis_like" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" faker:"slice_len=10"`
	Mentions    []Mention  `json:"mentions" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" faker:"slice_len=10"`
	Notifes     []Notify   `json:"notifes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" faker:"slice_len=10"`
	Views       []View     `json:"views" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" faker:"slice_len=10"`
	ReTooti     []ReTooti  `json:"re_tooti" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" faker:"slice_len=10"`
}

type ReTooti struct {
	ReTootiID   uint64     `json:"re_tooti_id" gorm:"primaryKey" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	UserID      uint64     `json:"user_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	Hashtags    []*Hashtag `json:"hashtags" gorm:"many2many:retooti_hashtags"`
	TootiID     uint64     `json:"tooti_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	QouteID     uint64     `json:"qoute_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	FromReTooti uint64     `json:"from_re_tooti" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"` //  self refrencial has many
	ReTooties   []ReTooti  `json:"re_tooties" gorm:"foreignkey:FromReTooti"`            // self refrencial has many
	Bookmark    []Bookmark `json:"bookmarks" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" `
	Qoutes      []Qoute    `json:"qoutes"`
	Comments    []Comment  `json:"comments" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" `
	Likes       []Like     `json:"likes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" `
	DisLike     []DisLike  `json:"dis_like" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" `
	Mentions    []Mention  `json:"mentions" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Notifes     []Notify   `json:"notifes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Views       []View     `json:"views" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
