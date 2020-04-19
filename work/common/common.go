package common

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

//func GenPassword(pwd string) string  {
//	salt := "ljie,er9:09+_LJl"
//	tmpPwd := md5.Sum([]byte(fmt.Sprintf("%s%s", pwd,salt)))
//	return fmt.Sprintf("%x", tmpPwd)
//}
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
}
func GenJWT(userId uint64, email string, ip string) string  {
	key := "ljie,er9:09+_LJl-05l"
	mySigningKey := []byte(key)
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
			ExpiresAt: time.Now().Unix() + 60,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256 , claims)
	jwt, err := token.SignedString(mySigningKey)
	fmt.Println("err:", err)
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
		return []byte("ljie,er9:09+_LJl-05l"), nil
	})
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		info = &CurrentUserInfo{
			Email:  claims.Email,
			UserId: claims.UserId,
			Ip:     claims.Ip,
			ExpiresAt: claims.StandardClaims.ExpiresAt,
		}
	} else {
		err = errors.New("token is Valid")
	}
	return
}
