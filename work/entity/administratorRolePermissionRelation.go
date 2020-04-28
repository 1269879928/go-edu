package entity

// 角色-权限中间表
type AdministratorRolePermissionRelation struct {
	AdministratorRolesId uint64 `json:"administrator_roles_id"`
	AdministratorPermissionsId uint64 `json:"administrator_permissions_id"`
}
func (AdministratorRolePermissionRelation) TableName() string {
	return "administrator_role_permission_relation"
}