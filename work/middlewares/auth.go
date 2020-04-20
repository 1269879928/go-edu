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
			c.Abort()
			c.JSON(http.StatusUnauthorized, serializer.Response{
				Code: 401,
				Msg:  "token is invalid",
			})
			return
		}
		result, err := common.VerifyJWT(token)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, serializer.Response{
				Code: http.StatusUnauthorized,
				Msg:  "token is invalid",
			})
			return
		}
		remoteAddr := c.ClientIP()
		if result.Ip != remoteAddr {
			c.Abort()
			c.JSON(http.StatusUnauthorized, serializer.Response{
				Code: 401,
				Msg:  "token is invalid",
			})
			return
		}
		c.Set("Email", result.Email)
		c.Set("UserId", result.UserId)
		//c.Set("Token", result.Token)
		c.Next()
		refreshJwt(c)
		//c.Abort()
	}
}
func refreshJwt(c *gin.Context) {
	fmt.Printf("c::%#v\n", c.ContentType())
	c.JSON(200, gin.H{
		"token": c.Keys["Token"],
	})

}
