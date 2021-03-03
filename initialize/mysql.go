package initialize

import (
	"fmt"

	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/model"

	"go.uber.org/zap"

	g "54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/global"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// MySql 初始化MySql连接 使用sqlx
func MySql(cfg *model.MySqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.PassWrod,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
	// 也可以使用MustConnect连接不成功就panic
	g.MDB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	g.MDB.SetMaxOpenConns(cfg.MaxOpenConns)
	g.MDB.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}

// MySqlClose 关闭mysql连接  废弃 改为公共的
func MySqlClose() {
	fmt.Println("关闭mysql")
	_ = g.MDB.Close()
}
