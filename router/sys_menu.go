package router

import (
	v1 "54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/api/v1"
	"github.com/gin-gonic/gin"
)

func InitMenuRooter(r *gin.RouterGroup) {
	UserRouter := r.Group("menu")
	{
		UserRouter.GET("getMenu", v1.GetMenu)
		UserRouter.GET("getMenuByPId", v1.GetMenuByPId)
	}
}
