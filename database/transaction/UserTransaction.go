package transaction

import (
	"github.com/pilinux/gorest/database"
	"github.com/pilinux/gorest/database/model"
)

const (
	EMPTY = 0
	USERNAME  = 1
	EMAIL  = 2
	PHONE_NUMBER = 3
)


func GetUserByUserName(userName string) model.User {
	var db  = database.GetDB()
	user := model.User{}
	db.Find(&user , "user_name = ?" , userName)
	return user
}

func GetUserByEmail(email string) model.User {
	var db  = database.GetDB()
	user := model.User{}
	db.Find(&user , "email = ?" , email)
	return user
}

func GetUserByPhoneNumber(phoneNumber string) model.User {
	var db  = database.GetDB()
	user := model.User{}
	db.Find(&user , "phone_number = ?" , phoneNumber)
	return user
}

func UserIsExsits(user *model.User) int8 {
	username := GetUserByUserName(user.UserName)
	if (username.UserID > 0) {
		return USERNAME
	}

	email := GetUserByEmail(user.Email)
	if (email.UserID > 0) {
		return EMAIL
	}

	phoneNumber := GetUserByPhoneNumber(user.PhoneNumber)
	if (phoneNumber.UserID > 0) {
		return PHONE_NUMBER
	}


	return EMPTY
}

//func CreateNewUser(user *model.User) bool {
//	var db = database.GetDB()
//	db.Create(user)
//}

func GetUserByID(id int64) model.User {
	var db  = database.GetDB()
	user := model.User{}
	db.Find(&user , id)
	return user
}

