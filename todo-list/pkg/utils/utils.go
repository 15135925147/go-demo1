package utils

import (
	"github.com/golang-jwt/jwt"
	"time"
)

var JWTsecret = []byte("ABAB")

type Claims struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 签发token 参考https://cloud.tencent.com/developer/article/1770768
func GenerateToken(id uint, username, password string) (string, error) {
	// 设定过期时间
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * 24) //token 24小时失效
	//组织token所需要的参数
	claims := Claims{
		Id:       id,
		UserName: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "todo_list",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JWTsecret)
	return token, err
}

// 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTsecret, nil
	})

	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
