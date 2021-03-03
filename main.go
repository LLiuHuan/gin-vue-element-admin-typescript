package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/initialize"
	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/model"

	"go.uber.org/zap"
)

func init() {
	// 0. 获取配置文件路径
	filePath := flag.String("f", "./config/config.yaml", "配置文件路径")
	fmt.Printf("配置的配置文件路径为：%s \v", *filePath)
	// 1. 加载配置
	if err := initialize.Settings(*filePath); err != nil {
		fmt.Printf("init Settings failed, err: %v\n", err)
		os.Exit(0)
		return
	}
	// 2. 加载日志
	if err := initialize.Logger(model.SettingsConf.LogConfig, model.SettingsConf.Mode); err != nil {
		fmt.Printf("init Logger failed, err: %v\n", err)
		zap.L().Error("init Logger failed, err", zap.Error(err))
		os.Exit(0)
		return
	}
	// 3. 初始化MySQL连接
	// 初始化sqlx
	if err := initialize.MySql(model.SettingsConf.MySqlConfig); err != nil {
		fmt.Printf("init MySql failed, err: %v\n", err)
		zap.L().Error("init MySql failed, err", zap.Error(err))
		os.Exit(0)
		return
	}
	// 初始化gorm
	if err := initialize.GormMySql(model.SettingsConf.MySqlConfig); err != nil {
		fmt.Printf("init GormMySql failed, err: %v\n", err)
		zap.L().Error("init GormMySql failed, err", zap.Error(err))
		os.Exit(0)
		return
	}
	// 4. 初始化Redis连接
	if err := initialize.Redis(model.SettingsConf.RedisConfig); err != nil {
		fmt.Printf("init Redis failed, err: %v\n", err)
		zap.L().Error("init Redis failed, err", zap.Error(err))
		os.Exit(0)
		return
	}
	// 5. 初始化雪花算法
	if err := initialize.Snowflake(model.SettingsConf.StartTime, model.SettingsConf.MachineID); err != nil {
		fmt.Printf("init Snowflake failed, err: %v\n", err)
		zap.L().Error("init Snowflake failed, err", zap.Error(err))
		os.Exit(0)
		return
	}
	// 6. 初始化gin框架内置的校验器使用的翻译器
	if err := initialize.Trans("zh"); err != nil {
		fmt.Printf("init validator failed, err: %v\n", err)
		zap.L().Error("init validator failed, err", zap.Error(err))
		os.Exit(0)
		return
	}
}

// Go Web 开发通用脚手架

// @title Swagger Example API
// @version 0.0.1
// @description 描述信息
// @termsOfService http://swagger.io/terms/

// @contact.name 刘欢
// @contact.url 54cc.cc
// @contact.email liuhuan@54cc.cc

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 54cc.cc
// @BasePath /
func main() {
	zap.L().Info("Server starting")
	// 延迟关闭init
	defer zap.L().Sync()
	defer initialize.MySqlClose()
	defer initialize.RedisClose()
	// 1. 注册路由
	r := initialize.Router(model.SettingsConf.Mode)
	// 2. 启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", model.SettingsConf.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
