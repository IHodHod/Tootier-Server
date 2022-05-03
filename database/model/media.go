package model


type VoiceURL string
type ImageURL string


// dl.tootier.ir/tweet/media/voice/shhriyar2000-14001-13-8-13-34-50.mp3
// dl.tootier.ir/tweet/media/image/shhriyar2000-14001-13-8-13-34-50.webp


type Media struct {
	MediaID uint64 `json:"media_id" gorm:"primaryKey"`
	VoiceURL `json:"voice_url"` // TODO non null
	ImageURL `json:"image_url"` // TODO null
}
