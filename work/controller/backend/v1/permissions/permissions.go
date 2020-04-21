package permissions

import (
	"github.com/gin-gonic/gin"
	"go-edu/work/serializer"
	"net/http"
)

func GetPermissions(c *gin.Context) {
	var menus []string = []string{"Home", "system", "Administrator", "AdministratorRole", "Roles"}
	c.JSON(http.StatusOK, serializer.Response{
		Code:  0,
		Data:  menus,
		Msg:   "ok",
		Error: nil,
	})
	return
}
