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
	Role []*Role `gorm:"many2many:user_roles;"`
	ProfileID uint64 `json:"profile_id"`
	Profile Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Device []Device `json:"devices" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Bookmark []Bookmark `json:"bookmarks" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Following []Following `json:"followings" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FollowingUser []Following `json:"following_users" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:FollowingUserID"`

	Follower 	 []Follower  `json:"followers" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FollowerUser []Follower  `json:"follower_users" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:FollowerUserID"`
	//IDAuth    uint64         `json:"-"`
}

type Role struct {
	RoleID uint64 `json:"role_id" gorm:"primaryKey"`
	RoleName string `json:"role_name"`
	User []*User `gorm:"many2many:user_roles;"`
}

type UserGeo struct {
	LatitudeX string `json:"latitude_x"`
	LongitudeY string `json:"longitude_y"`
}
