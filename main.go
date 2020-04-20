package main

import (
	"fmt"
	"go-edu/router"
	"go-edu/work/base/inits"
	"os"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	inits.Init()
	mysqlDsn := os.Getenv("MysqlDSN")
	fmt.Println("mysql dsn", mysqlDsn)
	r := router.Router()
	r.Run(":3000")
}
