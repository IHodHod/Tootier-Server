package model

import "time"

type Device struct {
	DeviceID uint64 `json:"device_id" gorm:"primaryKey"`
	Hardware `gorm:"embedded"`
	Token string `json:"token"`
	LastTimeVisit time.Time `json:"last_time_visit"`
	RegisterTimeDevice time.Time `json:"register_time_device"`
	UserID uint64 `json:"user_id"` // TODO Relation
}


type Hardware struct {
	DeviceName string `json:"device_name"`
	MacAddress string `json:"mac_address"` // TODO unique
	OS string `json:"os"`
	IP string `json:"ip"`
	DeviceUUID string `json:"device_uuid"` // TODO Unique - FireBase
}


