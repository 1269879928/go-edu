package courseCategoryies

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-edu/work/common"
	"go-edu/work/services"
	"net/http"
	"strconv"
)

// 创建分类
func Create(c *gin.Context)  {
	service := &services.CreateCourseCategoriesService{}
	if err := c.ShouldBind(service); err !=nil {
		common.ValidateResponse(err)
		return
	}
	fmt.Printf("%#v\n", service)
	resp := service.Create()
	c.JSON(http.StatusOK, resp)
}

func Index(c *gin.Context) {
	service := &services.IndexCourseCategoriesService{}
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
		common.ValidateResponse(errors.New("id is invalid"))
		return
	}
	service := &services.EditCourseCategoriesService{Id: id}
	resp := service.Edit()
	c.JSON(http.StatusOK, resp)
}
func Update(c *gin.Context)  {
	service := &services.UpdateCourseCategoriesService{}
	if err:= c.ShouldBind(service);err !=nil {
		fmt.Println("err;", err)
		common.ValidateResponse(err)
		return
	}
	fmt.Printf("%#v\n", service)
	resp := service.Update()
	c.JSON(http.StatusOK, resp)
}
func Delete(c *gin.Context)  {
	service := &services.DeleteCourseCategoriesService{}
	if err:= c.ShouldBind(service);err !=nil {
		common.ValidateResponse(err)
		return
	}
	fmt.Printf("%#v\n", service)
	resp := service.Delete()
	c.JSON(http.StatusOK, resp)
}