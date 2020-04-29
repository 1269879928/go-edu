package main

import (
	"go-edu/router"
	"go-edu/work/base/inits"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	inits.Init()
	r := router.Router()
	r.Run(":3000")
}
