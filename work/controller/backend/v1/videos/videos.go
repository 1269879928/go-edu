package videos

import (
	"github.com/gin-gonic/gin"
	"go-edu/work/common"
	"go-edu/work/services"
	"net/http"
)

func AliyunVodAuthTokenCreate(c *gin.Context)  {
	service := &services.AliyunVodUploadCreate{}
	if err := c.ShouldBind(service); err !=nil {
		resp :=common.ValidateResponse(err)
		c.JSON(http.StatusOK, resp)
		return
	}
	resp := service.AliyunAuthTokenCreate()
	c.JSON(http.StatusOK, resp)
}
