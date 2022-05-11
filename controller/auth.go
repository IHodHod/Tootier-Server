package controller

import (
	"github.com/gin-gonic/gin"
)

// CreateUserAuth - POST /register // user => {
//			username , password , name , phoneNumber
// 			os 		 , device_name , mac_address , ip
//  		    , device: {
// 					os : "android" ,
// 					device_name: "" ,
// 					macAddress: "" ,
// 					ip: "",
// 				}
//		}
func CreateUserAuth(c *gin.Context) {
//	db := database.GetDB()
//	user := model.User{}
//	register := io_models.Register{}
//	regissterFinal := io_models.Register{}
//
//	// bind JSON
//	if err := c.ShouldBindJSON(&register); err != nil {
//		renderer.Render(c, gin.H{"status": err.Error()}, http.StatusBadRequest)
//		return
//	}
//
//	// email validation
//	if !service.IsEmailValid(register.Email) {
//		renderer.Render(c, gin.H{"msg": "wrong email address"}, http.StatusBadRequest)
//		return
//	}
//
//	// email must be unique
//	if err := db.Where("email = ?", register.Email).Select("UserID").First(&user).Error; err == nil {
//		renderer.Render(c, gin.H{"msg": "email already registered"}, http.StatusForbidden)
//		return
//	}
//
//	// user must not be able to manipulate all fields
//	authFinal.Email = auth.Email
//	authFinal.Password = auth.Password
//
//	// one unique email for each account
//	tx := db.Begin()
//	if err := tx.Create(&authFinal).Error; err != nil {
//		tx.Rollback()
//		log.WithError(err).Error("error code: 1001")
//		renderer.Render(c, gin.H{"msg": "internal server error"}, http.StatusInternalServerError)
//	} else {
//		tx.Commit()
//		renderer.Render(c, authFinal, http.StatusCreated)
//	}
}
