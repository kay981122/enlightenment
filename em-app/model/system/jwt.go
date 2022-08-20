package system

import (
	"github.com/golang-jwt/jwt"
)

// MyCustomClaims 第一步：定义结构体
// MyClaims 定义结构体并继承jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们需要额外记录一个username和id字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyCustomClaims struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
