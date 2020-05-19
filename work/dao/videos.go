package dao

import (
	"go-edu/work/base/inits"
	"go-edu/work/entity"
)

type VideosDao struct {}
var VideosDaoObj *VideosDao

var VideosColumn = "id, title, course_id, chapter_id, is_free,status, description, seo_description, seo_keywords, url, aliyun_video_id, duration,created_at"
func (*VideosDao)Create(data *entity.Videos)(result *entity.Videos, err error)  {
	err = inits.Gorm.Create(&data).Error
	result = data
	return
}
func (*VideosDao)GetByPaginate(page,pageSize, courseId, chapterId uint64, title string)(res []*entity.Videos, total int, err error)  {
	tx := inits.Gorm.Model(&entity.Videos{})
	if courseId > 0 {
		tx = tx.Where("course_id = ?", courseId)
	}
	if chapterId > 0 {
		tx = tx.Where("chapter_id = ?", chapterId)
	}
	if len(title) > 0 {
		tx = tx.Where("title LIKE ?", "%" + title + "%")
	}
	err = tx.Select(VideosColumn).Offset((page - 1) * pageSize).Limit(pageSize).Find(&res).Error
	tx.Count(&total)
	return
}
func (*VideosDao)GetOneById(id uint64)(result entity.Videos, err error) {
	err = inits.Gorm.Select(VideosColumn).Where("id = ?", id).First(&result).Error
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