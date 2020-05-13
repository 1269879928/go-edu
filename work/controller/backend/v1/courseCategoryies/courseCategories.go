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


// @Tags 课程分类
// @Summary 创建分类
// @version 1.0
// @Accept json
// @Param name path string true "name"
// @Param sort path int true  "sort"
// @Success 200 {object} serializer.Response 成功后返回值
// @Failure 500 {object} serializer.Response 失败返回值
// @Router /course-categories [post]
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

// @Tags 课程分类
// @Summary 获取课程分类列表(分页)
// @version 1.0
// @Accept json
// @Param page path string true "page"
// @Param pageSize path int true  "pageSize"
// @Success 200 {object} serializer.Response 成功后返回值
// @Failure 500 {object} serializer.Response 失败返回值
// @Router /course-categories [GET]
func Index(c *gin.Context) {
	service := &services.IndexCourseCategoriesService{}
	if err := c.ShouldBind(service);err != nil {
		common.ValidateResponse(err)
		return
	}
	resp := service.Index()
	c.JSON(http.StatusOK, resp)
}
// @Tags 课程分类
// @Summary 获取课程分类信息
// @version 1.0
// @Accept json
// @Success 200 {object} serializer.Response 成功后返回值
// @Failure 500 {object} serializer.Response 失败返回值
// @Router /course-categories/{id}/edit [GET]
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
// @Tags 课程分类
// @Summary 更新课程分类
// @version 1.0
// @Accept json
// @Param id path string true "id"
// @Param name path string true "name"
// @Param sort path int true  "sort"
// @Param status path int true  "status"
// @Success 200 {object} serializer.Response 成功后返回值
// @Failure 500 {object} serializer.Response 失败返回值
// @Router /course-categories [PATCH]
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
// @Tags 课程分类
// @Summary 删除课程分类
// @version 1.0
// @Accept json
// @Param id path string true "id"
// @Param status path int true "status"
// @Success 200 {object} serializer.Response 成功后返回值
// @Failure 500 {object} serializer.Response 失败返回值
// @Router /course-categories [DELETE]
func Delete(c *gin.Context)  {
	service := &services.DeleteCourseCategoriesService{}
	if err:= c.ShouldBind(service);err !=nil {
		common.ValidateResponse(err)
		return
	}
	resp := service.Delete()
	c.JSON(http.StatusOK, resp)
}
// @Tags 课程分类
// @Summary 获取所有课程分类
// @version 1.0
// @Accept json
// @Success 200 {object} serializer.Response 成功后返回值
// @Failure 500 {object} serializer.Response 失败返回值
// @Router /course-categories-all [GET]
func GetAll(c *gin.Context)  {
	service := &services.GetAllCourseCategoriesService{}
	resp := service.GetAllCategories()
	c.JSON(http.StatusOK, resp)
}
