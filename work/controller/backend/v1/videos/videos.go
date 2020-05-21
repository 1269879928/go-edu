package videos

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shijting/go-edu/work/common"
	"github.com/shijting/go-edu/work/httpStatus"
	"github.com/shijting/go-edu/work/serializer"
	"github.com/shijting/go-edu/work/services"
	"net/http"
	"strconv"
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
		c.JSON(http.StatusOK, resp)
		return
	}
	resp := service.Index()
	c.JSON(http.StatusOK, resp)
}
func Edit(c *gin.Context)  {
	_id := c.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	if err !=nil {
		resp := &serializer.Response{
			Code:  httpStatus.OPERATION_WRONG,
			Data:  nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
			Error: nil,
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	service := &services.EditVideosService{Id: id}
	fmt.Printf("data:%v\n",service)
	resp := service.Edit()
	c.JSON(http.StatusOK, resp)
}
func Update(c *gin.Context) {
	service := &services.UpdateVideosService{}
	if err := c.ShouldBind(service);err != nil {
		resp := common.ValidateResponse(err)
		c.JSON(http.StatusOK, resp)
		return
	}
	resp := service.Update()
	c.JSON(http.StatusOK, resp)
}
func Delete(c *gin.Context) {
	service := &services.DeleteVideosService{}
	if err := c.ShouldBind(service);err != nil {
		resp := common.ValidateResponse(err)
		c.JSON(http.StatusOK, resp)
		return
	}
	resp := service.Delete()
	c.JSON(http.StatusOK, resp)
}