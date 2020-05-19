package services

import (
	"fmt"
	"go-edu/libs/aliVod"
	"go-edu/work/dao"
	"go-edu/work/entity"
	"go-edu/work/httpStatus"
	"go-edu/work/serializer"
	"time"
)

// 阿里云视频点播逻辑处理
// 获取上传凭证
type AliyunVodUploadCreate struct {
	FileName string `form:"file_name" bingding:"required" json:"file_name"`
	//CourseId uint64 `form:"course_id" bingding:"-" json:"course_id"`
}
// 创建上传凭证
func (f *AliyunVodUploadCreate)AliyunAuthTokenCreate() (resp *serializer.Response)  {
	vodClient, err := aliVod.InitVodClient()
	if err != nil {
		resp = &serializer.Response{
			Code:  httpStatus.OPERATION_WRONG,
			Data: nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
			Error: nil,
		}
		return
	}
	//course, err := dao.CoursesObj.GetOneById(f.CourseId)
	if err != nil {
		resp = &serializer.Response{
			Code:  httpStatus.OPERATION_WRONG,
			Data: nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
			Error: nil,
		}
		return
	}
	// http://img03.sogoucdn.com/app/a/100520021/8de3c081b9c92c249460c305a934b1f2
	upload := &aliVod.CreateUploadVideo{
		Client:      vodClient,
		Title:       f.FileName,
		Description: f.FileName,
		CoverURL:    "",
		Tags:        "test",
	}
	response, err := upload.MyCreateUploadVideo()
	if err != nil {
		resp = &serializer.Response{
			Code:  httpStatus.OPERATION_WRONG,
			Data: nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
			Error: nil,
		}
		fmt.Println("UploadVideo failed , err:", err)
		return
	}
	//fmt.Printf("videoid:%#v\n, auth: %#v\n, address:%#v\n", response.VideoId, response.UploadAddress, response.UploadAuth)
	resp = &serializer.Response{
		Code:  httpStatus.SUCCESS_STATUS,
		Data:  map[string]string{"video_id": response.VideoId, "upload_address": response.UploadAddress, "upload_auth": response.UploadAuth},
		Msg:   "ok",
		Error: nil,
	}
	return
}
// 刷新上传凭证
type AliyunVodUploadRefresh struct {
	VideoId string
}
func (f *AliyunVodUploadRefresh)AliyunAuthTokenRefresh() (resp *serializer.Response) {
	vodClient, err := aliVod.InitVodClient()
	if err != nil {
		fmt.Println("init vodclient failed , err:", err)
		resp = &serializer.Response{
			Code:  httpStatus.OPERATION_WRONG,
			Data: nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
			Error: nil,
		}
		return
	}
	response, err := aliVod.MyRefreshUploadVideo(vodClient, f.VideoId)
	if err != nil {
		fmt.Println("UploadVideo failed , err:", err)
		resp = &serializer.Response{
			Code:  httpStatus.OPERATION_WRONG,
			Data: nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
			Error: nil,
		}
		return
	}
	fmt.Printf("videoid:%#v\n, auth: %#v\n, address:%#v\n", response.VideoId, response.UploadAddress, response.UploadAuth)
	resp = &serializer.Response{
		Code:  httpStatus.SUCCESS_STATUS,
		Data:  map[string]string{"video_id": response.VideoId, "upload_address": response.UploadAddress, "upload_auth": response.UploadAuth},
		Msg:   "ok",
		Error: nil,
	}
	return
}

type CreateVideosService struct {
	CourseId uint64 `form:"course_id" binding:"required" json:"course_id"`
	ChapterId uint64 `form:"chapter_id" binding:"required" json:"chapter_id"`
	IsFree uint64 `form:"is_free" json:"is_free"`
	Title string `form:"title" binding:"required" json:"title"`
	Description string `form:"description" json:"description"`
	SeoDescription string  `form:"seo_description" json:"seo_description"`
	SeoKeywords string  `form:"seo_keywords" json:"seo_keywords"`
	Url string `form:"url" json:"url"`
	AliyunVideoId string `form:"aliyun_video_id" json:"aliyun_video_id"`
	Duration uint64 `form:"duration" json:"duration"`
	Status uint64 `form:"status" json:"status"`
}
func (f *CreateVideosService)Create() (resp *serializer.Response) {
	data := &entity.Videos{
		CourseId:       f.CourseId,
		ChapterId:      f.ChapterId,
		IsFree:         f.IsFree,
		Title:          f.Title,
		Description:    f.Description,
		SeoDescription: f.SeoDescription,
		SeoKeywords:    f.SeoKeywords,
		Url:            f.Url,
		AliyunVideoId:  f.AliyunVideoId,
		Duration:       f.Duration,
		Status:         f.Status,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}
	_, err := dao.VideosDaoObj.Create(data)
	if err != nil {
		resp = &serializer.Response{
			Code:  httpStatus.OPERATION_WRONG,
			Data: nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
			Error: nil,
		}
		return
	}
	resp = &serializer.Response{
		Code:  httpStatus.SUCCESS_STATUS,
		Msg:   "ok",
	}
	return
}

