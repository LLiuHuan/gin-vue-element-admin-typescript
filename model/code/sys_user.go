package code

import "errors"

var (
	CodeUserExist       = "用户名已存在"
	CodeUserNotExist    = "用户名不存在"
	CodeInvalidPassword = "用户名或密码错误"
	CodeLoginSuccess    = "登录成功"
	CodeRefreshTokenOK  = "Token 刷新成功"

	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)
