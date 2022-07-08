package system

import (
	v1 "em-app/api/v1"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouterWithoutRecord := Router.Group("user")
	userApi := v1.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouterWithoutRecord.POST("login", userApi.Login)
	}
}
