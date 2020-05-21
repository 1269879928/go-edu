package dao

import (
	"github.com/shijting/go-edu/work/base/inits"
	"github.com/shijting/go-edu/work/entity"
)

type AdministratorRoleRelation struct {
	AdministratorsId []uint64
}
func (d * AdministratorRoleRelation)DeleteByAdminId()  {
	for _, id := range d.AdministratorsId {
		inits.Gorm.Delete(&entity.AdministratorRoleRelation{AdministratorsId: id})
	}
}
// 获取管理员角色
func (d * AdministratorRoleRelation)GetRolesById()  {
	inits.Gorm.Where("")
}