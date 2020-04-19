package services

import (
	"fmt"
	"myedu/work/common"
	"myedu/work/dao"
	"myedu/work/httpStatus"
	"myedu/work/serializer"
)

type LoginForm struct {
	Email string `form:"email" binding:"required,email,gt=5,lt=20"`
	Password string `form:"password" binding:"required"`
	RemoteAddr string
}
func (l *LoginForm)Login() (result *serializer.Response) {

	admin, err :=dao.AdminstratorObj.GetAdministratorByEmail(l.Email)
	//if err == sql.ErrNoRows {
	//	result = &serializer.Response{
	//		Code:  httpStatus.WRONG_EMAIL_PASSWORD,
	//		Msg:   httpStatus.GetCode2Msg(httpStatus.WRONG_EMAIL_PASSWORD),
	//	}
	//	return
	//}
	if err != nil || admin.ID == 0{
		fmt.Println("err123 ", err)
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
	menus := []string{"Home", "Administrator"}
	jwt := common.GenJWT(admin.ID, admin.Email, l.RemoteAddr)
	result = &serializer.Response{
		Code:  httpStatus.SUCCESS_STATUS,
		Msg:   "login success",
		Data: map[string]interface{}{"menus":menus, "token": jwt},
	}
	return
}


