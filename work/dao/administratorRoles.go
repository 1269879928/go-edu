package dao

import (
	"go-edu/work/base/inits"
	"go-edu/work/entity"
)

type administratorRolesDao struct {}
var AdministratorRoles *administratorRolesDao

func (*administratorRolesDao) Create(data *entity.AdministratorRoles)(err error)  {
	err = inits.Gorm.Model(&entity.AdministratorRoles{}).Create(data).Error
	return
}
func (*administratorRolesDao) GetById(id int)(data entity.AdministratorRoles,err error)  {
	err = inits.Gorm.Select("id, role_name, description").Where("id = ?", id).First(&data).Error
	return
}
func (*administratorRolesDao) GetByPage(page, pageSize int)(list []entity.AdministratorRoles,total int,err error)  {
	tx := inits.Gorm.Model(&entity.AdministratorRoles{})
	err = tx.Select("id, role_name,description, status, created_at").Offset((page-1)*pageSize).Limit(pageSize).Find(&list).Error
	tx.Count(&total)
	return
}
func (*administratorRolesDao) UpdateById(data *entity.AdministratorRolesData)(err error)  {
	err = inits.Gorm.Model(&entity.AdministratorRoles{}).Save(data).Error
	return
}
func (*administratorRolesDao) UpdateStatusById(id, status int)(err error)  {
	err = inits.Gorm.Model(&entity.AdministratorRoles{}).Where("id = ?", id).Update("status", status).Error
	return
}
func (*administratorRolesDao) DeleteById(data *entity.AdministratorRolesData)(err error)  {
	err = inits.Gorm.Create(data).Error
	return
}