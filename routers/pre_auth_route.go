package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/pilinux/gorest/controller"
)

func PreAuthRouter(c *gin.RouterGroup , handlers ...gin.HandlerFunc) {
	c.GET("find/username/:username" , controller.UserNameAvailable) // api/v1/find/username/:username
}
