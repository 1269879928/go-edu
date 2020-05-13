package entity

import "time"

// 课程
type Courses struct {
	Id uint64 `gorm:"primary_key" json:"id"`
	Title string `gorm:"type:varchar(200);not null;default:'';comment:'课程标题'" json:"title"`
	Price float32 `gorm:"type:decimal(8,2);not null;default:'0.00';comment:'收费'" json:"price"`
	Description string `gorm:"type:text;comment:'描述'" json:"description"`
	SeoKeywords string `gorm:"type:varchar(255);not null;default:'';comment:'seo关键词(多个以,分开)'" json:"seo_keywords"`
	SeoDescription string `gorm:"type:varchar(255);not null;default:'';comment:'seo描述'" json:"seo_description"`
	Thumb string `gorm:"type:varchar(200);not null;default:'';comment:'封面图片'" json:"thumb"`
	ThumbStoreType uint64 `gorm:"type:tinyint(1);not null;default:1;comment:'封面图片存储类型(1:本地2:七牛云3:阿里云)'" json:"thumb_store_type"`
	Status uint64 `gorm:"type:tinyint(1);not null;default:1;comment:'状态(0隐藏，1显示)'" json:"status"`
	PublishedAt time.Time `gorm:"comment:'发布时间'" json:"published_at"`
	CategoryId uint64 `gorm:"type:int;not null;default:0;comment:'所属分类'" json:"category_id"`
	IsRec uint64 `gorm:"type:tinyint(1);not null;default:0;comment:'推荐(0否1是)'" json:"is_rec"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
