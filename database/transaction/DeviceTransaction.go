package transaction

import (
	"github.com/pilinux/gorest/core/xlogger"
	"github.com/pilinux/gorest/database"
	"github.com/pilinux/gorest/database/model"
)

func DeviceUpdate(oldToken string , token string) bool {
	db := database.GetDB()
	device := model.Device{Token: oldToken}
	err := db.Model(&device).Where("token = ?" , oldToken).Update("token" , token).Error ; if err != nil {
		xlogger.Err(err.Error())
		return false
	}

	return true
}

func DeleteDevice(device *model.Device) bool {
	db := database.GetDB()
	err := db.Delete(device , "token = ?" , device.Token).Error ; if err != nil {
		xlogger.Err(err.Error())
		return false
	}


	return  true
}
