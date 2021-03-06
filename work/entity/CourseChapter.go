package entity

import "time"

// 课程章节
type CourseChapter struct {
	Id uint64 `gorm:"primary_key" json:"id"`
	CourseId uint64 `gorm:"type:int;not null;comment:'课程id'" json:"course_id"`
	Title string `gorm:"type:varchar(200);not null;comment:'章节标题'" json:"title"`
	Sort uint64 `gorm:"type:smallint;not null;comment:'排序：升序'" json:"sort"`
	Videos []Videos `gorm:"foreignkey:ChapterId" json:"videos"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
