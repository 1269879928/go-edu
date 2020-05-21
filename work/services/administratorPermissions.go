package services

import (
	"fmt"
	"github.com/shijting/go-edu/work/dao"
	"github.com/shijting/go-edu/work/entity"
	"github.com/shijting/go-edu/work/httpStatus"
	"github.com/shijting/go-edu/work/serializer"
	"strings"
	"time"
)

type CreatePermissionService struct {
	PermissionName string `form:"permission_name" binding:"required,gt=1,lt=30" json:"permission_name"`
	UniqueKey      string `form:"unique_key" binding:"required,gt=2,lt=50" json:"unique_key"`
	Method         string `form:"method" binding:"required,gt=2,lt=10" json:"method"`
	Url            string `form:"url" binding:"required" json:"url"`
	Description    string `form:"description" binding:"lt=200" json:"description"`
	Pid            int    `form:"pid" json:"pid"`
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
		Pid:            d.Pid,
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

type PermissionsList struct {
}

// 获取所有权限
func (*PermissionsList) GetPermissions() (resp *serializer.Response) {
	list, err := dao.AdministratorPermissionsObj.GetPermissionByPage(1)
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
		Data: list,
		Msg:  httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
	}
	return
}

type IndexPermissionService struct {
	Page     int
	PageSize int
}

//var tree []entity.AdministratorPermissionsTree
func (d *IndexPermissionService) Index(excludeId int ) (resp *serializer.Response) {

	list, err := dao.AdministratorPermissionsObj.GetPermissionByPage(excludeId)
	if err != nil {
		fmt.Printf("%#v\n", err)
		resp = &serializer.Response{
			Code: httpStatus.GETTING_DATA_FAIL,
			Msg:  httpStatus.GetCode2Msg(httpStatus.GETTING_DATA_FAIL),
		}
		return
	}
	data := buildData(list)
	result := buildTree(0, data)


	resp = &serializer.Response{
		Code: httpStatus.SUCCESS_STATUS,
		Data: map[string]interface{}{"list": result},
		Msg:  httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
	}
	return
}
func getPermissionTree(list []entity.AdministratorPermissions) (tree []*entity.AdministratorPermissionsTree) {
	//var parent []*entity.AdministratorPermissionsTree
	for _, val := range list {
		if val.Pid == 0 {
			node := &entity.AdministratorPermissionsTree{
				Id:             val.Id,
				PermissionName: val.PermissionName,
				UniqueKey:      val.UniqueKey,
				Method:         val.Method,
				Url:            val.Url,
				Pid:            val.Pid,
				Description:    val.Description,
				CreatedAt:      val.CreatedAt,
				UpdatedAt:      val.UpdatedAt,
			}
			tree = append(tree, node)
		}
	}
	return
}
func buildData(list []entity.AdministratorPermissions) map[int]map[int]entity.AdministratorPermissions {
	data := make(map[int]map[int]entity.AdministratorPermissions)
	for _, v := range list {
		id := v.Id                   // 主ID
		pid := v.Pid                 // 父ID
		if _, ok := data[pid]; !ok { // 如果不存在则创建一个新节点
			data[pid] = make(map[int]entity.AdministratorPermissions)
		}
		data[pid][id] = v
	}
	return data
}

// 构建树的结构.
// a. 判断parent_id是否存在.
// b. 如果parent_id存在继续递归.至到data没有找到parent_id节点的数据.
func buildTree(parentId int, data map[int]map[int]entity.AdministratorPermissions) []entity.AdministratorPermissions {
	node := make([]entity.AdministratorPermissions, 0)
	for id, item := range data[parentId] {
		if data[id] != nil {
			item.Children = buildTree(id, data)
		}
		node = append(node, item)
	}
	return node
}

type EditPermissionService struct {
	Id int `form:"id" binding:"required"`
}

