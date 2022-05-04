package model

// dl.tootier.ir/tweet/media/voice/shhriyar2000-14001-13-8-13-34-50.mp3
// dl.tootier.ir/tweet/media/image/shhriyar2000-14001-13-8-13-34-50.webp
type Media struct {
	MediaID   uint64 `json:"media_id" gorm:"primaryKey"`
	VoiceURL  string
	TootiID   uint64  `json:"tooti_id"`
	Tooti     Tooti   `json:"tooti" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	QouteID   uint64  `json:"qoute_id"`
	Qoute     Qoute   `json:"qoute" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CommentID uint64  `json:"comment_id"`
	Comment   Comment `json:"comment" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
