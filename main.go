package main

import (
	"fmt"
	"myedu/router"
	"myedu/work/base/inits"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	inits.Init()
	r := router.Router()
	fmt.Printf("%#v\n", inits.Config)
	r.Run(":3000")
}
