package main

import (
	"fmt"
	"go-edu/router"
	"go-edu/work/base/inits"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	inits.Init()
	r := router.Router()
	fmt.Printf("%#v\n", inits.Config)
	r.Run(":3000")
}
