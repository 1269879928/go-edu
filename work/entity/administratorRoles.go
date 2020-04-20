package entity

import (
	"time"
)
// 角色表
type AdministratorRoles struct {
	ID          uint `gorm:"primary_key"`
	RoleName    string `gorm:"type:varchar(60);not null;default:'';comment:'角色名称'"`
	Description string `gorm:"type:varchar(200);not null;default:'';comment:'描述'"`
	Status      uint8 `gorm:"type:tinyint(1);not null;default:1;comment:'状态(1正常，0无效)'"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
func (AdministratorRoles) TableName() string {
	return "administrator_roles"
}