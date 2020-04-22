package entity

import "time"

// 权限表(菜单)
type AdministratorPermissions struct {
	Id int `gorm:"primary_key" json:"id"`
	PermissionName string `gorm:"type:varchar(60);not null;default:'';comment:'菜单名称'" json:"permission_name"`
	UniqueKey string `gorm:"type:varchar(50);not null;unique_index:unq_key;comment:'唯一标识字段，与vue路由name一致'" json:"unique_key"`
	Method string `gorm:"type:varchar(20);not null;comment:'http请求方法'" json:"method"`
	Url string	`gorm:"type:varchar(200);not null;comment:'http路由'" json:"url"`
	Description string 	`gorm:"type:varchar(200);not null;default:'';comment:'描述'" json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
func (AdministratorPermissions) TableName() string {
	return "administrator_permission"
}

type AdministratorPermissionsData struct {
	Id int64
	PermissionName string
	UniqueKey string
	Method string
	Url string
	Description string
}
func (AdministratorPermissionsData) TableName() string {
	return "administrator_permission"
}