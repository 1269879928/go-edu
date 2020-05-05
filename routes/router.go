package routes

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-edu/libs/aliVod"
	"go-edu/work/controller/backend/v1/administrator"
	"go-edu/work/controller/backend/v1/administratorPermissions"
	"go-edu/work/controller/backend/v1/administratorRoles"
	"go-edu/work/controller/backend/v1/courseCategoryies"
	"go-edu/work/middlewares"
	"net/http"
)

func Routes() (r *gin.Engine)  {
	r = gin.Default()
	r.Use(middlewares.Cors())

	v1 := r.Group("/backend/v1")
	{
		v1.POST("/administrator/login", administrator.Login)
		v1.GET("/video/auth", func(context *gin.Context) {
			vodClient, err := aliVod.InitVodClient()
			if err != nil {
				fmt.Println("init vodclient failed , err:", err)
				return
			}
			id := "1ab11273da8a4921b021e52b899b87b0"
			res , err := aliVod.MyGetPlayAuth(vodClient, id)
			if err != nil {
				fmt.Println("get player auth failed , err:", err)
				return
			}
			context.JSON(http.StatusOK, res)
		})
		//v1 := r.Group("/")
		v1.Use(middlewares.AuthRequired())
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
			v1.GET("/roles", administratorRoles.GetRoles)
			v1.PATCH("/role", administratorRoles.Update)
			v1.PATCH("/role/status", administratorRoles.UpdateStatus)
			// 分配权限
			v1.PATCH("/role/update-permissions", administratorRoles.UpdatePermissions)
			// 权限
			v1.GET("/permissions", administratorPermissions.GetPermissions)
			v1.GET("/permission", administratorPermissions.Index)
			v1.GET("/permission-list", administratorPermissions.PermissionsList)
			v1.POST("/permission", administratorPermissions.Create)
			v1.GET("/permission/:id/edit", administratorPermissions.Edit)
			v1.PATCH("/permission", administratorPermissions.Update)
			v1.DELETE("/permission/:id", administratorPermissions.Delete)
			v1.GET("/set-permission", administratorPermissions.SetPermission)
			// 课程分类
			v1.POST("/course-categories", courseCategoryies.Create)
			v1.GET("/course-categories", courseCategoryies.Index)
			v1.GET("/course-categories/:id/edit", courseCategoryies.Edit)
			v1.PATCH("/course-categories", courseCategoryies.Update)
			v1.DELETE("/course-categories", courseCategoryies.Delete)

		}
		url := ginSwagger.URL("http://192.168.1.104:3000/swagger/doc.json")
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
		r.POST("/video", func(context *gin.Context) {
			var localFile string = "D:\\bg\\5588.flv"
			// 初始化vod
			vodClient, err := aliVod.InitVodClient()
			if err != nil {
				fmt.Println("init vodclient failed , err:", err)
				return
			}
			// 获取上传凭证
			upload := &aliVod.CreateUploadVideo{
				Client:      vodClient,
				Title:       "毛片",
				Description: "小姐姐很正，我很喜欢",
				CoverURL:    "http://img03.sogoucdn.com/app/a/100520021/8de3c081b9c92c249460c305a934b1f2",
				Tags:        "行为艺术",
			}
			response, err := upload.MyCreateUploadVideo()
			if err != nil {
				fmt.Println("UploadVideo failed , err:", err)
				return
			}
			// 执行成功会返回VideoId、UploadAddress和UploadAuth
			var videoId = response.VideoId
			var uploadAuthDTO aliVod.UploadAuthDTO
			var uploadAddressDTO aliVod.UploadAddressDTO
			var uploadAuthDecode, _ = base64.StdEncoding.DecodeString(response.UploadAuth)
			var uploadAddressDecode, _ = base64.StdEncoding.DecodeString(response.UploadAddress)
			_ = json.Unmarshal(uploadAuthDecode, &uploadAuthDTO)
			_= json.Unmarshal(uploadAddressDecode, &uploadAddressDTO)
			// 使用UploadAuth和UploadAddress初始化OSS客户端
			var ossClient, _ = aliVod.InitOssClient(uploadAuthDTO, uploadAddressDTO)
			// 上传文件，注意是同步上传会阻塞等待，耗时与文件大小和网络上行带宽有关
			aliVod.UploadLocalFile(ossClient, uploadAddressDTO, localFile)
			//MultipartUploadFile(ossClient, uploadAddressDTO, localFile)
			fmt.Println("Succeed, VideoId:", videoId)
		})

	}
	return
}
