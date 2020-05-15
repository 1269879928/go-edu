package common

import (
	"github.com/gin-gonic/gin"
	"go-edu/libs/upload"
	"go-edu/work/common"
	"go-edu/work/serializer"
	"net/http"
)

func UploadImage(c *gin.Context)  {
	uploader := upload.UploadFile{
		Context:   c,
		File:      "file",
		BasePath:  "course/",
		AcceptExt: []string{"jpg", "png"},
		StoreType: common.QINIU_STORE,
	}
	path, err := uploader.Upload()
	if err != nil {
		c.JSON(http.StatusOK, serializer.Response{
			Code:  4000,
			Data:  nil,
			Msg:   "上传失败",
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, serializer.Response{
		Code:  0,
		Data:  map[string]string{"path": "/"+path},
		Msg:   "上传成功",
		Error: nil,
	})
}