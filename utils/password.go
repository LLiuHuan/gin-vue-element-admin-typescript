package utils

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func SetPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		zap.L().Error("")
	}
	return string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
}

func CheckPassword(encodePWD, loginPWD string) error {
	return bcrypt.CompareHashAndPassword([]byte(encodePWD), []byte(loginPWD)) //验证（对比）
}
