package administrator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-edu/work/common"
	"go-edu/work/services"
	"net/http"
)

// @Summary 后台登录接口
// @Tags 登录
// @version 1.0
// @Accept json
// @Param email path string true "email"
// @Param password path string true "password"
// @Success 200 {object} serializer.Response 成功后返回值
// @Failure 500 {object} serializer.Response 失败返回值
// @Router /administrator/login [post]
func Login(c *gin.Context)  {
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
