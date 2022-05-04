package model

import "time"

type Tooti struct {
	TootiID   uint64     `json:"tooti_id" gorm:"primaryKey"`
	Title     string     `json:"title" gorm:"default:null"`
	CreatedAt time.Time  `json:"created_at"`
	MediaID   uint64     `json:"media_id"`
	Hashtags  []*Hashtag `json:"hashtags" gorm:"many2many:tooti_hashtags"`
	UserID    uint64     `json:"user_id"`
	Bookmark  []Bookmark `json:"bookmarks" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Qoute     []Qoute    `json:"qoute" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ReTooti   []ReTooti  `json:"re_tooti" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comments  []Comment  `json:"comments" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Likes     []Like     `json:"likes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DisLike   []DisLike  `json:"dis_like" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Mentions  []Mention  `json:"mentions" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Notifes   []Notify   `json:"notifes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Views     []View     `json:"views" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Qoute struct {
	QouteID     uint64     `json:"qoute_id" gorm:"primaryKey"`
	MediaID     uint64     `json:"media_id"`
	Title       string     `json:"title" gorm:"default:null"`
	CreatedAt   time.Time  `json:"created_at"`
	Hashtags    []*Hashtag `json:"hashtags" gorm:"many2many:qoute_hashtags"`
	UserID      uint64     `json:"user_id"`
	TootiID     uint64     `json:"tooti_id"`                             // one to has many
	FromQouteID *uint64    `json:"from_qoute_id"`                        // self refrencial has many
	Qoutes      []Qoute    `json:"qoutes" gorm:"foreignkey:FromQouteID"` // self refrencial has many
	Comments    []Comment  `json:"comments" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ReTootiID   uint64     `json:"re_tooti_id"`
	Bookmark    []Bookmark `json:"bookmarks" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Likes       []Like     `json:"likes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DisLike     []DisLike  `json:"dis_like" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Mentions    []Mention  `json:"mentions" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Notifes     []Notify   `json:"notifes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Views       []View     `json:"views" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ReTooti     []ReTooti  `json:"re_tooti" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ReTooti struct {
	ReTootiID   uint64     `json:"re_tooti_id" gorm:"primaryKey"`
	UserID      uint64     `json:"user_id"`
	Hashtags    []*Hashtag `json:"hashtags" gorm:"many2many:retooti_hashtags"`
	TootiID     uint64     `json:"tooti_id"`
	QouteID     uint64     `json:"qoute_id"`
	FromReTooti uint64     `json:"from_re_tooti"`                            //  self refrencial has many
	ReTooties   []ReTooti  `json:"re_tooties" gorm:"foreignkey:FromReTooti"` // self refrencial has many
	Bookmark    []Bookmark `json:"bookmarks" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Qoutes      []Qoute    `json:"qoutes"`
	Comments    []Comment  `json:"comments" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Likes       []Like     `json:"likes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DisLike     []DisLike  `json:"dis_like" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Mentions    []Mention  `json:"mentions" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Notifes     []Notify   `json:"notifes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Views       []View     `json:"views" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
