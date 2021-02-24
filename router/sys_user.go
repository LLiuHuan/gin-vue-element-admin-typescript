package router

import (
	v1 "54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/api/v1"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.RouterGroup) {
	UserRouter := r.Group("user")
	{
		UserRouter.POST("refreshToken", v1.RefreshToken)
	}
}
