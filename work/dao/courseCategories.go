package dao

import (
	"go-edu/work/base/inits"
	"go-edu/work/entity"
)

type CreateCourseCategories struct {}
var CreateCourseCategoriesObj *CreateCourseCategories

func (*CreateCourseCategories)Create(data *entity.CourseCategories)(result *entity.CourseCategories, err error)  {
	err = inits.Gorm.Create(data).Error
	result = data
	return
}
func (*CreateCourseCategories)GetByPaginate(page,pageSize uint64)(res []*entity.CourseCategories, total int, err error)  {
	tx := inits.Gorm.Model(&entity.CourseCategories{})
	err = tx.Select("id, name, sort,status,created_at").Offset((page - 1) * pageSize).Limit(pageSize).Find(&res).Error
	tx.Count(&total)
	return
}
func (*CreateCourseCategories)GetAll()(res []*entity.CourseCategories, err error)  {
	tx := inits.Gorm.Model(&entity.CourseCategories{})
	err = tx.Select("id, name, sort,status,created_at").Where("status = 1").Find(&res).Error
	return
}
func (*CreateCourseCategories)GetOneById(id uint64)(result entity.CourseCategories, err error) {
	err = inits.Gorm.Select("id, name, sort,status,created_at").Where("id = ?", id).First(&result).Error
	return
}
func (*CreateCourseCategories)Update(id uint64, data map[string]interface{}) (err error) {

	err = inits.Gorm.Model(&entity.CourseCategories{Id: id}).Omit("id").Update(data).Error
	return
}
//func (*CreateCourseCategories)(id uint64, data map[string]interface{}) (err error) {
//
//	err = inits.Gorm.Model(&entity.CourseCategories{Id: id}).Omit("id").Update(data).Error
//	return
//}