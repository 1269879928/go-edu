package conf

type Mysql struct {
	Host string `json:"host"`
	Port int64 `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}
type Redis struct {
	Host string `json:"host"`
	Port int64 `json:"port"`
	Auth string `json:"auth"`
}
type Jwt struct {
	Key string `json:"key"`
	Expires int64 `json:"expires"`
}
type Oss struct {
	AccessKeyId string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	RegionId string `json:"region_id"`
}
type Qiniu struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Bucket string `json:"bucket"`
}
// easyjson:json
type Config struct {
	Mysql Mysql `json:"mysql"`
	Redis Redis `json:"redis"`
	Jwt Jwt `json:"jwt"`
	Oss Oss `json:"oss"`
	Qiniu Qiniu `json:"qiniu"`
}
