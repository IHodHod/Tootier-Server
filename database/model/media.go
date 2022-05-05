package model

// dl.tootier.ir/tweet/media/voice/shhriyar2000-14001-13-8-13-34-50.mp3
// dl.tootier.ir/tweet/media/image/shhriyar2000-14001-13-8-13-34-50.webp
type Media struct {
	MediaID   uint64  `json:"media_id" gorm:"primaryKey" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	VoiceURL  string  `json:"voice_url" faker:"sentence"`
	ImageURL  string  `json:"image_url" faker:"sentence"`
	TootiID   uint64  `json:"tooti_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	Tooti     Tooti   `json:"tooti" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	QouteID   uint64  `json:"qoute_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	Qoute     Qoute   `json:"qoute" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CommentID uint64  `json:"comment_id" faker:"oneof: 1, 2, 3,4,5,6,7,8,9,10"`
	Comment   Comment `json:"comment" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
