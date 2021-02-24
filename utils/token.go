package utils

import (
	"strings"

	g "54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/global"
	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/model/request"
	"go.uber.org/zap"

	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/model/code"
	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/model/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetaToken() func(c *gin.Context) string {
	return func(c *gin.Context) string {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.FailWithMessage(code.CodeInvalidToken+"请求头缺少Authorization", c)
			c.Abort()
			return ""
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.FailWithMessage(code.CodeInvalidToken+"请求头中auth格式有误", c)
			c.Abort()
			return ""
		}
		return parts[1]
	}
}

func GetrToken() func(c *gin.Context) string {
	return func(c *gin.Context) string {
		p := new(request.ParamToken)
		if err := c.ShouldBindJSON(p); err != nil {
			zap.L().Error("RefreshToken with invalid param", zap.Error(err))
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				response.FailWithMessage(code.CodeInvalidParam, c)
				c.Abort()
				return ""
			}
			response.FailWithMessage(errs.Translate(g.Trans), c)
			c.Abort()
			return ""
		}
		return p.Rtoken
	}
}
