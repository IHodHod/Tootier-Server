package controller

import (
	"github.com/gin-gonic/gin"
)

// GetUsers - GET /users
func GetUsers(c *gin.Context) {
	//db := database.GetDB()
	//users := []model.User{}
	//posts := []model.Post{}
	//hobbies := []model.Hobby{}
	//
}

// GetUser - GET /users/:id
func GetUser(c *gin.Context) {
	//db := database.GetDB()
	//id := c.Params.ByName("id")
	//user := model.User{}
	//posts := []model.Post{}
	//hobbies := []model.Hobby{}
	//
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
