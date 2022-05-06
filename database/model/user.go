package model

// User model - `users` table
type User struct {
	UserID          uint64 `json:"user_id" gorm:"primaryKey"`
	UserName        string `json:"user_name" gorm:"unique"`
	Password        string `json:"password"`
	Name            string `json:"name"`
	Email           string `json:"email" gorm:"unique"`
	PhoneNumber     string `json:"phone_number" gorm:"unique"`
	RegisterTime    string `json:"register_time"`
	UserGeo         `gorm:"embedded"`
	Comments        []Comment        `json:"comments" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CommentLikes    []CommentLike    `json:"comment_likes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CommentDisLikes []CommentDisLike `json:"comment_dis_likes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" `
	Likes           []Like           `json:"likes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DisLikes        []DisLike        `json:"dis_likes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Role            []*Role          `gorm:"many2many:user_roles;" `
	ProfileID       uint64           `json:"profile_id" `
	Profile         Profile          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" `
	Devices          []Device         `json:"devices" `
	Bookmark        []Bookmark       `json:"bookmarks" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Following       []Following      `json:"followings" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FollowingUser   []Following      `json:"following_users" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:FollowingUserID"`
	Follower        []Follower       `json:"followers" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FollowerUser    []Follower       `json:"follower_users" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:FollowerUserID"`
	Mentioners      []Mention        `json:"mentioners" gorm:"foreignkey:MentionerUserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Mentioneds      []Mention        `json:"mentioneds" gorm:"foreignkey:MentionedUserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Notifes         []Notify         `json:"notifes" gorm:"foreignkey:ToUserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Views           []View           `json:"views" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Tooties         []Tooti          `json:"tooties"`
	ReTooties       []ReTooti        `json:"re_tooties"`
	Qoutes          []Qoute          `json:"qoutes"`

	//IDAuth    uint64         `json:"-"`
}

type Role struct {
	RoleID   uint64  `json:"role_id" gorm:"primaryKey"`
	RoleName string  `json:"role_name"`
	User     []*User `gorm:"many2many:user_roles;"`
}

type UserGeo struct {
	LatitudeX  string `json:"latitude_x"`
	LongitudeY string `json:"longitude_y"`
}
