package dao

import (
	"go-edu/work/base/inits"
	"go-edu/work/entity"
)

type AdministratorRoleRelation struct {
	AdministratorsId []uint64
}
func (d * AdministratorRoleRelation)DeleteByAdminId()  {
	for _, id := range d.AdministratorsId {
		inits.Gorm.Delete(&entity.AdministratorRoleRelation{AdministratorsId: id})
	}

}
