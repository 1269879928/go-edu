package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-edu/work/dao"
	"net/http"
)

func Test(c *gin.Context)  {
	res := dao.TestInsert()
	fmt.Printf("%v\n", res)
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
