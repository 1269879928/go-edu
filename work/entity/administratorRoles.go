package entity

import (
	"time"
)
// 角色表
type AdministratorRoles struct {
	ID          int64 `gorm:"primary_key" json:"id"`
	RoleName    string `gorm:"type:varchar(60);not null;default:'';comment:'角色名称'" json:"role_name"`
	Description string `gorm:"type:varchar(200);not null;default:'';comment:'描述'" json:"description"`
	Status      uint8 `gorm:"type:tinyint(1);not null;default:1;comment:'状态(1正常，0无效)'" json:"status"`
	Permissions []AdministratorPermissions `gorm:"many2many:administrator_role_permission_relation" json:"permissions"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time	`json:"updated_at"`
}
type CreatedAt struct {
	CreatedAt   string `json:"created_at"`
}
func (AdministratorRoles) TableName() string {
	return "administrator_roles"
}
type AdministratorRolesData struct {
	Id int64
	RoleName    string
	Description string
	Status      uint8 `gorm:"default:1"`
}
func (AdministratorRolesData) TableName() string {
	return "administrator_roles"
}

// 事务测试
type Test1 struct {
	Id int64 `gorm:"primary_key"`
	T1Name    string
	T1Age int
}
type Test2 struct {
	Id int64 `gorm:"primary_key"`
	T2Name    string
	T1Id int64
}