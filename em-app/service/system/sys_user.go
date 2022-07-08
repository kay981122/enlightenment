package system

import (
	"em-app/global"
	"em-app/model/system"
	"em-app/utils"
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

type UserService struct {
}

func (userService UserService) VerifyPassword(user system.User) (string, error) {
	var userInfo system.User
	// 加密
	user.Password = utils.EncryptionByMD5(user.Password)
	db := global.Db.Model(&system.User{})
	db.Where("username = ?", user.Username).Where("password = ?", user.Password).Where("status = '1'").Find(&userInfo)
	if userInfo == (system.User{}) {
		return "", errors.New("账号或密码有误")
	}
	token, _ := GenToken(system.MyCustomClaims{
		userInfo.Id,
		userInfo.Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
			Issuer:    "kay",
		}})
	return token, nil
}
