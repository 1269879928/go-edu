package upload

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-edu/work/base/inits"
	"go-edu/work/common"
	"os"
	"path"
	"strings"
	"time"
)
// 单文件上传
type UploadFile struct {
	Context   *gin.Context
	File      string // form file name
	BasePath string // upload/tinymce ->tinymce
	AcceptExt []string // 接受上传的文件后缀 ["jpg", "png", "gif"]
	StoreType int // 存储类型（1:本地，2：七牛云，3阿里云）
}


func (ts *UploadFile)Upload() (savePath string, err error) {
	file, err := ts.Context.FormFile(ts.File)
	if err != nil {
		return
	}
	fileExt := strings.ToLower(path.Ext(file.Filename))[1:]
	marker := false
	for _, ext := range ts.AcceptExt {
		if fileExt == ext {
			marker = true
		}
	}
	if !marker {
		err = errors.New("unsupported file type ." + fileExt)
		return
	}
	fileName := fmt.Sprintf("%s%s%s.%s", ts.BasePath, time.Now().Format("20060102"), common.RandomString(8),fileExt)
	err = os.MkdirAll(common.BASE_PATH+ts.BasePath, 0755)
	if err != nil {
		fmt.Println("mkdir for record dir error,", err)
		return
	}
	localFile := common.BASE_PATH  + fileName
	err = ts.Context.SaveUploadedFile(file, localFile)

	if err != nil {
		return
	}
	storeType := 1
	if ts.StoreType == 0 {
		storeType = 1
	} else {
		storeType = ts.StoreType
	}
	if storeType == common.LOCAL_STORE {
		savePath = localFile
	}
	if storeType == common.QINIU_STORE {
		var uploadObj Uploader
		uploadObj = NewQiniuUpload(localFile, fileName, inits.Config.Qiniu.AccessKey, inits.Config.Qiniu.SecretKey, inits.Config.Qiniu.Bucket)
		savePath ,err = uploadObj.Upload()
		if err == nil {
			os.Remove(localFile)
		}
	}
	if storeType == common.ALIYUN_STORE {
		// TODO
		os.Remove(localFile)
	}
	return
}

func errorHandler(err error)  {

}