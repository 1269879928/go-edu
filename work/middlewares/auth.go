package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myedu/work/common"
	"myedu/work/serializer"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("url", c.Request.RequestURI)
		token := c.GetHeader("Authorization")
		if len(token) > 20{
			fmt.Println("1",)
			result, err := common.VerifyJWT(token)
			if err != nil {
				c.Abort()
				c.JSON(401, serializer.Response{
					Code: 401,
					Msg:    "token is invalid",
				})
				return
			}
			remoteAddr := c.ClientIP()
			if result.Ip != remoteAddr {
				c.Abort()
				c.JSON(401, serializer.Response{
					Code: 401,
					Msg:    "token is invalid",
				})
				return
			}
			fmt.Printf("r:%s, s:%s\n", remoteAddr, result.Ip )
			fmt.Printf("info:%#v\n", result)
			c.Set("Email", result.Email)
			c.Set("UserId", result.UserId)
			c.Next()
		} else {
			c.Abort()
			fmt.Println("2")
			c.JSON(401, serializer.Response{
				Code: 401,
				Msg:    "token is invalid",
			})
			return
		}

		//c.Abort()
	}
}
