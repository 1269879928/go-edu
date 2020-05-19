package dao

import (
	"go-edu/work/base/inits"
	"go-edu/work/entity"
)

type Courses struct {}
var CoursesObj *Courses

var coursesColumn = "id, title, price,status, description, seo_description, seo_keywords, thumb,thumb_store_type,category_id, published_at, is_rec,created_at"
func (*Courses)Create(data *entity.Courses)(result *entity.Courses, err error)  {
	err = inits.Gorm.Create(&data).Error
	result = data
	return
}
func (*Courses)GetByPaginate(page,pageSize uint64)(res []*entity.Courses, total int, err error)  {
	tx := inits.Gorm.Model(&entity.Courses{})
	err = tx.Select(coursesColumn).Offset((page - 1) * pageSize).Limit(pageSize).Find(&res).Error
	tx.Count(&total)
	return
}
func (*Courses)GetAllByStatus(status uint64)(result []entity.Courses, err error) {
	err = inits.Gorm.Select(coursesColumn).Where("status = ?", status).Find(&result).Error
	return
}
func (*Courses)GetOneById(id uint64)(result entity.Courses, err error) {
	err = inits.Gorm.Select(coursesColumn).Where("id = ?", id).First(&result).Error
	return
}
func (*Courses)Update(id uint64, data map[string]interface{}) (err error) {

	err = inits.Gorm.Model(&entity.Courses{Id: id}).Omit("id").Update(data).Error
	return
}
//func (*CreateCourseCategories)(id uint64, data map[string]interface{}) (err error) {
//
//	err = inits.Gorm.Model(&entity.CourseCategories{Id: id}).Omit("id").Update(data).Error
//	return
//}
// in 范围查询
func (*Courses)GetSomeByIds(ids []uint64) (res []*entity.Courses, err error) {
	err = inits.Gorm.Model(&entity.Courses{}).Select(coursesColumn).Where("id IN (?)", ids).Find(&res).Error
	return
}