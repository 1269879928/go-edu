package common

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/shijting/go-edu/work/base/inits"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)


//func GenPassword(pwd string) string  {
//	salt := "ljie,er9:09+_LJl"
//	tmpPwd := md5.Sum([]byte(fmt.Sprintf("%s%s", pwd,salt)))
//	return fmt.Sprintf("%x", tmpPwd)
//}
func RandomString(length int) string  {
	seedStr := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	str := ""
	byteStr := []byte(seedStr)
	lenStr := len(seedStr)
	rand.Seed(time.Now().UnixNano())
	for i:= length; i>0; i-- {
		str += string(byteStr[rand.Intn(lenStr)])
	}
	return str
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
type CurrentUserInfo struct {
	Email string `json:"email"`
	UserId uint64 `json:"user_id"`
	Ip string `json:"ip"`
	ExpiresAt int64 `json:"expires_at"`
	Token string `json:"token"`
}
func GenJWT(userId uint64, email string, ip string, expires int64) string  {

	mySigningKey := []byte(inits.Config.Jwt.Key)
	type MyCustomClaims struct {
		Email string `json:"email"`
		UserId uint64 `json:"user_id"`
		Ip string `json:"ip"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		email,
		userId,
		ip,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + expires,
			Issuer:    "shjting",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256 , claims)
	jwt, _ := token.SignedString(mySigningKey)
	return jwt
}
func VerifyJWT(tokenString string) (info *CurrentUserInfo, err error)  {
	type MyCustomClaims struct {
		Email string `json:"email"`
		UserId uint64 `json:"user_id"`
		Ip string `json:"ip"`
		jwt.StandardClaims
	}
	// sample token is expired.  override time so it parses as valid
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(inits.Config.Jwt.Key), nil
	})
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		//fmt.Printf("%#v, %#v\n", claims, token.Raw)
		email := claims.Email
		userId := claims.UserId
		expiresAt :=claims.StandardClaims.ExpiresAt
		ip := claims.Ip
		// 当前token即将过期，生成新token
		if expiresAt - time.Now().Unix() < 20 {
			tokenString = GenJWT(userId, email, ip, inits.Config.Jwt.Expires)
		}
		info = &CurrentUserInfo{
			Email:  email,
			UserId: userId,
			Ip:     ip,
			ExpiresAt: expiresAt,
			Token: tokenString,
		}
	} else {
		err = errors.New("token is Valid")
	}
	return
}
