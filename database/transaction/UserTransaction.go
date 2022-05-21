package transaction

import (
	"github.com/pilinux/gorest/core/xlogger"
	"github.com/pilinux/gorest/database"
	"github.com/pilinux/gorest/database/model"
	"github.com/pilinux/gorest/global"
	"gorm.io/gorm"
)

const (
	EMPTY        = 0
	USERNAME     = 1
	EMAIL        = 2
	PHONE_NUMBER = 3
)

func GetUserByUserName(userName string) model.User {
	var db = database.GetDB()
	user := model.User{}
	db.Find(&user, "user_name = ?", userName)
	return user
}

func GetUserByEmail(email string) model.User {
	var db = database.GetDB()
	user := model.User{}
	db.Find(&user, "email = ?", email)
	return user
}

func GetUserByPhoneNumber(phoneNumber string) model.User {
	var db = database.GetDB()
	user := model.User{}
	db.Find(&user, "phone_number = ?", phoneNumber)
	return user
}

func UserIsExsits(user *model.User) int8 {
	username := GetUserByUserName(user.UserName)
	if username.UserID > 0 {
		return USERNAME
	}

	email := GetUserByEmail(user.Email)
	if email.UserID > 0 {
		return EMAIL
	}

	phoneNumber := GetUserByPhoneNumber(user.PhoneNumber)
	if phoneNumber.UserID > 0 {
		return PHONE_NUMBER
	}

	return EMPTY
}

func UpdateGeoLocation(user *model.User, longitude string, latitude string) bool {
	userGeo := model.UserGeo{
		LatitudeX:  latitude,
		LongitudeY: longitude,
	}

	result := database.GetDB().Model(user).Updates(model.User{
		UserGeo: userGeo,
	})

	if result.RowsAffected > 0 {
		return true
	}

	return false
}

func AddDevice(user *model.User, device *model.Device) bool {
	err := database.GetDB().Model(user).Association("Devices").Append(device)
	if err != nil {
		xlogger.Err("Error in AddDevice " + global.STR(err))
		return false
	}

	return true
}

func CreateNewUser(user *model.User, device *model.Device) bool {
	err := database.GetDB().Transaction(func(tx *gorm.DB) error {
		result := tx.Create(user)
		if result.Error != nil || result.RowsAffected < 1 {
			return result.Error
		}

		device.UserID = user.UserID
		errDevice := tx.Create(device).Error
		if errDevice != nil {
			return errDevice
		}

		return nil
	})

	if err != nil {
		return false
	}

	return true
}

func GetUserByID(id int64) model.User {
	var db = database.GetDB()
	user := model.User{}
	db.Find(&user, id)
	return user
}

func GetUserByUsernameAndPassword(username string, password string) model.User {
	db := database.GetDB()
	user := model.User{}
	db.Find(&user, "user_name = ? AND password = ?", username, password)
	return user
}
