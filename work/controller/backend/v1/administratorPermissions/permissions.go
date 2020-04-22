package administratorPermissions

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-edu/work/common"
	"go-edu/work/serializer"
	"go-edu/work/services"
	"net/http"
	"strconv"
)

// 获取当前管理员权限
func GetPermissions(c *gin.Context) {
	var menus []string = []string{"Home", "system", "Administrator", "AdministratorRole", "Roles", "AdministratorPermission"}
	c.JSON(http.StatusOK, serializer.Response{
		Code:  0,
		Data:  menus,
		Msg:   "ok",
		Error: nil,
	})
	return
}

// 首页
func Index(c *gin.Context) {
	pageTmp := c.DefaultQuery("page", "1")
	pageSizeTmp := c.DefaultQuery("pageSize", "20")
	page,_ := strconv.Atoi(pageTmp)
	pageSize,_ := strconv.Atoi(pageSizeTmp)
	obj := &services.IndexPermissionService{
		Page: page,
		PageSize: pageSize,
	}
	result := obj.Index()
	c.JSON(http.StatusOK, result)
}
// 新增
func Create(c *gin.Context)  {
	service := &services.CreatePermissionService{}
	if err := c.ShouldBindJSON(service); err != nil {
		fmt.Printf("params err2:%v\n", err)
		c.JSON(http.StatusOK, common.ValidateResponse(err))
		return
	}
	resp := service.Create()
	c.JSON(http.StatusOK, resp)
}
// 更新
func Edit(c *gin.Context)  {
	_id := c.Param("id")
	id,_ := strconv.Atoi(_id)
	service := services.EditPermissionService{Id: id}
	resp := service.Edit()
	c.JSON(http.StatusOK, resp)
}


func Update(c *gin.Context) {
	service := &services.UpdatePermissionService{}
	if err := c.ShouldBindJSON(service);err !=nil {
		c.JSON(http.StatusOK, common.ValidateResponse(err))
	}
	resp := service.Update()
	c.JSON(http.StatusOK, resp)
}
func Delete(c *gin.Context)  {
	_id := c.Param("id")
	id , _ := strconv.Atoi(_id)
	//if id == 0 {
	//	c.JSON(http.StatusOK, )
	//}
	service := &services.DeletePermissionService{Id: id}
	resp := service.Delete()
	c.JSON(http.StatusOK, resp)
}