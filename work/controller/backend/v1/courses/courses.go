package courses

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

func Create(c *gin.Context)  {
	service := &services.CreateCoursesService{}
	if err := c.ShouldBind(service); err !=nil {
		fmt.Printf("dddddddddddata:%#v\n", service)
		resp :=common.ValidateResponse(err)
		c.JSON(http.StatusOK, resp)
		return
	}
	resp := service.Create()
	c.JSON(http.StatusOK, resp)
}
func Index(c *gin.Context) {
	service := &services.IndexCoursesService{}
	if err := c.ShouldBind(service);err != nil {
		common.ValidateResponse(err)
		return
	}
	resp := service.Index()
	c.JSON(http.StatusOK, resp)
}
func Edit(c *gin.Context)  {
	_id := c.Param("id")
	id,_ := strconv.ParseUint(_id, 10, 64)
	if id == 0 {
		c.JSON(http.StatusOK, serializer.Response{
			Code:  httpStatus.PARAM_WRONG,
			Data:  nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.PARAM_WRONG),
			Error: nil,
		})
		return
	}
	service := &services.EditCourseService{Id: id}
	resp := service.Edit()
	c.JSON(http.StatusOK, resp)
}
func Update(c *gin.Context)  {
	service := &services.UpdateCourseService{}
	if err:= c.ShouldBind(service);err !=nil {
		fmt.Println("err;", err)
		resp := common.ValidateResponse(err)
		c.JSON(http.StatusOK, resp)
		return
	}
	fmt.Printf("%#v\n", service)
	resp := service.Update()
	c.JSON(http.StatusOK, resp)
}