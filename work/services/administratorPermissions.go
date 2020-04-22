package services

import (
	"fmt"
	"go-edu/work/dao"
	"go-edu/work/entity"
	"go-edu/work/httpStatus"
	"go-edu/work/serializer"
	"strings"
	"time"
)

type CreatePermissionService struct {
	PermissionName string `form:"permission_name" binding:"required,gt=2,lt=30" json:"permission_name"`
	UniqueKey      string `form:"unique_key" binding:"required,gt=2,lt=50" json:"unique_key"`
	Method         string `form:"method" binding:"required,gt=2,lt=10" json:"method"`
	Url            string `form:"url" binding:"required,gt=2,lt=100" json:"url"`
	Description    string `form:"description" binding:"lt=200" json:"description"`
}

func (d *CreatePermissionService) Create() (resp *serializer.Response) {
	data := &entity.AdministratorPermissions{
		PermissionName: d.PermissionName,
		UniqueKey:      d.UniqueKey,
		Method:         strings.ToUpper(d.Method),
		Url:            d.Url,
		Description:    d.Description,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}
	err := dao.AdministratorPermissionsObj.CreatePermission(data)
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
		Code: httpStatus.SUCCESS_STATUS,
		Msg:  httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
	}
	return
}

type IndexPermissionService struct {
	Page     int
	PageSize int
}

func (d *IndexPermissionService) Index() (resp *serializer.Response) {

	list, total, err := dao.AdministratorPermissionsObj.GetPermissionByPage(d.Page, d.PageSize)
	if err != nil {
		fmt.Printf("%#v\n", err)
		resp = &serializer.Response{
			Code: httpStatus.GETTING_DATA_FAIL,
			Msg:  httpStatus.GetCode2Msg(httpStatus.GETTING_DATA_FAIL),
		}
		return
	}
	resp = &serializer.Response{
		Code: httpStatus.SUCCESS_STATUS,
		Data: map[string]interface{}{"list": list, "total": total},
		Msg:  httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
	}
	return
}

type EditPermissionService struct {
	Id int `form:"id" binding:"required"`
}

func (d *EditPermissionService) Edit() (resp *serializer.Response) {
	data, err := dao.AdministratorPermissionsObj.GetPermissionById(d.Id)
	fmt.Printf("%#v\n", data)
	if err != nil {
		resp = &serializer.Response{
			Code:  httpStatus.GETTING_DATA_FAIL,
			Data:  nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.GETTING_DATA_FAIL),
			Error: nil,
		}
		return
	}
	resp = &serializer.Response{
		Code:  httpStatus.SUCCESS_STATUS,
		Data:  data,
		Msg:   httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
		Error: nil,
	}
	return
}

// 更新
type UpdatePermissionService struct {
	Id             int64  `form:"id" binding:"required" json:"id"`
	PermissionName string `form:"permission_name" binding:"required,gt=2,lt=30" json:"permission_name"`
	Description    string `form:"description" binding:"lt=200" json:"description"`
	Method         string `form:"method" binding:"required,gt=2,lt=30" json:"method"`
	Url            string `form:"url" binding:"required,gt=2,lt=200" json:"url"`
	UniqueKey      string `form:"unique_key" binding:"required,gt=2,lt=50" json:"unique_key"`
}

func (d *UpdatePermissionService) Update() (resp *serializer.Response) {
	data := &entity.AdministratorPermissionsData{
		Id:             d.Id,
		PermissionName: d.PermissionName,
		UniqueKey:      d.UniqueKey,
		Method:         d.Method,
		Url:            d.Url,
		Description:    d.Description,
	}
	err := dao.AdministratorPermissionsObj.UpdatePermission(data)
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
		Code: httpStatus.SUCCESS_STATUS,
		Msg:  httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
	}
	return
}

// 删除
type DeletePermissionService struct {
	Id int `form:"id" binding:"gt=0"`
}

func (d *DeletePermissionService)Delete()(resp *serializer.Response)  {
	obj := &entity.AdministratorPermissions{
		Id:             d.Id,
	}
	err := dao.AdministratorPermissionsObj.DeletePermission(obj)
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
		Code: httpStatus.SUCCESS_STATUS,
		Msg:  httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
	}
	return
}