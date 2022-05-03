package model


type VoiceURL string
type ImageURL string


// dl.tootier.ir/tweet/media/voice/shhriyar2000-14001-13-8-13-34-50.mp3
// dl.tootier.ir/tweet/media/image/shhriyar2000-14001-13-8-13-34-50.webp


type Media struct {
	MediaID uint64 `json:"media_id" gorm:"primaryKey"`
	VoiceURL string

	//VoiceURL `json:"voice_url"` // TODO non null
	//ImageURL `json:"image_url"` // TODO null

	TootiID uint64 `json:"tooti_id"`
	Tooti Tooti `json:"tooti" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	QouteID uint64 `json:"qoute_id"`
	Qoute Qoute `json:"qoute" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CommentID uint64 `json:"comment_id"`
	Comment Comment `json:"comment" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
