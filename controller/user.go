package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid2 "github.com/gofrs/uuid"
	"github.com/pilinux/gorest/core/xlogger"
	"github.com/pilinux/gorest/database"
	"github.com/pilinux/gorest/database/model"
	"github.com/pilinux/gorest/database/transaction"
	"github.com/pilinux/gorest/global"
	"github.com/pilinux/gorest/io_models"
	"github.com/pilinux/gorest/lib/renderer"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

func UserNameAvailable(c *gin.Context) {
	findUserName := io_models.FindUserName{}
	status := global.CreateStatus()
	err := c.BindUri(&findUserName)
	if err != nil {
		status.Code = http.StatusBadRequest
		status.Message = global.GetLang().MSG_ERR
		status.Status = "error"
		renderer.Render(c, status.ToGin(), status.Code)
		return
	}

	result := transaction.GetUserByUserName(findUserName.Username)

	var defined = false
	if result.UserID > 0 {
		defined = true
	}

	status.Data = gin.H{"defined": defined}
	renderer.Render(c, status.ToGin(), status.Code)
}

func Test(c *gin.Context) {
	xlogger.Verbose("verbose")
	xlogger.Info("information test")
	xlogger.Debug("debug tests")
	xlogger.Warn("warning tests")
	xlogger.Err("error test")

	status := global.CreateStatus()
	register := io_models.Register{}

	err := c.BindJSON(&register)
	if err != nil {
		status.Code = http.StatusBadRequest
		status.Message = global.GetLang().MSG_ERR
		status.Status = "error"

		fmt.Println("err" + err.Error())

		renderer.Render(c, status.ToGin(), status.Code)
		return
	}

	user := model.User{
		UserName:    register.UserName,
		Email:       register.Email,
		PhoneNumber: register.PhoneNumber,
	}


	switch transaction.UserIsExsits(&user) {
	case transaction.USERNAME:
		status.Code = http.StatusConflict
		status.Status = "error"
		status.Message = global.GetLang().MSG_ERR_USERNAME_EXISTS
		renderer.Render(c, status.ToGin(), status.Code)
		return
	case transaction.EMAIL:
		status.Code = http.StatusConflict
		status.Status = "error"
		status.Message = global.GetLang().MSG_ERR_EMAIL_EXISTS
		renderer.Render(c, status.ToGin(), status.Code)
		return
	case transaction.PHONE_NUMBER:
		status.Code = http.StatusConflict
		status.Status = "error"
		status.Message = global.GetLang().MSG_ERR_PHONENUMBER_EXISTS
		renderer.Render(c, status.ToGin(), status.Code)
		return
	}

	uuid , err := uuid2.NewV1() ; if err != nil {
		status.Code = http.StatusInternalServerError
		status.Status = "error"
		status.Message = global.GetLang().MSG_ERR_CAN_NOT_CERATE_USER
		renderer.Render(c, status.ToGin(), status.Code)
		return
	}


	user.Password = register.Password
	user.LatitudeX = register.Latitude
	user.LongitudeY = register.Longitude
	user.Name = register.Name
	user.RegisterTime = global.CurrentTimeUnix()


	device := model.Device{
		Hardware:       model.Hardware{
			DeviceName: register.DeviceName,
			MacAddress: register.MacAddress,
			OS:         register.OS,
			IP:         register.Ip,
			DeviceUUID: uuid.String(),
		},
		Token:              "fsfsfbibsfsfsfsfsfr3e33434343434fwifbwfi",
		LastTimeVisit:      global.CurrentTimeUnix(),
		RegisterTimeDevice: global.CurrentTimeUnix(),
		UserID:             user.UserID,
	}

	if !transaction.CreateNewUser(&user , &device) {
		status.Code = http.StatusInternalServerError
		status.Status = "error"
		status.Message = global.GetLang().MSG_ERR_CAN_NOT_CERATE_USER
		renderer.Render(c, status.ToGin(), status.Code)
		return
	}


	// TODO : Send Verify Code To Email OR PhoneNumber


	renderer.Render(c, gin.H{"result": "اوکیه"}, 200)
}

func goid() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func GetUsers(c *gin.Context) {
	db := database.GetDB()
	users := model.User{
		UserID: 1,
	}
	err := db.Select("Devices").Delete(users)

	if err != nil {
		fmt.Println(err)
		renderer.Render(c, gin.H{"status": "error"}, 500)
	}

	//db.Preload(clause.Associations).Find(&users)

	//device := []model.Device{}

	//db.Find(&users).Model()
	//db.Find(&device)

	//:= db.Model(&model.User{}).Select("*" +
	//"").Joins("left join user.user_id on devices.user_id = user.user.id").
	//Scan(&result{})

	renderer.Render(c, gin.H{"status": "success"}, 202)
	//renderer.Render(c, users, http.StatusOK)
}

