package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-edu/work/common"
	"go-edu/work/serializer"
	"net/http"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if len(token) < 20 {
			fmt.Println("1")
			c.Abort()
			c.JSON(http.StatusUnauthorized, serializer.Response{
				Code: 401,
				Msg:  "token is invalid",
			})
			return
		}
		result, err := common.VerifyJWT(token)
		if err != nil {
			fmt.Println("2")
			c.Abort()
			c.JSON(http.StatusUnauthorized, serializer.Response{
				Code: http.StatusUnauthorized,
				Msg:  "token is invalid",
			})
			return
		}
		remoteAddr := c.ClientIP()
		if result.Ip != remoteAddr {
			fmt.Println("3")
			c.Abort()
			c.JSON(http.StatusUnauthorized, serializer.Response{
				Code: 401,
				Msg:  "token is invalid",
			})
			return
		}
		c.Set("Email", result.Email)
		fmt.Println("currentId:", result.UserId)
		c.Set("UserId", result.UserId)
		//c.Set("Token", result.Token)
		c.Next()
		//refreshJwt(c)
	}
}
func refreshJwt(c *gin.Context) {
	fmt.Printf("c::%#v\n", c.ContentType())
	c.JSON(200, gin.H{
		"token": c.Keys["Token"],
	})

}
