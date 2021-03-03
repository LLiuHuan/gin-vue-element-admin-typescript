package router

import (
	v1 "54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/api/v1"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.RouterGroup) {
	UserRouter := r.Group("user")
	{
		// 重置token
		UserRouter.POST("refreshToken", v1.RefreshToken)
		// 用户基础信息
		UserRouter.GET("userInfo", v1.UserInfo)
	}
}
