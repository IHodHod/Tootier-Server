package model

type Device struct {
	DeviceID           uint64 `json:"device_id" gorm:"primaryKey"`
	Hardware           `gorm:"embedded"`
	Token              string `json:"token"`
	LastTimeVisit      string `json:"last_time_visit"`
	RegisterTimeDevice string `json:"register_time_device"`
	UserID             uint64 `json:"user_id"`
}

type Hardware struct {
	DeviceName string `json:"device_name"`
	MacAddress string `json:"mac_address" gorm:"unique"`
	OS         string `json:"os"`
	IP         string `json:"ip"`
	DeviceUUID string `json:"device_uuid" gorm:"unique"` // FireBase
}
