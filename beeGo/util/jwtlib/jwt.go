package jwtlib

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//简易Token
type EasyToken struct {
	UserName string //用户名
	Uid      string //用户id
	Expires  int64  //过期时间
}

var (
	verifyKey  string //验证密钥
	ErrAbsent  = "token absent(令牌不存在)"
	ErrInvalid = "token invalid(无效令牌)"
	ErrExpired = "token expired(令牌过期)"
	ErrOther   = "other err(其他错误)"
)

func init() {
	verifyKey = beego.AppConfig.String("jwt::tokenKey")
}

//获取token令牌
func (e EasyToken) GetToken() (string, error) {
	mySigningKey := []byte(verifyKey)
	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: e.Expires,
		Id:        e.Uid,
		Issuer:    e.UserName,
		IssuedAt:  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	return tokenString, err
	//claims := &jwt.StandardClaims{
	//	ExpiresAt: e.Expires,
	//	Issuer:    e.UserName,
	//	IssuedAt:  time.Now().Unix(),
	//}
	//token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	//tokenString, err := token.SignedString([] byte(verifyKey))
	//fmt.Printf("%v%v",tokenString,err)
	//if err != nil {
	//	log.Println(err)
	//}
	//return tokenString, err

	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//	"foo": "bar",
	//	"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	//})
	//
	//// Sign and get the complete encoded token as a string using the secret
	//tokenString, err := token.SignedString([]byte(verifyKey))
	//
	//fmt.Println(tokenString, err)
	//return tokenString, err
}

//验证token
func (e EasyToken) ValidateToken(tokenString string) (bool, error, jwt.MapClaims) {
	if tokenString == "" {
		return false, errors.New(ErrAbsent), nil
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(verifyKey), nil
	})
	if token == nil {
		return false, errors.New(ErrInvalid), nil
	}
	if token.Valid {
		maps := token.Claims.(jwt.MapClaims)
		return true, nil, maps
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, errors.New(ErrInvalid), nil
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return false, errors.New(ErrExpired), nil
		} else {
			return false, errors.New(ErrOther), nil
		}
	} else {
		return false, errors.New(ErrOther), nil
	}
}
