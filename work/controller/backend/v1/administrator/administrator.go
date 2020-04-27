package administrator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-edu/work/common"
	"go-edu/work/httpStatus"
	"go-edu/work/serializer"
	"go-edu/work/services"
	"net/http"
	"strconv"
)

// 获取权限
func GetPermissions(c *gin.Context) {
	var menus []string = []string{"Home", "Icons", "tablelist", "TableBasic", "TableSearch"}
	c.JSON(http.StatusOK, serializer.Response{
		Code: httpStatus.SUCCESS_STATUS,
		Data: menus,
		Msg:  httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
	})
	return
}

// 管理员信息
func GetUserInfo(c *gin.Context) {
	tags := []string{"热情", "开放", "执着", "闷骚", "积极"}
	userinfo := make(map[string]interface{})
	userinfo["name"] = "vvvvvp"
	userinfo["desc"] = "执着于理想，纯粹于当下"
	userinfo["email"] = "HeyUI@some.com"
	userinfo["org"] = "某某公司"
	userinfo["dept"] = "某某部门"
	userinfo["title"] = "HeyUI@some.com"
	userinfo["location"] = "深圳"
	userinfo["tags"] = tags
	c.JSON(http.StatusOK, serializer.Response{
		Code: httpStatus.SUCCESS_STATUS,
		Data: userinfo,
		Msg:  httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
	})
	return
}

// 首页
func Index(c *gin.Context) {
	pageTmp := c.DefaultQuery("page", "1")
	pageSizeTmp := c.DefaultQuery("pageSize", "20")
	result:= services.AdministratorServices.Index(pageTmp, pageSizeTmp)
	c.JSON(http.StatusOK, result)
}


// 添加管理员
func Create(c *gin.Context) {
	service := &services.Administrator{}
	if err := c.ShouldBindJSON(service); err != nil {
		fmt.Printf("params err:%v\n", err)
		c.JSON(http.StatusOK, common.ValidateResponse(err))
		return
	}
	res := service.Create()
	c.JSON(http.StatusOK, res)
}


// 删除
func UpdateStatus(c *gin.Context) {
	status := &services.StatusForm{}
	if err := c.ShouldBindJSON(status); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, common.ValidateResponse(err))
		return
	}
	res := status.UpdateStatus()
	c.JSON(http.StatusOK, res)
}


// 更新
func Edit(c *gin.Context) {
	if c.Request.Method == "PATCH" {
		service := &services.UpdateForm{}
		if err := c.ShouldBind(service); err != nil {
			fmt.Printf("%#v\n",err)
			c.JSON(http.StatusOK, common.ValidateResponse(err))
			return
		}
		resp := service.Update()
		c.JSON(http.StatusOK, resp)
		return
	}
	idTmp := c.Param("id")
	id, err := strconv.ParseUint(idTmp, 10, 64)
	if err != nil {
		return
	}
	service := services.AdministratorDetail{Id: id}
	resp := service.GetAdministratorDetailById()
	c.JSON(http.StatusOK, resp)
}
