package entity

// 角色-权限中间表
type AdministratorRolePermissionRelation struct {
	RoleId uint64
	PermissionId uint64
}
func (AdministratorRolePermissionRelation) TableName() string {
	return "administrator_role_permission_relation"
}