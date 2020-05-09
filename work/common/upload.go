package common

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"strings"
	"time"
)
// 单文件上传
type UploadFile struct {
	Context   *gin.Context
	File      string
	AcceptExt []string // 接受上传的文件后缀 ["jpg", "png", "gif"]
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
	fileName := fmt.Sprintf("%d%d%s", time.Now().Format("20060102"),RandomString(8),fileExt)
	_ = fileName
	return
}

func errorHandler(err error)  {

}