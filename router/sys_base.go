package router

import (
	v1 "54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/api/v1"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(r *gin.RouterGroup) {
	UserRouter := r.Group("base")
	{
		UserRouter.POST("signup", v1.SignUp)
		UserRouter.POST("login", v1.Login)
	}
}
