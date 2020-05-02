package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-edu/docs"
	"go-edu/routes"
	"go-edu/work/base/inits"
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
	r := routes.Router()
	url := ginSwagger.URL("http://localhost:3000/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Run(":3000")
}
