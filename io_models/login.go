package io_models


type Login struct {
	UserName string `json:"user_name" binding:"required,lowercase"`
	Password string `json:"password" binding:"required"`

	Latitude    string `json:"latitude" binding:"required,latitude"`
	Longitude   string `json:"longitude" binding:"required,longitude"`
	Name        string `json:"name" binding:"required"`
	Ip          string `json:"ip" binding:"required,ip"`
	OS          string `json:"os" binding:"required"`
	MacAddress  string `json:"mac_address" binding:"required,mac"`
	DeviceName  string `json:"device_name" binding:"required"`
}
