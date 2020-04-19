package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 跨域中间件
func Cors() gin.HandlerFunc  {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie", "Authorization"}
	//config.AllowOrigins = []string{"http://localhost:9000"}
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	return cors.New(config)
}
