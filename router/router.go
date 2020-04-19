package router

import (
	"github.com/gin-gonic/gin"
	"go-edu/work/controller/backend/v1/administrator"
	"go-edu/work/controller/backend/v1/permissions"
	"go-edu/work/middlewares"
)

func Router() (r *gin.Engine)  {
	r = gin.Default()
	r.Use(middlewares.Cors())
	//r.POST("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})
	// 后端相关
	v1 := r.Group("/backend/v1")
	{
		v1.POST("/administrator/login", administrator.Login)
		v1.Use(middlewares.AuthRequired())
		{
			v1.GET("/administrator", administrator.Index)
			v1.POST("/administrator", administrator.Create)
			v1.PATCH("/administrator/status", administrator.UpdateStatus)
			// 获取权限
			v1.GET("/permissions", permissions.GetPermissions)
		}
	}
	return
}
