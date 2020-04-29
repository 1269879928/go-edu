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
