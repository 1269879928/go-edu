package entity
// 管理员角色中间表
type AdministratorRoleRelation struct {
	AdministratorId uint `gorm:"type:int(10);not null;index:idx_administrator_role_relation_administrator_id"`
	RoleId uint `gorm:"type:int(10);not null;index:idx_administrator_role_relation_role_id"`
}
func (AdministratorRoleRelation) TableName() string {
	return "administrator_role_relation"
}