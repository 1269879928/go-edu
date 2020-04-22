package administrator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-edu/work/common"
	"go-edu/work/dao"
	"go-edu/work/httpStatus"
	"go-edu/work/serializer"
	"go-edu/work/services"
	"net/http"
	"strconv"
	"strings"
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

func Index(c *gin.Context) {
	pageTmp := c.DefaultQuery("page", "1")
	pageSizeTmp := c.DefaultQuery("pageSize", "20")
	result:= services.AdministratorServices.Index(pageTmp, pageSizeTmp)
	c.JSON(http.StatusOK, result)
}

// 添加管理员
func Create(c *gin.Context) {
	service := &services.Administrator{}
	if err := c.ShouldBind(service); err != nil {
		fmt.Printf("params err:%v\n", err)
		c.JSON(http.StatusOK, common.ValidateResponse(err))
		return
	}
	res := service.Create()
	c.JSON(http.StatusOK, res)
}



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

type UpdateForm struct {
	Id       int64  `form:"id" binding:"required"`
	Name     string `form:"name" binding:"required"`
	Password string `form:"password"`
	Email    string `form:"email" binding:"required,email"`
}

func Edit(c *gin.Context) {
	if c.Request.Method == "PATCH" {
		var params UpdateForm
		if err := c.ShouldBindJSON(&params); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, serializer.Response{
				Code: httpStatus.OPERATION_WRONG,
				Data: nil,
				Msg:  httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
			})
			return
		}
		data := make(map[string]interface{})
		data["name"] = params.Name
		password := strings.TrimSpace(params.Password)
		if len(password) > 0 {
			c.JSON(http.StatusOK, serializer.Response{
				Code: httpStatus.PARAM_WRONG,
				Data: nil,
				Msg:  httpStatus.GetCode2Msg(httpStatus.PARAM_WRONG),
			})
		}
		err := dao.AdminstratorObj.UpdateById(params.Id, data)
		if err != nil {
			c.JSON(http.StatusOK, serializer.Response{
				Code: httpStatus.PARAM_WRONG,
				Data: nil,
				Msg:  httpStatus.GetCode2Msg(httpStatus.PARAM_WRONG),
			})
			return
		}
		c.JSON(http.StatusOK, serializer.Response{
			Code: httpStatus.SUCCESS_STATUS,
			Msg:  httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
		})
		return
	}
	idTmp := c.Param("id")
	id, err := strconv.ParseInt(idTmp, 10, 64)
	if err != nil {
		return
	}
	data, err := dao.AdminstratorObj.GetAdministratorById(id)
	if err != nil {
		c.JSON(http.StatusOK, serializer.Response{
			Code: httpStatus.OPERATION_WRONG,
			Data: data,
			Msg:  httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
		})
		return
	}
	c.JSON(http.StatusOK, serializer.Response{
		Code: httpStatus.SUCCESS_STATUS,
		Data: data,
		Msg:  httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
	})
}
