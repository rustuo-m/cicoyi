package route

import (
	"github.com/gin-gonic/gin"
	"github.com/rustuo-m/cicoyi/cicoyi_user/internal/handlers/userHandlers"
)

func InitUserGroupRouter(engine *gin.RouterGroup) {
	user := engine.Group("user")
	{
		user.GET("")
		user.GET("list")
		user.POST("", userHandlers.RegisterUser)
		user.DELETE("")
		user.PUT("")
	}
}
