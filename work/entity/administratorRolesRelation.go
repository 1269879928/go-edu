package entity
// 管理员角色中间表
type AdministratorRoleRelation struct {
	AdministratorsId uint64 `json:"administrators_id"`
	AdministratorRolesId uint64 `json:"administrator_roles_id"`
}
func (AdministratorRoleRelation) TableName() string {
	return "administrator_roles_relation"
}