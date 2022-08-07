package system

import (
	"em-app/global"
	"em-app/model/system"
	"em-app/utils"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

type UserService struct {
}

func (userService UserService) VerifyPassword(user system.User) (map[string]interface{}, error) {
	var userInfo system.User
	// 加密
	user.Password = utils.EncryptionByMD5(user.Password)
	db := global.Db.Model(&system.User{})
	db.Select("id", "username", "nickname", "phone", "email").Where("username = ?", user.Username).Where("password = ?", user.Password).Where("status = '1'").Find(&userInfo)
	if userInfo == (system.User{}) {
		return nil, errors.New("账号或密码有误")
	}
	fmt.Println(userInfo)
	token, _ := GenToken(system.MyCustomClaims{
		userInfo.Id,
		userInfo.Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
			Issuer:    "kay",
		}})
	result := make(map[string]interface{})
	result["userInfo"] = userInfo
	result["token"] = token
	return result, nil
}

func (userService UserService) AddUser(user system.User) (string, error) {
	user.Password = utils.EncryptionByMD5(user.Password)
	currentTime := global.GetCurrentTime(global.TimeTemplates[0])
	user.CreateTime = currentTime
	user.UpdateTime = currentTime
	user.Status = "1"
	mySnow, _ := global.NewSnowFlake(0, 0)
	id, _ := mySnow.NextId()
	user.Id = strconv.FormatInt(id, 10)
	global.Db.Create(&user)
	return user.Id, nil
}