// CreateUser - POST /users
func CreateUser(c *gin.Context) {
	//db := database.GetDB()
	//user := model.User{}
	//userFinal := model.User{}
	//
	//userIDAuth := middleware.AuthID
	//
	//// does the user have an existing profile
	//if err := db.Where("id_auth = ?", userIDAuth).First(&userFinal).Error; err == nil {
	//	renderer.Render(c, gin.H{"msg": "user profile found, no need to create a new one"}, http.StatusForbidden)
	//	return
	//}
	//
	//// bind JSON
	//if err := c.ShouldBindJSON(&user); err != nil {
	//	renderer.Render(c, gin.H{"msg": "bad request"}, http.StatusBadRequest)
	//	return
	//}
	//
	//// user must not be able to manipulate all fields
	//userFinal.FirstName = user.FirstName
	//userFinal.LastName = user.LastName
	//userFinal.IDAuth = userIDAuth
	//
	//tx := db.Begin()
	//if err := tx.Create(&userFinal).Error; err != nil {
	//	tx.Rollback()
	//	log.WithError(err).Error("error code: 1101")
	//	renderer.Render(c, gin.H{"msg": "internal server error"}, http.StatusInternalServerError)
	//} else {
	//	tx.Commit()
	//	renderer.Render(c, userFinal, http.StatusCreated)
	//}
}

// UpdateUser - PUT /users
func UpdateUser(c *gin.Context) {
	//db := database.GetDB()
	//user := model.User{}
	//userFinal := model.User{}
	//
	//userIDAuth := middleware.AuthID
	//
	//// does the user have an existing profile
	//if err := db.Where("id_auth = ?", userIDAuth).First(&userFinal).Error; err != nil {
	//	renderer.Render(c, gin.H{"msg": "no user profile found"}, http.StatusNotFound)
	//	return
	//}
	//
	//// bind JSON
	//if err := c.ShouldBindJSON(&user); err != nil {
	//	renderer.Render(c, gin.H{"msg": "bad request"}, http.StatusBadRequest)
	//	return
	//}
	//
	//// user must not be able to manipulate all fields
	//userFinal.UpdatedAt = time.Now()
	//userFinal.FirstName = user.FirstName
	//userFinal.LastName = user.LastName
	//
	//tx := db.Begin()
	//if err := tx.Save(&userFinal).Error; err != nil {
	//	tx.Rollback()
	//	log.WithError(err).Error("error code: 1111")
	//	renderer.Render(c, gin.H{"msg": "internal server error"}, http.StatusInternalServerError)
	//} else {
	//	tx.Commit()
	//	renderer.Render(c, userFinal, http.StatusOK)
	//}
}

// AddHobby - PUT /users/hobbies
func AddHobby(c *gin.Context) {
	//db := database.GetDB()
	//user := model.User{}
	//hobby := model.Hobby{}
	//hobbyNew := model.Hobby{}
	//hobbyFound := 0 // default (do not create new hobby) = 0, create new hobby = 1
	//
	//userIDAuth := middleware.AuthID
	//
	//// does the user have an existing profile
	//if err := db.Where("id_auth = ?", userIDAuth).First(&user).Error; err != nil {
	//	renderer.Render(c, gin.H{"msg": "no user profile found"}, http.StatusForbidden)
	//	return
	//}
	//
	//// bind JSON
	//if err := c.ShouldBindJSON(&hobby); err != nil {
	//	renderer.Render(c, gin.H{"msg": "bad request"}, http.StatusBadRequest)
	//	return
	//}
	//
	//if err := db.Where("hobby = ?", hobby.Hobby).First(&hobbyNew).Error; err != nil {
	//	hobbyFound = 1 // create new hobby
	//}
	//
	//if hobbyFound == 1 {
	//	hobbyNew.Hobby = hobby.Hobby
	//	tx := db.Begin()
	//	if err := tx.Create(&hobbyNew).Error; err != nil {
	//		tx.Rollback()
	//		log.WithError(err).Error("error code: 1121")
	//		renderer.Render(c, gin.H{"msg": "internal server error"}, http.StatusInternalServerError)
	//	} else {
	//		tx.Commit()
	//		hobbyFound = 0
	//	}
	//}
	//
	//if hobbyFound == 0 {
	//	user.Hobbies = append(user.Hobbies, hobbyNew)
	//	tx := db.Begin()
	//	if err := tx.Save(&user).Error; err != nil {
	//		tx.Rollback()
	//		log.WithError(err).Error("error code: 1131")
	//		renderer.Render(c, gin.H{"msg": "internal server error"}, http.StatusInternalServerError)
	//	} else {
	//		tx.Commit()
	//		renderer.Render(c, user, http.StatusOK)
	//	}
	//}
}
