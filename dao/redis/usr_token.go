package redis

import (
	"context"
	"strconv"

	g "54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/global"
)

func SetUserToken(userID int64, token string) (err error) {
	ctx := context.Background()
	err = g.RDB.Set(ctx, strconv.FormatInt(userID, 10), token, 0).Err()
	return
}

func GetUserToken(userID int64) (token string, err error) {
	ctx := context.Background()
	token, err = g.RDB.Get(ctx, strconv.FormatInt(userID, 10)).Result()
	return
}
