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
	err = inits.Gorm.Model(&entity.AdministratorPermissions{}).Select("id, permission_name, description, method, url, unique_key, created_at").Where("id = ?", id).First(&data).Error
	return
}
func (*administratorPermissionsDao)DeletePermission(data *entity.AdministratorPermissions) (err error)  {
	err = inits.Gorm.Delete(data).Error
	return
}
func  (*administratorPermissionsDao)GetPermissionByPage(page, pageSize int)(list []entity.AdministratorPermissions,total int,err error)  {
	tx := inits.Gorm.Model(&entity.AdministratorPermissions{})
	err = tx.Select("id, permission_name,description, method, unique_key, url, created_at").Order("id desc").Offset((page-1)*pageSize).Limit(pageSize).Find(&list).Error
	tx.Count(&total)
	return
}
func (*administratorPermissionsDao)UpdatePermission(data *entity.AdministratorPermissionsData) (err error)  {
	err = inits.Gorm.Save(data).Error
	return
}