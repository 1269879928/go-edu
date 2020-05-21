package migration

import (
	"github.com/shijting/go-edu/work/base/inits"
	"github.com/shijting/go-edu/work/entity"
)

// 数据迁移
func Migration()  {
	inits.Gorm.
		Set("gorm:table_options", "ENGINE=InnoDB").
		Set("gorm:table_options",  "charset=utf8mb4").
		AutoMigrate(&entity.AdministratorsInfo{}, &entity.AdministratorRoles{},&entity.AdministratorRoleRelation{}, &entity.AdministratorPermissions{}, &entity.AdministratorRolePermissionRelation{})
	inits.Gorm.Model(&entity.AdministratorRoleRelation{}).AddIndex("idx_administrator_role_relation_administrator_id", "AdministratorId").
		AddIndex("idx_administrator_role_relation_role_id", "RoleId")
	inits.Gorm.Model(&entity.AdministratorRolePermissionRelation{}).AddIndex("idx_administrator_role_permission_relation_permission_id", "PermissionId").
		AddIndex("idx_administrator_role_permission_relation_role_id", "RoleId")
}
