package services

import (
	"fmt"
	"go-edu/libs/aliVod"
	"go-edu/work/base/inits"
	"go-edu/work/dao"
	"go-edu/work/entity"
	"go-edu/work/httpStatus"
	"go-edu/work/serializer"
	"time"
)

// 阿里云视频点播逻辑处理
// 获取上传凭证
type AliyunVodUploadCreate struct {
	CourseId uint64 `form:"course_id" bingding:"required" json:"course_id"`
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
	course, err := dao.CoursesObj.GetOneById(f.CourseId)
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
		Title:       course.Title,
		Description: "",
		CoverURL:    inits.Config.Qiniu.Domain + course.Thumb,
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
	fmt.Printf("videoid:%#v\n, auth: %#v\n, address:%#v\n", response.VideoId, response.UploadAddress, response.UploadAuth)
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
	IsFree uint64 `form:"is_free" binding:"-" json:"is_free"`
	Title string `form:"title" binding:"required" json:"title"`
	Description string `form:"description" binding:"-" json:"description"`
	SeoDescription string  `form:"seo_description" binding:"-" json:"seo_description"`
	SeoKeywords string  `form:"seo_keywords" binding:"-" json:"seo_keywords"`
	Url string `form:"url" binding:"-" json:"url"`
	AliyunVideoId uint64 `form:"aliyun_video_id" binding:"-" json:"aliyun_video_id"`
	Duration uint64 `form:"duration" binding:"-" json:"duration"`
	Status uint64 `form:"status" binding:"status" json:"status"`
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

type EditVideosService struct {
	Id uint64 `form:"id" binding:"required" json:"id"`
}

func (f *EditVideosService)Edit()  {

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
	AliyunVideoId uint64 `form:"aliyun_video_id" binding:"-" json:"aliyun_video_id"`
	Duration uint64 `form:"duration" binding:"-" json:"duration"`
	Status uint64 `form:"status" binding:"status" json:"status"`
}
func (f *UpdateVideosService)Update()  {

}
