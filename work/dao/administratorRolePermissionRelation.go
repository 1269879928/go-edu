package dao

import (
	"go-edu/work/base/inits"
	"go-edu/work/entity"
)

type AdministratorRolePermissionRelation struct {}
var AdministratorRolePermissionRelationObj *AdministratorRolePermissionRelation

func (*AdministratorRolePermissionRelation)GetPermissionByRoleId(roleId uint64) (res []*entity.AdministratorRolePermissionRelation,err error)  {
	err = inits.Gorm.Select("administrator_roles_id, administrator_permissions_id").Where("administrator_roles_id = ?", roleId).Find(&res).Error
	return
}
func (*AdministratorRolePermissionRelation)GetPermissionByRoleIds(roleIds []uint64) (res []*entity.AdministratorRolePermissionRelation,err error)  {
	err = inits.Gorm.Select("administrator_roles_id, administrator_permissions_id").Where("administrator_roles_id in (?)", roleIds).Find(&res).Error
	return
}