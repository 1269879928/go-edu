package main

import (
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
	r := routes.Routes()
	r.Run(":3000")
}
