package courseCategoryies

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-edu/work/common"
	"go-edu/work/services"
	"net/http"
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
