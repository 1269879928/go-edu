package dao

import (
	"errors"
	"github.com/shijting/go-edu/work/base/inits"
	"github.com/shijting/go-edu/work/entity"
)

type AdminstratorDao struct {}
var AdminstratorObj AdminstratorDao
func (a *AdminstratorDao) GetAdministratorById(id int64) (result entity.Administrators, err error)  {
	err = inits.Gorm.Where("id = ?", id).First(&result).Error
	return
}
// 创建管理员
func (a *AdminstratorDao) CreateAdministrator(data *entity.AdministratorInsert) (result *entity.AdministratorInsert, err error) {
	err = inits.Gorm.Create(&data).Error
	result = data
	return
}
// 创建管理员-角色
func (a *AdminstratorDao) CreateAdministratorRole(admin *entity.Administrators, roleIds []uint64) ( err error) {
	tx := inits.Gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("insert failed ")
			tx.Rollback()
		}
	}()
	if err = tx.Error; err != nil {
		err = errors.New("insert failed ")
		return
	}
	if err = tx.Create(&admin).Error; err !=nil {
		err = errors.New("insert failed ")
		tx.Rollback()
		return
	}
	if len(roleIds)>0 {
		for _, v := range roleIds {
			adminRole := &entity.AdministratorRoleRelation{
				AdministratorsId: admin.ID,
				AdministratorRolesId:          v,
			}
			if err = tx.Create(&adminRole).Error; err !=nil {
				tx.Rollback()
				err = errors.New("insert failed ")
				return
			}
		}
	}

	//result = data
	tx.Commit()
	return
}
func (a *AdminstratorDao)GetAdministratorsByPagination(page, pageSize int64) (result []entity.Administrators, count int64, err error)  {
	db := inits.Gorm.Model(&entity.Administrators{})
	db.Count(&count)
	err = db.Offset((page-1)*pageSize).Order("id desc").Limit(pageSize).Find(&result).Error
	return
}

type AdministratorUpdate struct {
	Id uint64
	Name string
	RoleIds []uint64
	Password string
}

func (*AdminstratorDao)Update(id int64, data map[string]interface{}) (err error) {
	err = inits.Gorm.Model(&entity.Administrators{}).Where("id = ?", id).Update(data).Error
	return
}
// 更新
func (d *AdministratorUpdate)UpdateById() (err error)  {
	tx := inits.Gorm.Begin()
	defer func() {
		if r := recover(); r!=nil {
			err = errors.New("update failed")
			tx.Rollback()
		}
	}()
	adminData := make(map[string]interface{})
	adminData["name"] = d.Name
	if d.Password != "" {
		adminData["password"] = d.Password
	}
	if err = tx.Model(&entity.Administrators{}).Where("id = ?", d.Id).Update(adminData).Error;err !=nil {
		tx.Rollback()
		return
	}
	if err = tx.Where("administrators_id = ?", d.Id).Delete(entity.AdministratorRoleRelation{}).Error;err !=nil {
		tx.Rollback()
		return
	}
	if len(d.RoleIds) > 0 {
		for _, roleId := range d.RoleIds {
			adminRole := &entity.AdministratorRoleRelation{
				AdministratorsId: d.Id,
				AdministratorRolesId:          roleId,
			}
			if err = tx.Create(&adminRole).Error; err !=nil {
				tx.Rollback()
				return
			}
		}
	}
	tx.Commit()
	return
}
func (a *AdminstratorDao)GetAdministratorByEmail(email string) (info entity.Administrators, err error)  {
	err = inits.Gorm.Model(&entity.Administrators{}).Where("email = ?", email).First(&info).Error
	return
}
// 根据id  查询管理员-角色 一对多
func (a *AdminstratorDao)GetAdministratorDetailById(data *entity.Administrators)(result *entity.Administrators, err error)  {
	if err = inits.Gorm.Where("status = 1").First(data).Error; err != nil {
		return
	}
	var roles []entity.AdministratorRoles
	err = inits.Gorm.Model(data).Where("status = 1").Related(&roles, "Roles").Error
	data.Roles = roles
	result = data
	return
}

