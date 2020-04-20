package entity

import "time"

// 权限表(菜单)
type AdministratorPermissions struct {
	Id uint `gorm:"primary_key"`
	PermissionName string `gorm:"type:varchar(60);not null;default:'';comment:'菜单名称'"`
	Key string `gorm:"type:varchar(50);not null;comment:'唯一标识字段，与vue路由name一致'"`
	Method string `gorm:"type:varchar(20);not null;comment:'http请求方法'"`
	Url string	`gorm:"type:varchar(200);not null;comment:'路由'"`
	Description string 	`gorm:"type:varchar(200);not null;default:'';comment:'描述'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
func (AdministratorPermissions) TableName() string {
	return "administrator_permission"
}