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
	Name string `form:"name" binding:"required,gt=2,lt=10" json:"name"`
	Password string `form:"password" binding:"required,gt=5" json:"password"`
	Email string `form:"email" binding:"required,email" json:"email"`
	RoleId []uint64 `form:"role_id" json:"role_id"`
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
func (admin *Administrator)Create() (res *serializer.Response) {
	pwd, err := common.HashPassword(admin.Password)
	if err != nil {
		res = &serializer.Response{
			Code:  httpStatus.OPERATION_WRONG,
			Data:  nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
		}
		return
	}
	adminData := &entity.Administrators{
		Name:          admin.Name,
		Email:         admin.Email,
		Password:      pwd,
		LastLoginDate: time.Now(),
		Status:        1,
	}
	fmt.Printf("id string%#v\n", admin)
	roleIds := make([]uint64, 0)
	if len(admin.RoleId) > 0 {
		for _, v := range admin.RoleId {
			//id,_ := strconv.ParseUint(v, 10, 64)
			roleIds = append(roleIds, v)
		}
	}
	err = dao.AdminstratorObj.CreateAdministratorRole(adminData, roleIds)
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
	err := dao.AdminstratorObj.Update(s.Id, data)
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
type AdministratorDetail struct {
	Id uint64
}
// 详情
func (d *AdministratorDetail) GetAdministratorDetailById() (resp *serializer.Response)  {
	data, err := dao.AdminstratorObj.GetAdministratorDetailById(&entity.Administrators{ID: d.Id})
	if err != nil {
		resp = &serializer.Response{
			Code:  httpStatus.OPERATION_WRONG,
			Data:  nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
			Error: nil,
		}
		return
	}
	roles, err := dao.AdministratorRoles.GetAllByStatus(1)
	if err != nil {
		resp = &serializer.Response{
			Code:  httpStatus.OPERATION_WRONG,
			Data:  nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
			Error: nil,
		}
		return
	}
	resp = &serializer.Response{
		Code:  httpStatus.SUCCESS_STATUS,
		Data:  map[string]interface{}{"roles": roles, "role":data},
		Msg:   httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
		Error: nil,
	}
	return
}
// 更新
type UpdateForm struct {
	Id       uint64  `form:"id" binding:"required"`
	Name     string `form:"name" binding:"required"`
	Password string `form:"password"`
	Email    string `form:"email" binding:"required,email"`
	RoleId []uint64 `form:"role_id" json:"role_id"`
}

func (d *UpdateForm)Update() (resp *serializer.Response) {
	data := dao.AdministratorUpdate{
		Id:       d.Id,
		Name:     d.Name,
		RoleIds:  d.RoleId,
	}
	if d.Password != "" {
		pwd, err := common.HashPassword(d.Password)
		if err != nil {
			resp = &serializer.Response{
				Code:  httpStatus.OPERATION_WRONG,
				Data:  nil,
				Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
				Error: nil,
			}
			return
		}
		data.Password = pwd
	}
	if err := data.UpdateById(); err != nil {
		resp = &serializer.Response{
			Code:  httpStatus.OPERATION_WRONG,
			Data:  nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
			Error: nil,
		}
		return
	}
	resp = &serializer.Response{
		Code:  httpStatus.SUCCESS_STATUS,
		Data:  nil,
		Msg:   httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
		Error: nil,
	}
	return
}