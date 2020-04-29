package entity

import "time"

// 课程分类
type CourseCategories struct {
	Id uint64 `gorm:"primary_key" json:"id"`
	Name string `gorm:"type:varchar(50);not null;default:'';comment:'课程分类名称'" json:"name"`
	ParentId uint64 `gorm:"type:int;not null;default:0;comment:'上级分类id，0为顶级'" json:"parent_id"`
	Sort uint64 `gorm:"type:int;not null;default:0;comment:'排序: 倒序'" json:"sort"`
	Status uint64 `gorm:"type:tinyint(1);not null;default:1;comment:'状态,1显示，0隐藏'" json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
