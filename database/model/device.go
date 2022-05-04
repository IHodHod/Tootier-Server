package model

import "time"

type Device struct {
	DeviceID           uint64 `json:"device_id" gorm:"primaryKey"`
	Hardware           `gorm:"embedded"`
	Token              string     `json:"token"`
	LastTimeVisit      time.Time  `json:"last_time_visit"`
	RegisterTimeDevice time.Time  `json:"register_time_device"`
	UserID             uint64     `json:"user_id"`
	Tooties            []Tooti    `json:"tooties" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // has many
	Qoutes             []Qoute    `json:"qoutes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`  // has many
	ReTooties          []*ReTooti `json:"re_tooties" gorm:"many2many:retooti_devices"`                   // many to many
}

type Hardware struct {
	DeviceName string `json:"device_name"`
	MacAddress string `json:"mac_address" gorm:"unique"`
	OS         string `json:"os"`
	IP         string `json:"ip"`
	DeviceUUID string `json:"device_uuid" gorm:"unique"` // FireBase
}