func (d *EditPermissionService) Edit() (resp *serializer.Response) {
	data, err := dao.AdministratorPermissionsObj.GetPermissionById(d.Id)
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
	Url            string `form:"url" binding:"required,min=1,lt=200" json:"url"`
	UniqueKey      string `form:"unique_key" binding:"required,gt=2,lt=50" json:"unique_key"`
	Pid            int    `form:"pid" json:"pid"`
}

func (d *UpdatePermissionService) Update() (resp *serializer.Response) {
	data := &entity.AdministratorPermissionsData{
		Id:             d.Id,
		PermissionName: d.PermissionName,
		UniqueKey:      d.UniqueKey,
		Method:         d.Method,
		Url:            d.Url,
		Description:    d.Description,
		Pid:            d.Pid,
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

func (d *DeletePermissionService) Delete() (resp *serializer.Response) {
	obj := &entity.AdministratorPermissions{
		Id: d.Id,
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
type PermissionsWithRoleService struct {
	RoleId uint64 `form:"role_id" binding:"required" json:"role_id"`
}
func (f *PermissionsWithRoleService)GetPermissionsWithRole() (resp *serializer.Response)  {
	list, err := dao.AdministratorPermissionsObj.GetPermissionByPage(0)
	if err != nil {
		fmt.Printf("%#v\n", err)
		resp = &serializer.Response{
			Code: httpStatus.GETTING_DATA_FAIL,
			Msg:  httpStatus.GetCode2Msg(httpStatus.GETTING_DATA_FAIL),
		}
		return
	}
	data := buildData(list)
	permissionList := buildTree(0, data)
	havePermissions, err := dao.AdministratorRolePermissionRelationObj.GetPermissionByRoleId(f.RoleId)
	if err !=nil {
		resp = &serializer.Response{
			Code: httpStatus.OPERATION_WRONG,
			Msg: httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
		}
		return
	}
	resp = &serializer.Response{
		Code: httpStatus.SUCCESS_STATUS,
		Data: map[string]interface{}{"list": permissionList, "permissions": havePermissions},
		Msg:  httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
	}
	return
}
// 获取管理员对应的权限
type AdministratorHasPermissionsService struct {
	Id uint64 // 管理员id
}
func (f *AdministratorHasPermissionsService)GetPermissionById() (resp *serializer.Response) {
	// 获取 当前登录用户扮演的角色id
	admin, err := dao.AdminstratorObj.GetAdministratorDetailById(&entity.Administrators{ID: f.Id})
	if err != nil {
		resp = &serializer.Response{
			Code: httpStatus.OPERATION_WRONG,
			Msg: httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
		}
		return
	}
	roleIds := make([]uint64, 0)
	if len(admin.Roles) > 0 {
		for _, v := range admin.Roles {
			roleIds = append(roleIds, v.ID)
		}
	}
	fmt.Println("roles",roleIds)
	// 获取权限id
	rolesPermissions, err := dao.AdministratorRolePermissionRelationObj.GetPermissionByRoleIds(roleIds)
	if err != nil {
		resp = &serializer.Response{
			Code: httpStatus.OPERATION_WRONG,
			Msg: httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
		}
		return
	}
	permissionIds := make([]uint64, 0)
	if len(rolesPermissions) == 0 {
		return
	}
	for _, v := range rolesPermissions {
		permissionIds = append(permissionIds, v.AdministratorPermissionsId)
	}
	//
	permissionsInfo, err := dao.AdministratorPermissionsObj.GetPermissionsWithIds(permissionIds)
	if err != nil {
		resp = &serializer.Response{
			Code: httpStatus.OPERATION_WRONG,
			Msg: httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
		}
		return
	}
	permissions := make([]string, 0)
	if len(permissionsInfo) == 0 {
		return
	}
	for _, v := range permissionsInfo {
		permissions = append(permissions, v.UniqueKey)
	}
	resp = &serializer.Response{
		Code:  httpStatus.SUCCESS_STATUS,
		Data:  permissions,
		Msg:   "OK",
		Error: nil,
	}
	return
}

