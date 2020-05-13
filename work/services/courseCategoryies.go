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
}
func (f *CreateCourseCategoriesService)Create()(resp *serializer.Response)  {
	data := &entity.CourseCategories{
		Name:      f.Name,
		ParentId:  0,
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
type IndexCourseCategoriesService struct {
	Page uint64 `form:"page" binding:"required"`
	PageSize uint64 `form:"pageSize" binding:"required"`
}

func (f *IndexCourseCategoriesService)Index() (resp *serializer.Response) {
	data,count, err := dao.CreateCourseCategoriesObj.GetByPaginate(f.Page, f.PageSize)
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
		Data:  map[string]interface{}{"list":data, "total":count},
		Msg:   httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
		Error: nil,
	}
	return
}
type EditCourseCategoriesService struct {
	Id uint64 `form:"id" binding:"required"`
}

func (f *EditCourseCategoriesService)Edit() (resp *serializer.Response)  {
	data, err := dao.CreateCourseCategoriesObj.GetOneById(f.Id)
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
		Data:  data,
		Msg:   httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
		Error: nil,
	}
	return
}
type UpdateCourseCategoriesService struct {
	Id uint64 `form:"id" binding:"required" json:"id"`
	Name string `form:"name" binding:"required,gt=0,lt=20" json:"name"`
	Status uint64 `form:"status" json:"status"`
	Sort uint64 `form:"sort" json:"sort"`
}

func (f *UpdateCourseCategoriesService)Update()(resp *serializer.Response)  {
	data := map[string]interface{}{
		"name": f.Name,
		"status": f.Status,
		"sort": f.Sort,
	}
	if err := dao.CreateCourseCategoriesObj.Update(f.Id, data);err != nil {
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
		Data:  data,
		Msg:   httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
		Error: nil,
	}
	return
}
type DeleteCourseCategoriesService struct {
	Id uint64 `form:"id" binding:"required" json:"id"`
	Status uint64 `form:"status"`
}

func (f *DeleteCourseCategoriesService)Delete() (resp *serializer.Response) {
	data := map[string]interface{} {"status": f.Status}
	if err := dao.CreateCourseCategoriesObj.Update(f.Id, data);err != nil {
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
		Data:  data,
		Msg:   httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
		Error: nil,
	}
	return
}

type GetAllCourseCategoriesService struct{}
func (f *GetAllCourseCategoriesService)GetAllCategories() (resp *serializer.Response) {
	data, err := dao.CreateCourseCategoriesObj.GetAll()
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
		Data:  data,
		Msg:   httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
		Error: nil,
	}
	return
}