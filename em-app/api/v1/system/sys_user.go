package system

import (
	"em-app/model/common/response"
	"em-app/model/system"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

func (u *UserApi) Login(c *gin.Context) {
	var user system.User
	_ = c.ShouldBind(&user)
	if user.Username == "" || user.Password == "" {
		response.FailWithMessage("请输入账号或密码", c)
		return
	}
	userInfo, err := userService.VerifyPassword(user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed(userInfo, "登录验证成功", c)
	}
}
