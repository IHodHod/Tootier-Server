package model


type Like struct {
	LikeID uint64 `json:"like_id" gorm:"primaryKey"`
	UserID uint64 `json:"user_id"` // TODO relation
	TootiLikeTime uint64 `json:"tooti_like_time"`

	TootiID uint64 `json:"tooti_id"` // TODO relation
	QouteID uint64 `json:"qoute_id"` 	  // TODO relation
	ReTootiID uint64 `json:"re_tooti_id"` // TODO relation
}

type DisLike struct {
	DisLikeID uint64 `json:"dis_like_id" gorm:"primaryKey"`
	UserID uint64 `json:"user_id"` // TODO relation
	TootiLikeTime uint64 `json:"tooti_like_time"`

	TootiID uint64 `json:"tooti_id"` // TODO relation
	QouteID uint64 `json:"qoute_id"` 	  // TODO relation
	ReTootiID uint64 `json:"re_tooti_id"` // TODO relation
}
