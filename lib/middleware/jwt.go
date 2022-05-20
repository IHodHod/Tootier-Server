package middleware

import (
	"fmt"
	"github.com/pilinux/gorest/core/xlogger"
	"github.com/pilinux/gorest/database/model"
	"github.com/pilinux/gorest/database/transaction"
	"github.com/pilinux/gorest/global"
	"github.com/pilinux/gorest/io_models"
	"github.com/pilinux/gorest/lib/renderer"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// AccessKey ...
var AccessKey []byte

// AccessKeyTTL ...
var AccessKeyTTL int

// RefreshKey ...
var RefreshKey []byte

// RefreshKeyTTL ...
var RefreshKeyTTL int

// MyCustomClaims ...
type MyCustomClaims struct {
	ID    uint64 `json:"Id"`
	Email string `json:"Email"`
	jwt.StandardClaims
}

// JWTPayload ...
type JWTPayload struct {
	AccessJWT  string `json:"AccessJWT"`
	RefreshJWT string `json:"RefreshJWT"`
}

// JWT ...
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		val := c.Request.Header.Get("Authorization")
		if len(val) == 0 || !strings.Contains(val, "Bearer ") {
			xlogger.Debug("token haven't bearer or len is 0")
			// log.Println("no vals or no Bearer found")
			renderer.Render(c , gin.H{"status" : "error" ,
				"code" : http.StatusUnauthorized , "msg" : global.GetLang().MSG_ERR_FORBIDDEN}, http.StatusUnauthorized)
			return
		}
		vals := strings.Split(val, " ")
		if len(vals) != 2 {
			xlogger.Debug("token is can not to split")
			renderer.Render(c , gin.H{"status" : "error" ,
				"code" : http.StatusUnauthorized , "msg" : global.GetLang().MSG_ERR_FORBIDDEN}, http.StatusUnauthorized)
			return
		}

		claims := io_models.Claims{}
		tokenString , err := jwt.ParseWithClaims(vals[1] , &claims , func(token *jwt.Token) (interface{} , error) {
			return AccessKey , nil
		})

		if err != nil {
			xlogger.Info("token is can not to parse " + err.Error() )
			isRemoved := transaction.DeleteDevice(&model.Device{Token: vals[1]})
			xlogger.Info("token is not valid and delete device row -> " + global.STR(isRemoved))
			renderer.Render(c , gin.H{"status" : "error" ,
				"code" : http.StatusUnauthorized , "msg" : global.GetLang().MSG_ERR_FORBIDDEN}, http.StatusUnauthorized)
			return
		}

		if !tokenString.Valid {
			isRemoved := transaction.DeleteDevice(&model.Device{Token: vals[1]})
			xlogger.Info("token is not valid and delete device row -> " + global.STR(isRemoved))
			renderer.Render(c , gin.H{"status" : "error" ,
				"code" : http.StatusUnauthorized , "msg" : global.GetLang().MSG_ERR_FORBIDDEN}, http.StatusUnauthorized)
			return
		}

 		if time.Unix(claims.ExpiresAt , 0).Sub(time.Now()) < global.DAY_TO_SEC * 23 * time.Second{
			expiretionTime := time.Now().Add(global.MONTH_TO_SEC * time.Second)
			xlogger.Info("re new , token should to be right expire -> " + global.STR(claims.ExpiresAt))

			newClaims := io_models.Claims{
				Username:       claims.Username,
				XXLG:           claims.XXLG,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expiretionTime.Unix() ,
				},
			}
			newToken := jwt.NewWithClaims(jwt.SigningMethodHS256 , newClaims)
			newTokenString , err := newToken.SignedString(AccessKey); if err != nil {
				xlogger.Err("can not sign a new token ")
				renderer.Render(c , gin.H{"status" : "error" ,
					"code" : http.StatusInternalServerError , "msg" : global.GetLang().MSG_ERR}, http.StatusInternalServerError)
				return
			}


			if !transaction.DeviceUpdate(vals[1] , newTokenString) {
				xlogger.Err("can not update device ")
				renderer.Render(c , gin.H{"status" : "error" ,
					"code" : http.StatusInternalServerError , "msg" : global.GetLang().MSG_ERR_FORBIDDEN}, http.StatusInternalServerError)
				return
			}

			c.Request.Header.Set("Authorization" , "Bearer " + newTokenString)
			c.Header("Authorization" , newTokenString)
		} else {
			c.Header("Authorization" , vals[1])
		}

		c.Next()
	}
}

// RefreshJWT ...
func RefreshJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var jwtPayload JWTPayload
		if err := c.ShouldBindJSON(&jwtPayload); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		//token, err := jwt.ParseWithClaims(jwtPayload.RefreshJWT, &MyCustomClaims{}, validateRefreshJWT)
		//if err != nil {
		//	c.AbortWithStatus(http.StatusUnauthorized)
		//	return
		//}

		//if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		//	AuthID = claims.ID
		//	Email = claims.Email
		//}
	}
}

// validateAccessJWT ...
func validateAccessJWT(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return AccessKey, nil
}

// validateRefreshJWT ...
func validateRefreshJWT(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return RefreshKey, nil
}

// GetJWT ...
func GetJWT(id uint64, email string, tokenType string) (string, error) {
	var key []byte
	var ttl int
	if tokenType == "access" {
		key = AccessKey
		ttl = AccessKeyTTL
	}
	if tokenType == "refresh" {
		key = RefreshKey
		ttl = RefreshKeyTTL
	}
	// Create the Claims
	claims := MyCustomClaims{
		id,
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(ttl)).Unix(),
			Issuer:    "GoRest API",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtValue, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return jwtValue, nil
}
