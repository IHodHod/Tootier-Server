package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid2 "github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	"github.com/pilinux/gorest/core/xlogger"
	"github.com/pilinux/gorest/database/model"
	"github.com/pilinux/gorest/database/transaction"
	"github.com/pilinux/gorest/global"
	"github.com/pilinux/gorest/io_models"
	"github.com/pilinux/gorest/lib/middleware"
	"github.com/pilinux/gorest/lib/renderer"
	"net/http"
	"time"
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

func Signup(c *gin.Context) {
	xlogger.Verbose("verbose")
	xlogger.Info("information test")
	xlogger.Debug("debug tests")
	xlogger.Warn("warning tests")
	xlogger.Err("error test")

	status 	 := global.CreateStatus()
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

	// check this user is exsits or not before new create user account
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

	uuid, err := uuid2.NewV1()
	if err != nil {
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

	// expiretionTime for Expire Token until 30 days after created
	expiretionTime := time.Now().Add(global.MONTH_TO_SEC * time.Second)
	deviceUUID := uuid.String()

	// Create the JWT claims, which includes the username and xxlg[deviceUUID] and expiry time
	claims := &io_models.Claims{
		Username: register.UserName,
		XXLG:     deviceUUID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiretionTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(middleware.AccessKey)
	if err != nil {
		status.Code = http.StatusInternalServerError
		status.Status = "error"
		status.Message = global.GetLang().MSG_ERR
		renderer.Render(c, status.ToGin(), status.Code)
		return
	}

	// model device for user signup
	device := model.Device{
		Hardware: model.Hardware{
			DeviceName: register.DeviceName,
			MacAddress: register.MacAddress,
			OS:         register.OS,
			IP:         register.Ip,
			DeviceUUID: deviceUUID,
		},
		Token:              tokenString,
		LastTimeVisit:      global.CurrentTimeUnix(),
		RegisterTimeDevice: global.CurrentTimeUnix(),
		UserID:             user.UserID,
	}

	if !transaction.CreateNewUser(&user, &device) {
		status.Code = http.StatusInternalServerError
		status.Status = "error"
		status.Message = global.GetLang().MSG_ERR_CAN_NOT_CERATE_USER
		renderer.Render(c, status.ToGin(), status.Code)
		return
	}

	// TODO : Send Verify Code To Email OR PhoneNumber

	status.Code = http.StatusCreated
	status.Data = gin.H{"token": tokenString}
	status.Message = global.GetLang().MSG_SUCCESS_SIGNUP

	renderer.Render(c, status.ToGin(), status.Code)
}

func Login(c *gin.Context) {
	login := io_models.Login{}
	status := global.Status{}

	err := c.BindJSON(&login); if err != nil {
		xlogger.Err("issue in binding json login " + err.Error())
		status.Set("error" , global.GetLang().MSG_ERR , http.StatusBadRequest , nil)
		renderer.Render(c , status.ToGin() , status.Code)
		return
	}

	xlogger.Info(global.STR(login))
	user := transaction.GetUserByUsernameAndPassword(login.UserName , login.Password)
	if user.UserID < 1 {
		xlogger.Err("issue in get user by username and pass ")
		status.Set("error" , global.GetLang().MSG_ERR_LOGIN , http.StatusNotFound , nil)
		renderer.Render(c , status.ToGin() , status.Code)
		return
	}

	uuid, err := uuid2.NewV1()
	if err != nil {
		status.Set("error" , global.GetLang().MSG_ERR_CAN_NOT_CERATE_USER , http.StatusInternalServerError , nil)
		renderer.Render(c, status.ToGin(), status.Code)
		return
	}

	// expiretionTime for Expire Token until 30 days after created
	expiretionTime := time.Now().Add(global.MONTH_TO_SEC * time.Second)
	deviceUUID := uuid.String()

	// Create the JWT claims, which includes the username and xxlg[deviceUUID] and expiry time
	claims := &io_models.Claims{
		Username: login.UserName,
		XXLG:     deviceUUID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiretionTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256 , claims)
	tokenString , err := token.SignedString(middleware.AccessKey); if err != nil {
		status.Set("error" , global.GetLang().MSG_ERR , http.StatusInternalServerError , nil)
		renderer.Render(c, status.ToGin(), status.Code)
		return
	}

	device := model.Device{
		Hardware: model.Hardware{
			DeviceName: login.DeviceName,
			MacAddress: login.MacAddress,
			OS:         login.OS,
			IP:         login.Ip,
			DeviceUUID: deviceUUID,
		},
		Token:              tokenString,
		LastTimeVisit:      global.CurrentTimeUnix(),
		RegisterTimeDevice: global.CurrentTimeUnix(),
		UserID:             user.UserID,
	}

	if !transaction.UpdateGeoLocation(&user , login.Longitude , login.Latitude) {
		status.Set("error" , global.GetLang().MSG_ERR , http.StatusInternalServerError , nil)
		renderer.Render(c , status.ToGin() , status.Code)
		return
	}

	if !transaction.AddDevice(&user , &device) {
		status.Set("error" , global.GetLang().MSG_ERR , http.StatusInternalServerError , nil)
		renderer.Render(c , status.ToGin() , status.Code)
		return
	}

	status.Set("success" , global.GetLang().MSG_SUCCESS_LOGIN , http.StatusOK , gin.H{"token": tokenString ,"user" : user})
	renderer.Render(c , status.ToGin() , status.Code)
}
