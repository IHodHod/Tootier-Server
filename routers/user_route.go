package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/pilinux/gorest/controller"
)

func UserRouter (c *gin.RouterGroup , handlers ...gin.HandlerFunc) {
	user := c.Group("user" , handlers...) // localhost:300//api/v1/user
	user.GET("username/:username" , controller.UserNameAvailable)
}
