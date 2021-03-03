package initialize

import (
	"fmt"
	"os"

	"gorm.io/gorm/logger"

	g "54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/global"

	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Gorm 初始化Gorm v2
// 初始化数据库并返回数据库全局变量
func Gorm(DbType string, cfg *model.MySqlConfig) (err error) {
	switch DbType {
	case "mysql":
		err = GormMySql(cfg)
	default:
		err = GormMySql(cfg)
	}
	return
}

// GormMySql 初始化mysql数据库
func GormMySql(cfg *model.MySqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.PassWrod,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	g.GROM, err = gorm.Open(mysql.New(mysqlConfig), gormConfig())
	if err != nil {
		zap.L().Error("gorm启动异常", zap.Error(err))
		os.Exit(0)
		return err
	}
	sqlDB, err := g.GROM.DB()
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	return
}

// 获取配置信息 后期可能会扩展配置信息 所以单独抽离出来一个函数
func gormConfig() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	config.Logger = logger.Default.LogMode(logger.Warn)
	return config
}
