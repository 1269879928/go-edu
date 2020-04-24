package router

import (
	"github.com/gin-gonic/gin"
	"go-edu/work/controller/backend/v1/administrator"
	"go-edu/work/controller/backend/v1/administratorPermissions"
	"go-edu/work/controller/backend/v1/administratorRoles"
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
		authorized := r.Group("/")
		authorized.Use(middlewares.AuthRequired())
		{
			// 管理员
			v1.GET("/administrator", administrator.Index)
			v1.POST("/administrator", administrator.Create)
			v1.PATCH("/administrator/status", administrator.UpdateStatus)
			v1.GET("/administrator/edit/:id", administrator.Edit)
			v1.PATCH("/administrator/edit", administrator.Edit)
			// 角色
			v1.GET("/role", administratorRoles.Index)
			v1.POST("/role", administratorRoles.Create)
			v1.GET("/role/:id/edit", administratorRoles.Edit)
			v1.PATCH("/role", administratorRoles.Update)
			v1.PATCH("/role/status", administratorRoles.UpdateStatus)
			// 权限
			v1.GET("/permissions", administratorPermissions.GetPermissions)
			v1.GET("/permission", administratorPermissions.Index)
			v1.GET("/permission-list", administratorPermissions.PermissionsList)
			v1.POST("/permission", administratorPermissions.Create)
			v1.GET("/permission/:id/edit", administratorPermissions.Edit)
			v1.PATCH("/permission", administratorPermissions.Update)
			v1.DELETE("/permission/:id", administratorPermissions.Delete)

		}
	}
	return
}
