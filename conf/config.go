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
	Post int64 `json:"post"`
	Auth string `json:"auth"`
}
type Jwt struct {
	Key string `json:"key"`
}

type Config struct {
	Mysql Mysql `json:"mysql"`
	Redis Redis `json:"redis"`
	Jwt Jwt `json:"jwt"`
}
