package dao

import (
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