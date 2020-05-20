package services

import (
	"go-edu/work/base/inits"
	"go-edu/work/common"
	"go-edu/work/dao"
	"go-edu/work/httpStatus"
	"go-edu/work/serializer"
)

type LoginForm struct {
	Email string `form:"email" binding:"required,email,gt=5,lt=20"`
	Password string `form:"password" binding:"required"`
	Token string `form:"token" binding:"required"`
	RemoteAddr string
}
// 登录
func (l *LoginForm)Login() (result *serializer.Response) {

	admin, err :=dao.AdminstratorObj.GetAdministratorByEmail(l.Email)
	if err != nil || admin.ID == 0{
		result = &serializer.Response{
			Code:  httpStatus.WRONG_EMAIL_PASSWORD,
			Msg:   httpStatus.GetCode2Msg(httpStatus.WRONG_EMAIL_PASSWORD),
		}
		return
	}

	verifyResult := common.CheckPasswordHash(l.Password, admin.Password)
	if !verifyResult{
		result = &serializer.Response{
			Code:  httpStatus.WRONG_EMAIL_PASSWORD,
			Msg:   httpStatus.GetCode2Msg(httpStatus.WRONG_EMAIL_PASSWORD),
		}
		return
	}
	jwt := common.GenJWT(admin.ID, admin.Email, l.RemoteAddr, inits.Config.Jwt.Expires)
	result = &serializer.Response{
		Code:  httpStatus.SUCCESS_STATUS,
		Msg:   "login success",
		Data: map[string]interface{}{"token": jwt},
	}
	return
}


