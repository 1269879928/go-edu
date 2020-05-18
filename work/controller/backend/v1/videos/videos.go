package videos

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-edu/work/common"
	"go-edu/work/httpStatus"
	"go-edu/work/serializer"
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

func AliyunAuthTokenRefresh(c *gin.Context)  {

	videoId := c.Param("video_id")

	if len(videoId) == 0 {
		fmt.Println("invalid videoId")
		resp := serializer.Response{
			Code:  httpStatus.PARAM_WRONG,
			Msg:   "video_id is null",
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	service := &services.AliyunVodUploadRefresh{VideoId: videoId}
	resp := service.AliyunAuthTokenRefresh()
	c.JSON(http.StatusOK, resp)
}
func Create(c *gin.Context)  {
	service := &services.CreateVideosService{}
	if err := c.ShouldBind(service); err !=nil {
		fmt.Println(err)
		resp :=common.ValidateResponse(err)
		c.JSON(http.StatusOK, resp)
		return
	}
	resp := service.Create()
	c.JSON(http.StatusOK, resp)
}
func Index(c *gin.Context) {
	service := &services.IndexVideosService{}
	if err := c.ShouldBind(service);err != nil {
		resp := common.ValidateResponse(err)
		fmt.Println(err)
		c.JSON(http.StatusOK, resp)
		return
	}
	resp := service.Index()
	c.JSON(http.StatusOK, resp)
}