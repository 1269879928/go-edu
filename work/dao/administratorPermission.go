package dao

import (
	"errors"
	"go-edu/work/base/inits"
	"go-edu/work/entity"
)

type administratorPermissionsDao struct {}
var AdministratorPermissionsObj *administratorPermissionsDao
func (*administratorPermissionsDao)CreatePermission(data *entity.AdministratorPermissions) (err error) {
	err = inits.Gorm.Create(data).Error
	return
}
func (*administratorPermissionsDao)GetPermissionById(id int)(data entity.AdministratorPermissions, err error)  {
	err = inits.Gorm.Model(&entity.AdministratorPermissions{}).Select("id, permission_name, description, method, url, unique_key, created_at, pid").Where("id = ?", id).First(&data).Error
	return
}
func (*administratorPermissionsDao)DeletePermission(data *entity.AdministratorPermissions) (err error)  {
	err = inits.Gorm.Delete(data).Error
	return
}
// excludeId 排除某个id
func  (*administratorPermissionsDao)GetPermissionByPage(excludeId int)(list []entity.AdministratorPermissions,err error)  {
	tx := inits.Gorm.Model(&entity.AdministratorPermissions{}).Select("id, permission_name,description, method, unique_key, url, created_at,pid")
	if  excludeId > 0 {
		tx = tx.Where("id <> ?", excludeId)
	}

	err = tx.Order("id desc").Find(&list).Error
	return
}
func (*administratorPermissionsDao)UpdatePermission(data *entity.AdministratorPermissionsData) (err error)  {
	err = inits.Gorm.Save(data).Error
	return
}
// 分配权限
func (*administratorPermissionsDao)UpdatePermissionsForRole(roleId uint64, permissionIds []uint64) (err error)  {
	tx := inits.Gorm.Begin()
	if err = tx.Error; err != nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = errors.New("update permission failed")
		}
	}()
	if err = tx.Where("administrator_roles_id = ?", roleId).Delete(entity.AdministratorRolePermissionRelation{}).Error; err !=nil {
		tx.Rollback()
		return
	}
	if len(permissionIds) > 0 {
		for _, pid := range permissionIds {
			if err = tx.Create(&entity.AdministratorRolePermissionRelation{AdministratorRolesId: roleId, AdministratorPermissionsId: pid}).Error; err !=nil {
				tx.Rollback()
				return
			}
		}
	}
	tx.Commit()
	return
}
// 获取权限
func (*administratorPermissionsDao)GetPermissionsWithIds(permissionIds []uint64) (res []*entity.AdministratorPermissions, err error)  {
	err = inits.Gorm.Select("id, permission_name, unique_key, method, url, pid").Where("id in (?)", permissionIds).Find(&res).Error
	return
}