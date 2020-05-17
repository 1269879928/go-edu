package services

import (
	"fmt"
	"go-edu/work/dao"
	"go-edu/work/entity"
	"go-edu/work/httpStatus"
	"go-edu/work/serializer"
	"time"
)

type CreateCourseChapterService struct {
	Title string `form:"title" binding:"required,lt=60" json:"title"`
	Sort uint64 `form:"sort" binding:"required" json:"sort"`
	CourseId uint64 `form:"course_id" binding:"required" json:"course_id"`
}

func (f *CreateCourseChapterService)Create() (resp *serializer.Response)  {
	data := &entity.CourseChapter{
		Title:          f.Title,
		Sort: f.Sort,
		CourseId: f.CourseId,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}
	_, err := dao.CourseChapterObj.Create(data)
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
		Data:  nil,
		Msg:   httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
		Error: nil,
	}
	return
}
type IndexCourseChapterService struct {
	CourseId uint64 `form:"course_id" binding:"required"`
	Page uint64 `form:"page" binding:"required"`
	PageSize uint64 `form:"pageSize" binding:"required"`
}

func (f *IndexCourseChapterService)Index() (resp *serializer.Response) {
	data,count, err := dao.CourseChapterObj.GetByPaginate(f.CourseId, f.Page, f.PageSize)
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


type EditCourseChapterService struct {
	Id uint64 `form:"id" binding:"required"`
}
func (f *EditCourseChapterService)Edit() (resp *serializer.Response)  {
	data, err := dao.CourseChapterObj.GetOneById(f.Id)
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
type UpdateCourseChapterService struct {
	Id uint64 `form:"id" binding:"required" json:"id"`
	Title string `form:"title" binding:"required,lt=60" json:"title"`
	Sort uint64 `form:"sort" binding:"-" json:"sort"`
	CourseId uint64 `form:"course_id" binding:"required" json:"course_id"`
}

func (f *UpdateCourseChapterService)Update()(resp *serializer.Response)  {
	data := map[string]interface{}{
		"id":f.Id,
		"title": f.Title,
		"sort": f.Sort,
		"updated_at": time.Now(),
	}
	if err := dao.CourseChapterObj.Update(f.Id, data);err != nil {
		resp = &serializer.Response{
			Code:  httpStatus.OPERATION_WRONG,
			Data:  nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
			Error: nil,
		}
		fmt.Println(err)
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
type DeleteCourseChapterService struct {
	Id uint64 `form:"id" binding:"required" json:"id"`
}

func (f *DeleteCourseChapterService)Delete()(resp *serializer.Response)  {
	//publishedAt := f.PublishedAt[:19]
	data := &entity.CourseChapter{Id: f.Id}
	if err := dao.CourseChapterObj.Delete(data);err != nil {
		resp = &serializer.Response{
			Code:  httpStatus.OPERATION_WRONG,
			Data:  nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
			Error: nil,
		}
		fmt.Println(err)
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
type CourseChapterCourseService struct {
	CourseId uint64 `form:"course_id" binding:"required" json:"course_id"`
}

func (f *CourseChapterCourseService)Delete()(resp *serializer.Response)  {
	//publishedAt := f.PublishedAt[:19]
	data, err := dao.CourseChapterObj.GetOneByCourseId(f.CourseId);
	if err != nil {
		resp = &serializer.Response{
			Code:  httpStatus.OPERATION_WRONG,
			Data:  nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
			Error: nil,
		}
		fmt.Println(err)
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
