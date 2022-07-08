package system

import (
	"em-app/model/system"
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

type JWTService struct {
}

// 定义加密秘钥
var mySigningKey = []byte("Kay123")

func GenToken(claims system.MyCustomClaims) (string, error) {
	// 使用HS256加密方式
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signToken, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return signToken, nil
}
func ParserToken(signToken string) (*system.MyCustomClaims, error) {
	var claims system.MyCustomClaims
	token, err := jwt.ParseWithClaims(signToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if token.Valid {
		return &claims, nil
	} else {
		return nil, err
	}
	// 带详细错误的解析
	//} else if ve, ok := err.(*jwt.ValidationError); ok {
	//	if ve.Errors&jwt.ValidationErrorMalformed != 0 {
	//		return nil, errors.New("不是一个合法的token")
	//	} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
	//		return nil, errors.New("token过期了")
	//	} else {
	//		fmt.Println("Couldn't handle this token:", err)
	//		return nil, errors.New("无法处理这个token")
	//	}
	//} else {
	//	return nil, errors.New("无法处理这个token")
	//}
}

func RenewToken(claims system.MyCustomClaims) (string, error) {
	// 若token过期但不超过10分钟则给他续签
	if withinLimit(claims.ExpiresAt, 600) {
		claims.ExpiresAt = time.Now().Add(30 * time.Minute).Unix()
		return GenToken(claims)
	}
	return "", errors.New("登录已过期")
}

func withinLimit(at int64, i int64) bool {
	e := time.Now().Unix()
	return e-at < i
}
