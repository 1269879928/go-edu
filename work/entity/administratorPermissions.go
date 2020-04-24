package entity

import "time"

// 权限表(菜单)
type AdministratorPermissions struct {
	Id int `gorm:"primary_key" json:"id"`
	PermissionName string `gorm:"type:varchar(60);not null;default:'';comment:'菜单名称'" json:"permission_name"`
	UniqueKey string `gorm:"type:varchar(50);not null;unique_index:unq_key;comment:'唯一标识字段，与vue路由name一致'" json:"unique_key"`
	Method string `gorm:"type:varchar(20);not null;comment:'http请求方法'" json:"method"`
	Url string	`gorm:"type:varchar(200);not null;comment:'http路由'" json:"url"`
	Pid int `gorm:"type:int(10);not null;default 0;index:idx_pid;comment:'父级菜单id，0为顶级菜单'" json:"pid"`
	Description string 	`gorm:"type:varchar(200);not null;default:'';comment:'描述'" json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Children []AdministratorPermissions `json:"children,omitempty"`
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
	Pid int
}
func (AdministratorPermissionsData) TableName() string {
	return "administrator_permission"
}
type AdministratorPermissionsTree struct {
	Id int `json:"id"`
	PermissionName string `json:"permission_name"`
	UniqueKey string `json:"unique_key"`
	Method string `json:"method"`
	Url string `json:"url"`
	Pid int `json:"pid"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Children []*AdministratorPermissionsTree `json:"children,omitempty"`
}