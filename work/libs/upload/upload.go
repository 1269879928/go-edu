package upload

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
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
}

const BasePath = "upload/"
func (ts *UploadFile)Upload() (savePath string, err error) {
	file, err := ts.Context.FormFile(ts.File)
	file.Open()
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
	fileName := fmt.Sprintf("%s%s.%s", time.Now().Format("20060102"), common.RandomString(8),fileExt)
	err = os.MkdirAll(ts.BasePath, 0755)
	if err != nil {
		fmt.Println("mkdir for record dir error,", err)
		return
	}
	savePath = BasePath + ts.BasePath + fileName
	err = ts.Context.SaveUploadedFile(file, savePath)
	//localFile := "C:\\Users\\shjting\\Desktop\\9999.jpg"
	//var uploadObj  Uploader
	//uploadObj = NewQiniuUpload(localFile, savePath, inits.Config.Qiniu.AccessKey, inits.Config.Qiniu.SecretKey, inits.Config.Qiniu.Bucket)
	//savePath ,err = uploadObj.Upload()
	return
}

func errorHandler(err error)  {

}