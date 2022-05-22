package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/pilinux/gorest/controller"
)

func AuthRouter(c *gin.RouterGroup , handlers ...gin.HandlerFunc){
	auth := c.Group("auth" , handlers...)
	auth.POST("signup" , controller.Signup) // api/v1/test
	auth.POST("login", controller.Login)
}
