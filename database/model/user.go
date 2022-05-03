package model

import "time"

// User model - `users` table
type User struct {
	UserID    uint64 `json:"user_id" gorm:"primaryKey"`
	UserName  *string `json:"user_name"`
	Password string `json:"password"`
	Name string         `json:"name"`
	PhoneNumber *string `json:"phone_number"`
	RegisterTime time.Time `json:"register_time"`
	UserGeo `gorm:"embedded"`
	RoleID uint64 `json:"role_id"` // TODO Relation
	ProfileID uint64 `json:"profile_id"` // TODO Relation
	//IDAuth    uint64         `json:"-"`
}

type Role struct {
	RoleID uint64 `json:"role_id" gorm:"primaryKey"`
	RoleName string `json:"role_name"`
}

type UserGeo struct {
	LatitudeX string `json:"latitude_x"`
	LongitudeY string `json:"longitude_y"`
}
