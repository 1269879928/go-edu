package administrator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-edu/work/common"
	"go-edu/work/services"
	"net/http"
)

func Login(c *gin.Context)  {
	//var menus []string = []string{"Home", "Icons", "tablelist", "TableBasic", "TableSearch"}

	login := &services.LoginForm{
		RemoteAddr: c.ClientIP(),
	}
	if err := c.ShouldBind(login); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, common.ValidateResponse(err))
		return
	}
	result := login.Login()
	c.JSON(http.StatusOK, result)
	return
}
