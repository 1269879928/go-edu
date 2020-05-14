package dao

import (
	"go-edu/work/base/inits"
	"go-edu/work/entity"
)

type CourseChapter struct {}
var CourseChapterObj *CourseChapter

//var column = "id, title, price,status, description, seo_description, seo_keywords, thumb,thumb_store_type,category_id, published_at, is_rec,created_at"
func (*CourseChapter)Create(data *entity.CourseChapter)(result *entity.CourseChapter, err error)  {
	err = inits.Gorm.Create(&data).Error
	result = data
	return
}
func (*CourseChapter)GetByPaginate(courseId,page,pageSize uint64)(res []*entity.CourseChapter, total int, err error)  {
	tx := inits.Gorm.Model(&entity.CourseChapter{})
	err = tx.Select("id, title, sort, created_at, updated_at").Where("course_id = ?", courseId).Offset((page - 1) * pageSize).Limit(pageSize).Find(&res).Error
	tx.Count(&total)
	return
}
func (*CourseChapter)GetOneById(id uint64)(result entity.CourseChapter, err error) {
	err = inits.Gorm.Select("id, title, sort,course_id, created_at, updated_at").Where("id = ?", id).First(&result).Error
	return
}
func (*CourseChapter)Update(id uint64, data map[string]interface{}) (err error) {

	err = inits.Gorm.Model(&entity.CourseChapter{Id: id}).Omit("id").Update(data).Error
	return
}
func (*CourseChapter)Delete(data *entity.CourseChapter) (err error) {

	err = inits.Gorm.Delete(data).Error
	return
}
//func (*CreateCourseCategories)(id uint64, data map[string]interface{}) (err error) {
//
//	err = inits.Gorm.Model(&entity.CourseCategories{Id: id}).Omit("id").Update(data).Error
//	return
//}