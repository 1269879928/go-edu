package entity

// 角色-权限中间表
type AdministratorRolePermissionRelation struct {
	RoleId uint 	`gorm:"type:int(10);not null;index:idx_administrator_role_permission_relation_role_id"`
	PermissionId uint `gorm:"type:int(10);not null;index:idx_administrator_role_permission_relation_permission_id"`
}
func (AdministratorRolePermissionRelation) TableName() string {
	return "administrator_roles_permission_relation"
}