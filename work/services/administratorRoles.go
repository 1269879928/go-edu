package services

import (
	"fmt"
	"go-edu/work/dao"
	"go-edu/work/entity"
	"go-edu/work/httpStatus"
	"go-edu/work/serializer"
)

// 创建角色
type CreateRolesService struct {
	RoleName    string `form:"role_name" binding:"required,gt=2,lt=20" json:"role_name"`
	Description string `form:"description" binding:"lt=200" json:"description"`
}

func (d *CreateRolesService) Create() (resp serializer.Response) {
	data := &entity.AdministratorRoles{
		RoleName:    d.RoleName,
		Description: d.Description,
	}
	err := dao.AdministratorRoles.Create(data)
	if err != nil {
		resp = serializer.Response{
			Code:  httpStatus.OPERATION_WRONG,
			Data:  nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
			Error: nil,
		}
		return
	}
	resp = serializer.Response{
		Code: httpStatus.SUCCESS_STATUS,
		Msg:  httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
	}
	return
}

//
type IndexRolesService struct {
	Page     int
	PageSize int
}

func (d *IndexRolesService) Index() (resp serializer.Response) {

	list, total, err := dao.AdministratorRoles.GetByPage(d.Page, d.PageSize)
	if err != nil {
		resp = serializer.Response{
			Code: httpStatus.GETTING_DATA_FAIL,
			Msg:  httpStatus.GetCode2Msg(httpStatus.GETTING_DATA_FAIL),
		}
		return
	}
	resp = serializer.Response{
		Code: httpStatus.SUCCESS_STATUS,
		Data: map[string]interface{}{"list": list, "total": total},
		Msg:  httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
	}
	return
}

type EditRolesService struct {
	Id int `form:"id" binding:"required"`
}

func (d *EditRolesService) Edit() (resp serializer.Response) {
	data, err := dao.AdministratorRoles.GetById(d.Id)
	fmt.Printf("%#v\n", data)
	if err != nil {
		resp = serializer.Response{
			Code:  httpStatus.GETTING_DATA_FAIL,
			Data:  nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.GETTING_DATA_FAIL),
			Error: nil,
		}
		return
	}
	resp = serializer.Response{
		Code:  httpStatus.SUCCESS_STATUS,
		Data:  data,
		Msg:   httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
		Error: nil,
	}
	return
}

// 更新
type UpdateRolesService struct {
	Id          int64  `form:"id" binding:"required"`
	RoleName    string `form:"role_name" binding:"required,gt=2,lt=30" json:"role_name"`
	Description string `form:"description" binding:"lt=200" json:"description"`
}

func (d *UpdateRolesService) Update() (resp serializer.Response) {
	data := &entity.AdministratorRolesData{
		Id:          d.Id,
		RoleName:    d.RoleName,
		Description: d.Description,
		Status:      1,
	}
	err := dao.AdministratorRoles.UpdateById(data)
	if err != nil {
		resp = serializer.Response{
			Code:  httpStatus.OPERATION_WRONG,
			Data:  nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
			Error: nil,
		}
		return
	}
	resp = serializer.Response{
		Code: httpStatus.SUCCESS_STATUS,
		Msg:  httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
	}
	return
}
// 更新status
type StatusRolesService struct {
	Id          int  `form:"id" binding:"required,number"`
	Status int `form:"status" binding:"number"`
}

func (d *StatusRolesService)UpdateStatus()(resp serializer.Response)  {
	err := dao.AdministratorRoles.UpdateStatusById(d.Id, d.Status)
	if err != nil {
		resp = serializer.Response{
			Code:  httpStatus.OPERATION_WRONG,
			Data:  nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
			Error: nil,
		}
		return
	}
	resp = serializer.Response{
		Code: httpStatus.SUCCESS_STATUS,
		Msg:  httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
	}
	return
}
