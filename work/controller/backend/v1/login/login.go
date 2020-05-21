package login

import (
	"encoding/json"
	"fmt"
	"github.com/GeeTeam/gt3-golang-sdk/geetest"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/shijting/go-edu/work/common"
	"github.com/shijting/go-edu/work/httpStatus"
	"github.com/shijting/go-edu/work/serializer"
	"github.com/shijting/go-edu/work/services"
	"net/http"
	"time"
)

// @Summary 后台登录接口
// @Tags 登录
// @version 1.0
// @Accept json
// @Param email path string true "email"
// @Param password path string true "password"
// @Success 200 {object} serializer.Response 成功后返回值
// @Failure 500 {object} serializer.Response 失败返回值
// @Router /administrator/login [post]
func Login(c *gin.Context)  {
	clientIP := c.ClientIP()
	login := &services.LoginForm{
		RemoteAddr: clientIP,
	}
	if err := c.ShouldBind(login); err != nil {
		c.JSON(http.StatusOK, common.ValidateResponse(err))
		return
	}
	// 校验验证码
	validate, err := common.VerifyJWT(login.Token)
	if err != nil || validate.Ip != clientIP {
		c.JSON(http.StatusOK, &serializer.Response{
			Code: httpStatus.VERIFYCODE_ERROR,
			Msg:  httpStatus.GetCode2Msg(httpStatus.VERIFYCODE_ERROR),
		})
		return
	}
	result := login.Login()
	c.JSON(http.StatusOK, result)
	return
}
//var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var store = sessions.NewCookieStore([]byte("sldfjoiwjkjsd0239849+_()*^&^&%&^$%^$#LDJLSJD..."))
var captchaID = "d5c4bf5f74a1de33ecd30944b12c93a7"
var privateKey = "04e3c2b463a0763972e049cc3441c176"

type RegisterResponse struct {
	Success int `json:"success"`
	Challenge string `json:"challenge"`
	Gt string `json:"gt"`
	NewCaptcha int `json:"new_captcha"`
}
// 初始化Geetest
func RegisterGeetest(c *gin.Context) {
	geetest := geetest.NewGeetestLib(captchaID, privateKey, 2 * time.Second)
	status, response := geetest.PreProcess("", "")
	registerResponse := &RegisterResponse{}
	err := json.Unmarshal(response, registerResponse)
	if err != nil {
		fmt.Printf("Unmarshal err:%#v\n", err)
		return
	}
	session, _ := store.Get(c.Request, "geetest")
	session.Values["geetest_status"] = status
	session.Save(c.Request, c.Writer)
	c.JSON(http.StatusOK, registerResponse)
}
// 验证Geetest
type validateForm struct {
	GeetestChallenge string `form:"geetest_challenge" binding:"required" json:"geetest_challenge"`
	GeetestValidate string `form:"geetest_validate" binding:"required" json:"geetest_validate"`
	GeetestSeccode string `form:"geetest_seccode" binding:"required" json:"geetest_seccode"`
}
func ValidateGeetest(c *gin.Context) {
	form := &validateForm{}
	if err :=c.ShouldBind(form); err !=nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, &serializer.Response{
			Code:  httpStatus.VERIFYCODE_ERROR,
			Msg:  httpStatus.GetCode2Msg(httpStatus.VERIFYCODE_ERROR),
		})
		return
	}
	var geetestRes bool
	geetest := geetest.NewGeetestLib(captchaID, privateKey, 2 * time.Second)
	res := make(map[string]interface{})
	session, _ := store.Get(c.Request, "geetest")
	val := session.Values["geetest_status"]
	status := val.(int8)
	if status == 1 {
		geetestRes = geetest.SuccessValidate(form.GeetestChallenge, form.GeetestValidate, form.GeetestSeccode, "", "")
	} else {
		geetestRes = geetest.FailbackValidate(form.GeetestChallenge, form.GeetestValidate, form.GeetestSeccode)
	}
	if geetestRes {
		validateToken := common.GenJWT(0, "", c.ClientIP(), 300)
		res["code"] = 0
		res["msg"] = "Success"
		res["token"] = validateToken
	} else {
		res["code"] = -100
		res["msg"] = "Failed"
		res["token"] = ""
	}
	c.JSON(http.StatusOK,  &serializer.Response{
		Code:  httpStatus.SUCCESS_STATUS,
		Data: res,
		Msg:  httpStatus.GetCode2Msg(httpStatus.SUCCESS_STATUS),
	})
}