// 分页列表
type IndexVideosService struct {
	Page uint64 `form:"page" binding:"required"`
	PageSize uint64 `form:"pageSize" binding:"required"`
	Title string `form:"title"`
	CourseId uint64 `form:"course_id"`
	ChapterId uint64 `form:"chapter_id"`
}

func (f *IndexVideosService)Index() (resp *serializer.Response) {
	data,count, err := dao.VideosDaoObj.GetByPaginate(f.Page, f.PageSize, f.CourseId, f.ChapterId, f.Title)
	if err != nil {
		resp = &serializer.Response{
			Code:  httpStatus.OPERATION_WRONG,
			Data:  nil,
			Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
			Error: nil,
		}
		return
	}
	if len(data) > 0 {
		courseIds := make([]uint64,0)
		chapterIds := make([]uint64,0)
		for _, item := range data {
			courseIds = append(courseIds, item.CourseId)
			chapterIds = append(chapterIds, item.ChapterId)
		}
		chapters, err := dao.CourseChapterObj.GetSomeByIds(chapterIds)
		if err != nil {
			resp = &serializer.Response{
				Code:  httpStatus.OPERATION_WRONG,
				Data:  nil,
				Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
				Error: nil,
			}
			return
		}
		courses, err := dao.CoursesObj.GetSomeByIds(courseIds)
		if err != nil {
			resp = &serializer.Response{
				Code:  httpStatus.OPERATION_WRONG,
				Data:  nil,
				Msg:   httpStatus.GetCode2Msg(httpStatus.OPERATION_WRONG),
				Error: nil,
			}
			return
		}
		for _, video := range data {
			for _, chapter := range chapters {
				if video.ChapterId == chapter.Id {
					video.Chapter = chapter
				}
			}
			for _, course := range courses {
				if video.CourseId == course.Id {
					video.Course = course
				}
			}
		}
	}
	resp = &serializer.Response{
		Code:  httpStatus.SUCCESS_STATUS,
		Data:  map[string]interface{}{"list":data, "total":count},
		Msg:   httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
		Error: nil,
	}
	return
}


type EditVideosService struct {
	Id uint64 `form:"id" binding:"required" json:"id"`
}

func (f *EditVideosService)Edit() (resp *serializer.Response) {
	data , err := dao.VideosDaoObj.GetOneById(f.Id)
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
	}
	return
}
type UpdateVideosService struct {
	Id uint64 `form:"id" binding:"required" json:"id"`
	CourseId uint64 `form:"course_id" binding:"required" json:"course_id"`
	ChapterId uint64 `form:"chapter_id" binding:"required" json:"chapter_id"`
	IsFree uint64 `form:"is_free" binding:"-" json:"is_free"`
	Title string `form:"title" binding:"required" json:"title"`
	Description string `form:"description" binding:"-" json:"description"`
	SeoDescription string  `form:"seo_description" binding:"-" json:"seo_description"`
	SeoKeywords string  `form:"seo_keywords" binding:"-" json:"seo_keywords"`
	Url string `form:"url" binding:"-" json:"url"`
	AliyunVideoId string `form:"aliyun_video_id" binding:"-" json:"aliyun_video_id"`
	Duration uint64 `form:"duration" binding:"-" json:"duration"`
	Status uint64 `form:"status" json:"status"`
}
func (f *UpdateVideosService)Update() (resp *serializer.Response) {
	data := map[string]interface{}{
		"course_id": f.CourseId,
		"chapter_id": f.ChapterId,
		"is_free": f.IsFree,
		"title": f.Title,
		"description": f.Description,
		"seo_description": f.SeoDescription,
		"seo_keywords": f.SeoKeywords,
		"url": f.Url,
		"aliyun_video_id": f.AliyunVideoId,
		"duration": f.Duration,
		"status": f.Status,
	}
	err := dao.VideosDaoObj.Update(f.Id, data)
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
	}
	return
}
type DeleteVideosService struct {
	Id uint64 `form:"id" binding:"required" json:"id"`
}
func (f *DeleteVideosService)Delete() (resp *serializer.Response) {
	data := &entity.Videos{Id: f.Id}
	err := dao.VideosDaoObj.Delete(data)
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
	}
	return
}