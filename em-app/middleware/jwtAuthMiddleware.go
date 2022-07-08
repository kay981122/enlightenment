package middleware

import (
	"em-app/model/common/response"
	"em-app/service"
	"em-app/service/system"
	"github.com/gin-gonic/gin"
	"strings"
)

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JWTService

func JWTAuthMiddleWare(c *gin.Context) {
	// 从请求头中取出
	signToken := c.Request.Header.Get("Authorization")
	if signToken == "" {
		response.FailWithMessage("token为空", c)
		c.Abort()
		return
	}
	// 校验token
	myclaims, err := system.ParserToken(signToken)
	if err != nil {
		if strings.Contains(err.Error(), "expired") {
			// 若过期，调用续签函数
			newToken, _ := system.RenewToken(*myclaims)
			if newToken != "" {
				c.Header("token", newToken)
				c.Request.Header.Set("Authorization", newToken)
				c.Next()
				return
			}
		}
		response.FailWithMessage("token校验失败", c)
		c.Abort()
		return
	}
	// 将用户的id放在到请求的上下文c上
	c.Set("userId", myclaims.Id)
	// 后续的处理函数可以用过c.Get("userid")来获取当前请求的id
	c.Next()
}
