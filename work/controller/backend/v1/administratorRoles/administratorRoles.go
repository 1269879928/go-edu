package administratorRoles

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shijting/go-edu/work/common"
	"github.com/shijting/go-edu/work/dao"
	"github.com/shijting/go-edu/work/httpStatus"
	"github.com/shijting/go-edu/work/serializer"
	"github.com/shijting/go-edu/work/services"
	"net/http"
	"strconv"
)


func Index(c *gin.Context) {
	pageTmp := c.DefaultQuery("page", "1")
	pageSizeTmp := c.DefaultQuery("pageSize", "20")
	page,_ := strconv.Atoi(pageTmp)
	pageSize,_ := strconv.Atoi(pageSizeTmp)
	obj := &services.IndexRolesService{
		Page: page,
		PageSize: pageSize,
	}
	result := obj.Index()
	c.JSON(http.StatusOK, result)
}
func GetRoles(c *gin.Context)  {
	roles, err := dao.AdministratorRoles.GetAllByStatus(1)
	if err != nil {
		resp := &serializer.Response{
			Code:  httpStatus.OPERATION_WRONG,
			Data:  nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
			Error: nil,
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	resp := &serializer.Response{
		Code:  httpStatus.SUCCESS_STATUS,
		Data:  roles,
		Msg:   httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
		Error: nil,
	}
	c.JSON(http.StatusOK, resp)
}

// 添加角色
func Create(c *gin.Context) {
	service := &services.CreateRolesService{}
	if err := c.ShouldBind(service); err != nil {
		fmt.Printf("params err:%v\n", err)
		c.JSON(http.StatusOK, common.ValidateResponse(err))
		return
	}
	res := service.Create()
	c.JSON(http.StatusOK, res)
}



func UpdateStatus(c *gin.Context) {
	status := &services.StatusRolesService{}
	if err := c.ShouldBindJSON(status); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, common.ValidateResponse(err))
		return
	}
	res := status.UpdateStatus()
	c.JSON(http.StatusOK, res)
}

func Edit(c *gin.Context)  {
	_id := c.Param("id")
	id,_ := strconv.Atoi(_id)
	service := services.EditRolesService{Id: id}
	resp := service.Edit()
	 c.JSON(http.StatusOK, resp)
}


func Update(c *gin.Context) {
	service := &services.UpdateRolesService{}
	if err := c.ShouldBindJSON(service);err !=nil {
		c.JSON(http.StatusOK, common.ValidateResponse(err))
		return
	}
	resp := service.Update()
	c.JSON(http.StatusOK, resp)
}
// 分配权限
func UpdatePermissions(c *gin.Context)  {
	service := &services.RolePermissionsService{}
	if err := c.ShouldBindJSON(service); err != nil {
		c.JSON(http.StatusOK, common.ValidateResponse(err))
		return
	}
	resp := service.UpdatePermissionsForRole()
	c.JSON(http.StatusOK, resp)
}
