package initialize

import (
	"net/http"

	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/router"

	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/model"

	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/middlewares"
	"github.com/gin-gonic/gin"

	_ "54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/docs" // 导入生成的docs
)

func Router(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(GinLogger(), GinRecovery(true))
	// 跨域
	r.Use(middlewares.Cors())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, model.SettingsConf.Version)
		})
		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
		// swagger 文档
		v1.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	}

	v1.Use()
	{
		router.InitBaseRouter(v1)
	}

	v1.Use(middlewares.JWTAuthMiddleware(), middlewares.RepeatLogin())
	{
		router.InitUserRouter(v1)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
