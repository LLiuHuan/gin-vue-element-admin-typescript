package middlewares

import (
	"strings"

	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/model/code"

	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/model"

	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/model/response"

	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/dao/redis"
	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/pkg/jwt"
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种⽅式 1.放在请求头 2.放在请求体 3.放在URI
		// 这⾥假设Token放在Header的中
		// 这⾥的具体实现⽅式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.FailWithMessage("请求头缺少Authorization", c)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.FailWithMessage("请求头中auth格式有误", c)
			c.Abort()
			return
		}
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			response.FailWithMessage("无效的Token", c)
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下⽂c上
		c.Set(model.CtxUserIdKey, mc.UserID)
		c.Set(model.CtxTokenKey, parts[1])
		c.Next() // 后续的处理函数可以⽤过c.Get("userID")来获取当前请求的⽤户信息
	}
}

func RepeatLogin() func(c *gin.Context) {
	return func(c *gin.Context) {
		userID, ok := c.Get(model.CtxUserIdKey)
		if !ok {
			return
		}
		token, ok := c.Get(model.CtxTokenKey)
		if !ok {
			return
		}
		oToken, err := redis.GetUserToken(userID.(int64))
		if err != nil {
			c.Abort()
			return
		}
		if token != oToken {
			response.FailWithMessage(code.CodeRepeatLogin, c)
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}
