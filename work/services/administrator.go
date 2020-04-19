package services

import (
	"fmt"
	"go-edu/work/common"
	"go-edu/work/dao"
	"go-edu/work/entity"
	"go-edu/work/httpStatus"
	"go-edu/work/serializer"
	"strconv"
	"time"
)

type Administrator struct {
	Name string `form:"name" binding:"required,gt=2,lt=10"`
	Password string `form:"password" binding:"required,gt=5"`
	Email string `form:"email" binding:"required,email"`
}

type StatusForm struct {
	Id     int64 `form:"id" binding:"required" json:"id"`
	Status int64 `form:"status" json:"status"`
}

var AdministratorServices *Administrator

func (admin *Administrator)Index(pageTmp, pageSizeTmp string) (res *serializer.Response) {
	page, err := strconv.ParseInt(pageTmp, 10, 64)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		// todo
		return
	}
	pageSize, err := strconv.ParseInt(pageSizeTmp, 10, 64)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	data, count, err := dao.AdminstratorObj.GetAdministratorsByPagination(page, pageSize)
	if err!= nil {
		fmt.Printf("err2:%v\n", err)
		res = &serializer.Response{
			Code:  httpStatus.GETTING_DATA_FAIL,
			Data:  nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.GETTING_DATA_FAIL),
		}
		return
	}
	result := map[string]interface{} {"list": data, "total": count}
	res = &serializer.Response{
		Code:  httpStatus.SUCCESS_STATUS,
		Data:  result,
		Msg:   httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
	}
	return
}
func (admin *Administrator)Create() (res *serializer.Response)  {
	pwd, err := common.HashPassword(admin.Password)
	if err != nil {
		res = &serializer.Response{
			Code:  httpStatus.OPERATION_WRONG,
			Data:  nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
		}
		return
	}
	data := &entity.Administrator{
		Name:          admin.Name,
		Email:         admin.Email,
		Password:      pwd,
		LastLoginDate: time.Now(),
		Status:        1,
	}
	_, err = dao.AdminstratorObj.CreateAdministrator(data)
	if err != nil {
		// TODO
		fmt.Printf("err:%#v\n", err)
		res = &serializer.Response{
			Code:  httpStatus.OPERATION_WRONG,
			Data:  nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
		}
		return
	}
	res = &serializer.Response{
		Code:  httpStatus.SUCCESS_STATUS,
		Data:  nil,
		Msg:   httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
	}
	return
}
func (s StatusForm) UpdateStatus() (res *serializer.Response) {
	data := map[string]interface{}{"status": s.Status}
	err := dao.AdminstratorObj.UpdateById(s.Id, data)
	if err != nil {
		res = &serializer.Response{
			Code:  httpStatus.OPERATION_WRONG,
			Data:  nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
		}
		return
	}
	res = &serializer.Response{
		Code:  httpStatus.SUCCESS_STATUS,
		Data:  nil,
		Msg:   httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
	}
	return
}
