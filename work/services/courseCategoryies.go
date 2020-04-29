package services

import (
	"go-edu/work/dao"
	"go-edu/work/entity"
	"go-edu/work/httpStatus"
	"go-edu/work/serializer"
)

type CreateCourseCategoriesService struct {
	Name string `form:"name" binding:"required" json:"name"`
	Sort uint64 `form:"sort" json:"sort"`
	ParentId uint64 `form:"parent_id" json:"parent_id"`
}
func (f *CreateCourseCategoriesService)Create()(resp *serializer.Response)  {
	data := &entity.CourseCategories{
		Name:      f.Name,
		ParentId:  f.ParentId,
		Sort:      f.Sort,
		Status:    1,
	}
	res, err := dao.CreateCourseCategoriesObj.Create(data)
	if err != nil {
		resp = &serializer.Response{
			Code:  httpStatus.OPERATION_WRONG,
			Data:  nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
			Error: nil,
		}
		return
	}
	resp = &serializer.Response{
		Code:  httpStatus.SUCCESS_STATUS,
		Data:  res,
		Msg:   httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
		Error: nil,
	}
	return
}
