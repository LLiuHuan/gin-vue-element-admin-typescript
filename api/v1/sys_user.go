package v1

import (
	"errors"

	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/model"

	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/utils"

	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/model/code"

	g "54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/global"
	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/model/request"
	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/server"

	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/model/response"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// SignUp 注册接口
// @Summary 用户注册接口
// @Description 根据参数，注册用户
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param who query request.ParamSignUp true "用户名, 密码, 确认密码"
// @Security ApiKeyAuth
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /api/v1/user/signup [post]
func SignUp(c *gin.Context) {
	// 1. 获取参数 参数校验
	p := new(request.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// zap.String("xx", "vv"),
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.FailWithMessage(code.CodeInvalidParam, c)
		}
		response.FailWithMessage(errs.Translate(g.Trans), c)
		return
	}
	// 手动对请求参数进行详细的业务规则校验
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.RePassword != p.Password {
	//
	//}
	//fmt.Println(p)
	// 2. 业务处理
	if err := server.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, code.ErrorUserNotExist) {
			response.FailWithMessage(code.CodeUserExist, c)
		}
		response.FailWithMessage(code.CodeDataBaseError, c)
		return
	}
	// 3. 返回响应
	response.OkWithMessage("注册成功", c)
}

// Login 登录请求
func Login(c *gin.Context) {
	// 1. 获取请求参数及参数校验
	p := new(request.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.FailWithMessage(code.CodeInvalidParam, c)
			return
		}
		response.FailWithMessage(errs.Translate(g.Trans), c)
		return
	}
	// 2. 业务逻辑处理
	aToken, rToken, err := server.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, code.ErrorUserNotExist) {
			response.FailWithMessage(code.CodeUserNotExist, c)
			return
		}
		response.FailWithMessage(code.CodeInvalidPassword, c)
		return
	}
	// 3. 返回响应
	data := map[string]string{
		"aToken": aToken,
		"rToken": rToken,
	}
	response.OkWithDetailed(data, code.CodeLoginSuccess, c)
}

func RefreshToken(c *gin.Context) {
	// 1. 获取aToken和rToken
	aToken := utils.GetaToken()(c)
	rToken := utils.GetrToken()(c)
	if aToken == "" && rToken == "" {
		response.FailWithMessage(code.CodeInvalidToken, c)
	}
	userID, ok := c.Get(model.CtxUserIdKey)
	if !ok {
		response.FailWithMessage(code.CodeInvalidToken, c)
	}
	nAToken, nRToken, err := server.RefreshToken(aToken, rToken, userID.(int64))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	data := map[string]string{
		"aToken": nAToken,
		"rToken": nRToken,
	}
	response.OkWithDetailed(data, code.CodeRefreshTokenOK, c)
}
