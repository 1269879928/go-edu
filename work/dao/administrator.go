package dao

import (
	"go-edu/work/base/inits"
	"go-edu/work/entity"
)

type AdminstratorDao struct {}
var AdminstratorObj AdminstratorDao
func (a *AdminstratorDao) GetAdministratorById(id int64) (result entity.AdministratorsInfo, err error)  {
	err = inits.Gorm.Where("id = ?", id).First(&result).Error
	return
}

func (a *AdminstratorDao) CreateAdministrator(data *entity.Administrator) (result *entity.Administrator, err error) {
	err = inits.Gorm.Create(&data).Error
	result = data
	return
}
func (a *AdminstratorDao)GetAdministratorsByPagination(page, pageSize int64) (result []entity.AdministratorsInfo, count int64, err error)  {
	db := inits.Gorm.Model(&entity.AdministratorsInfo{})
	db.Count(&count)
	err = db.Offset((page-1)*pageSize).Order("id desc").Limit(pageSize).Find(&result).Error
	return
}
func (a *AdminstratorDao)UpdateById(id int64, data map[string]interface{}) (err error)  {
	err = inits.Gorm.Model(&entity.AdministratorsInfo{}).Where("id = ?", id).Update(data).Error
	return
}
func (a *AdminstratorDao)GetAdministratorByEmail(email string) (info entity.AdministratorsInfo, err error)  {
	err = inits.Gorm.Model(&entity.AdministratorsInfo{}).Where("email = ?", email).First(&info).Error
	return
}

