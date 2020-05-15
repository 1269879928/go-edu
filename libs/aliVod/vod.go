package aliVod

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go-edu/work/base/inits"
	"os"
)
type UploadAuthDTO struct {
	AccessKeyId string
	AccessKeySecret string
	SecurityToken string
}
type UploadAddressDTO struct {
	Endpoint string
	Bucket string
	FileName string
}
//  使用AK初始化VOD客户端
//var (
//	accessKeyId string = inits.Config.Oss.AccessKeyId
//	accessKeySecret string = inits.Config.Oss.AccessKeySecret
//	regionId string =inits.Config.Oss.RegionId
//)

type Vod struct {
	Client *vod.Client
}


func InitVodClient() (client *vod.Client, err error) {

	// 创建授权对象
	credential := &credentials.AccessKeyCredential{
		AccessKeyId:     inits.Config.Oss.AccessKeyId,
		AccessKeySecret: inits.Config.Oss.AccessKeySecret,
	}
	// 自定义config
	config := sdk.NewConfig()
	config.AutoRetry = true      // 失败是否自动重试
	config.MaxRetryTime = 3      // 最大重试次数
	config.Timeout = 3000000000  // 连接超时，单位：纳秒；默认为3秒
	// 创建vodClient实例
	return vod.NewClientWithOptions(inits.Config.Oss.RegionId, config, credential)
}
// 获取视频上传地址和凭证
type CreateUploadVideo struct {
	Client *vod.Client
	Title string
	Description string
	CoverURL string
	Tags string

}
func (f *CreateUploadVideo)MyCreateUploadVideo() (response *vod.CreateUploadVideoResponse, err error) {
	request := vod.CreateCreateUploadVideoRequest()
	request.Title = f.Title
	request.Description = f.Description
	request.FileName = "/opt/video/sample/video_file.mp4"
	//request.CateId = "-1"
	request.CoverURL = f.CoverURL
	request.Tags = f.Tags // "tag1,tag2"
	request.AcceptFormat = "JSON"
	return f.Client.CreateUploadVideo(request)
}
//3 使用上传凭证和地址初始化OSS客户端（注意需要先Base64解码并Json Decode再传入）
func InitOssClient(uploadAuthDTO UploadAuthDTO, uploadAddressDTO UploadAddressDTO) (*oss.Client, error) {
	client, err := oss.New(uploadAddressDTO.Endpoint,
		uploadAuthDTO.AccessKeyId,
		uploadAuthDTO.AccessKeySecret,
		oss.SecurityToken(uploadAuthDTO.SecurityToken),
		oss.Timeout(86400*7, 86400*7))
	return client, err
}
// 4. 上传本地文件
func UploadLocalFile(client *oss.Client, uploadAddressDTO UploadAddressDTO, localFile string) {
	// 获取存储空间。
	bucket, err := client.Bucket(uploadAddressDTO.Bucket)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 上传本地文件。
	err = bucket.PutObjectFromFile(uploadAddressDTO.FileName, localFile)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}
// 5 刷新上传凭证
func MyRefreshUploadVideo(client *vod.Client, videoId string) (response *vod.RefreshUploadVideoResponse, err error) {
	request := vod.CreateRefreshUploadVideoRequest()
	request.VideoId = videoId
	request.AcceptFormat = "JSON"
	return client.RefreshUploadVideo(request)
}
// 根据视频id 获取视频信息(播放地址)
func MyGetPlayInfo(client *vod.Client, videoId string) (response *vod.GetPlayInfoResponse, err error) {
	request := vod.CreateGetPlayInfoRequest()
	request.VideoId = videoId
	request.AcceptFormat = "JSON"
	return client.GetPlayInfo(request)
}
// 获取播放凭证
func MyGetPlayAuth(client *vod.Client, videoId string) (response *vod.GetVideoPlayAuthResponse, err error) {
	request := vod.CreateGetVideoPlayAuthRequest()
	request.VideoId = videoId
	request.AcceptFormat = "JSON"
	return client.GetVideoPlayAuth(request)
}
