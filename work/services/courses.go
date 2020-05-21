package services

import (
	"fmt"
	"github.com/shijting/go-edu/work/common"
	"github.com/shijting/go-edu/work/dao"
	"github.com/shijting/go-edu/work/entity"
	"github.com/shijting/go-edu/work/httpStatus"
	"github.com/shijting/go-edu/work/serializer"
	"time"
)

type CreateCoursesService struct {
	Title string `form:"title" binding:"required,lt=60" json:"title"`
	Price float32 `form:"price" binding:"-" json:"price"`
	Description string `form:"description" binding:"-" json:"description"`
	SeoKeywords string `form:"seo_keywords" binding:"-" json:"seo_keywords"`
	SeoDescription string `form:"seo_description" binding:"-" json:"seo_description"`
	Thumb string `form:"thumb" binding:"-" json:"thumb"`
	Status uint64 `form:"status" binding:"-" json:"status"`
	PublishedAt string `form:"published_at" binding:"-" json:"published_at"`
	CategoryId uint64 `form:"category_id" binding:"required" json:"category_id"`
	IsRec uint64 `form:"is_rec" binding:"-" json:"is_rec"`
}

func (f *CreateCoursesService)Create() (resp *serializer.Response)  {
	publishedAt, err := time.ParseInLocation("2006-01-02 15:04", f.PublishedAt, time.Local)
	if err != nil {
		resp = &serializer.Response{
			Code:  httpStatus.WRONG_REPEAT_FORM,
			Data:  nil,
			Msg:   "",
			Error: nil,
		}
		return
	}
	data := &entity.Courses{
		Title:          f.Title,
		Price:          f.Price,
		Description:    f.Description,
		SeoKeywords:    f.SeoKeywords,
		SeoDescription: f.SeoDescription,
		Thumb:          f.Thumb,
		ThumbStoreType: common.QINIU_STORE,
		Status:         f.Status,
		PublishedAt:    publishedAt,
		CategoryId:     f.CategoryId,
		IsRec:          f.IsRec,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}
	_, err = dao.CoursesObj.Create(data)
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
type IndexCoursesService struct {
	Page uint64 `form:"page" binding:"required"`
	PageSize uint64 `form:"pageSize" binding:"required"`
}

func (f *IndexCoursesService)Index() (resp *serializer.Response) {
	data,count, err := dao.CoursesObj.GetByPaginate(f.Page, f.PageSize)
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


type EditCourseService struct {
	Id uint64 `form:"id" binding:"required"`
}
func (f *EditCourseService)Edit() (resp *serializer.Response)  {
	data, err := dao.CoursesObj.GetOneById(f.Id)
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
type UpdateCourseService struct {
	Id uint64 `form:"id" binding:"required" json:"id"`
	Title string `form:"title" binding:"required,lt=60" json:"title"`
	Price float32 `form:"price" binding:"-" json:"price"`
	Description string `form:"description" binding:"-" json:"description"`
	SeoKeywords string `form:"seo_keywords" binding:"-" json:"seo_keywords"`
	SeoDescription string `form:"seo_description" binding:"-" json:"seo_description"`
	Thumb string `form:"thumb" binding:"-" json:"thumb"`
	Status uint64 `form:"status" binding:"-" json:"status"`
	PublishedAt string`form:"published_at" binding:"-" json:"published_at"`
	CategoryId uint64 `form:"category_id" binding:"required" json:"category_id"`
	IsRec uint64 `form:"is_rec" binding:"-" json:"is_rec"`
}

func (f *UpdateCourseService)Update()(resp *serializer.Response)  {
	publishedAt, _ := time.ParseInLocation("2006-01-02 15:04", f.PublishedAt, time.Local)
	//publishedAt := f.PublishedAt[:19]
	data := map[string]interface{}{
		"title": f.Title,
		"status": f.Status,
		"price": f.Price,
		"description": f.Description,
		"seo_description": f.SeoDescription,
		"seo_keywords": f.SeoKeywords,
		"thumb": f.Thumb,
		"published_at": publishedAt,
		"category_id": f.CategoryId,
		"is_rec": f.IsRec,
		"updated_at": time.Now(),
	}
	if err := dao.CoursesObj.Update(f.Id, data);err != nil {
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

type  GetCourseService struct {
	Status uint64
}
func (f *GetCourseService)GetAllCourses() (resp *serializer.Response) {
	couses,err := dao.CoursesObj.GetAllByStatus(f.Status)
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
		Data:  couses,
		Msg:   httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
		Error: nil,
	}
	return
}
