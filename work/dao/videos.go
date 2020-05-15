package dao

import (
	"go-edu/work/base/inits"
	"go-edu/work/entity"
)

type VideosDao struct {}
var VideosDaoObj *VideosDao

//var column = "id, title, price,status, description, seo_description, seo_keywords, thumb,thumb_store_type,category_id, published_at, is_rec,created_at"
func (*VideosDao)Create(data *entity.Videos)(result *entity.Videos, err error)  {
	err = inits.Gorm.Create(&data).Error
	result = data
	return
}
func (*VideosDao)GetByPaginate(courseId,page,pageSize uint64)(res []*entity.Videos, total int, err error)  {
	tx := inits.Gorm.Model(&entity.Videos{})
	err = tx.Select("id, title, sort, created_at, updated_at").Where("course_id = ?", courseId).Offset((page - 1) * pageSize).Limit(pageSize).Find(&res).Error
	tx.Count(&total)
	return
}
func (*VideosDao)GetOneById(id uint64)(result entity.Videos, err error) {
	err = inits.Gorm.Select("id, title, sort,course_id, created_at, updated_at").Where("id = ?", id).First(&result).Error
	return
}
func (*VideosDao)Update(id uint64, data map[string]interface{}) (err error) {
	err = inits.Gorm.Model(&entity.Videos{Id: id}).Omit("id").Update(data).Error
	return
}
func (*VideosDao)Delete(data *entity.Videos) (err error) {

	err = inits.Gorm.Delete(data).Error
	return
}