package code

import "errors"

const (
	CodeInvalidParam  = "请求参数错误"
	CodeSuccess       = "success"
	CodeServerBusy    = "服务器繁忙"
	CodeNeedLogin     = "需要登录"
	CodeInvalidToken  = "无效的Token"
	CodeRepeatLogin   = "该账号在其他设备登录"
	CodeInvalidHeader = "请求头缺少Authorization"
	CodeDataBaseError = "数据库内部错误"
)

var (
	ErrorInvalidToken = errors.New("token 被篡改")
)
