package httpStatus
var codeToMsg map[int]string
func init() {
	codeToMsg = make(map[int]string)
	codeToMsg[SUCCESS_STATUS] = "操作成功"
	codeToMsg[OPERATION_WRONG] = "操作失败"
	codeToMsg[PARAM_WRONG] = "参数有误"
	codeToMsg[WRONG_PASSWORD] = "旧密码不正确"
	codeToMsg[WRONG_EMAIL_PASSWORD] = "登录账户或者密码不正确"
	codeToMsg[PASSWD_TOO_SHORT] = "新密码不能小于6位或者超过20位"
	codeToMsg[WRONG_PASSWD_SAMPLE] = "新密码不能与旧密码相同"
	codeToMsg[WRONG_PASSWD_CONFIRM] = "新密码确认失败"
	codeToMsg[VERIFYCODE_EXPIRE] = "图片验证码过期"
	codeToMsg[VERIFYCODE_ERROR] = "图片验证码错误"
	codeToMsg[WRONG_REPEAT_FORM] = "操作不合法"
	codeToMsg[WRONG_NAME_EXIST] = "名称已存在"
	codeToMsg[WRONG_PATH_EXIST] = "路径已存在"
	codeToMsg[WRONG_NAME_ILLEGAL_CHAR_CHINESE] = "名称只能包含中英文"
	codeToMsg[WRONG_SUPER_ROLE_OPERATION] = "无权限操作此角色"
	codeToMsg[WRONG_NAME_ILLEGAL_CHAR] = "名称只能包含英文数字且开头不能为数字"
	codeToMsg[WRONG_PERMISSION_EMPTY] = "此角色权限分配不能为空"
	codeToMsg[WRONG_PERMISSION_NO_SELECT] = "请选择用户角色"
	codeToMsg[WRONG_PERMISSION_PATH_EMPTY] = "非顶级父菜单路径不能为空"
	codeToMsg[WRONG_PERMISSION_NO_HAVE] = "无权限访问"
	codeToMsg[WRONG_LOGRECORD_NO_CHECK] = "请勾选需要删除的选项"
	codeToMsg[FILE_UPLOAD_FAIL] = "文件上传失败"
	codeToMsg[FILE_UPLOAD_KEEP] = "文件上传中"
	codeToMsg[FAIL_MERGE_FILE] = "文件合并失败"
	codeToMsg[FAIL_SUFFIX_FILE] = "文件上传格式不支持"
	codeToMsg[GETTING_DATA_FAIL] = "获取数据失败"
	//codeToChnMsg[GETTING_DATA_FAIL] = "非法操作或无效参数"

}
func GetCode2Msg(code int) string  {
	if msg, ok :=codeToMsg[code]; ok {
		return msg
	}
	return ""
}
