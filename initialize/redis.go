package initialize

import (
	"context"
	"fmt"
	"time"

	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/model"

	g "54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/global"
	"github.com/go-redis/redis/v8"
)

// Redis 初始化Redis连接
func Redis(cfg *model.RedisConfig) (err error) {
	g.RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.PassWrod, // no password set
		DB:       cfg.DB,       // use default DB
		PoolSize: cfg.PoolSize, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = g.RDB.Ping(ctx).Result()
	return err
}

// RedisClose 关闭Redis连接 废弃 改为公共的
func RedisClose() {
	_ = g.RDB.Close()
}
