package serializer

// 响应序列化
type Response struct {
	Code int         `json:"code"` // 自定义状态码
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
	Error  interface{} `json:"error"`
	Token string `json:"token"`
}
