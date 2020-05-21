package main

import (
	_ "github.com/shijting/go-edu/docs"
	"github.com/shijting/go-edu/routes"
	"github.com/shijting/go-edu/work/base/inits"
)
// @title 后端相关
// @version 1.0
// @contact.name shjting
// @contact.url 106.53.5.146
// @contact.email 1269879928@qq.com
// @host localhost:3000
// @BasePath /backend/v1
func main() {
	inits.Init()
	r := routes.Routes()
	r.Run(":3000")
	//localFile := "C:\\Users\\shjting\\Desktop\\9999.jpg"
	//saveName := "aaaaaa.jpg"
	//ts := upload.NewQiniuUpload(localFile, saveName, inits.Config.Qiniu.AccessKey, inits.Config.Qiniu.SecretKey, inits.Config.Qiniu.Bucket)
	//ts.Upload()
}
