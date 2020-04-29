package entity

import (
	"time"
)

type Videos struct {
	Id uint64 `gorm:"primary_ke" json:"id"`
	CourseId uint64 `gorm:"type:int;not null;comment:'所属课程id';index:idx_course_id" json:"course_id"`
	ChapterId uint64 `gorm:"type:int;not null;comment:'所属章节id';index:idx_chapter_id" json:"chapter_id"`
	UserId uint64 `gorm:"type:int;not null;default:0;comment:'id';index:idx_user_id" json:"user_id"`
	Title string `gorm:"type:varchar(200);not null;comment'标题'" json:"title"`
	Description string `gorm:"type:text;comment:'描述'" json:"description"`
	SeoDescription string `gorm:"type:varchar(200);not null;default:'';comment:'seo描述'" json:"seo_description"`
	SeoKeywords string `gorm:"type:varchar(255);not null;default:'';comment:'seo关键词'" json:"seo_keywords"`
	Url string `gorm:"type:varchar(255);not null;comment:'播放地址'"`
	AliyunVideoId uint64 `gorm:"type:int;not null;default:0;comment:'阿里云短视频Id'" json:"aliyun_video_id"`
	PublishedAt time.Time `gorm:"comment:'发布时间'" json:"published_at"`
	Duration uint64 `gorm:"type:int;not null;default:0;comment:'时长(秒)'" json:"duration"`
	Status uint64 `gorm:"type:tinyint(1);not null;default:1;comment:'状态(0隐藏，1显示)'" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
 }
