package jwt

import (
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"

	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/model"

	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/model/code"

	"github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

// GenToken ⽣成access token 和 refresh token
func GenToken(userID int64, isRToken bool) (aToken, rToken string, err error) {
	var (
		issuer = model.SettingsConf.UserConfig.Issuer
		secret = []byte(model.SettingsConf.UserConfig.Secret)
	)
	fmt.Println(issuer, secret)
	// 创建⼀个我们⾃⼰的声明
	c := MyClaims{
		userID, // ⾃定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(viper.GetInt("auth.jwt_expire"))).Unix(), // 过期时间
			Issuer:    issuer,                                                                            // 签发⼈
		},
	}
	// 加密并获得完整的编码后的字符串token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256,
		c).SignedString(secret)
	// refresh token 不需要存任何⾃定义数据
	if isRToken {
		rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(), // 过期时间
			Issuer:    issuer,                                     // 签发⼈
		}).SignedString(secret)
	}
	fmt.Println(aToken, rToken, isRToken)
	// 使⽤指定的secret签名并获得完整的编码后的字符串token
	return
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (claims *MyClaims, err error) {
	// 解析token
	var token *jwt.Token
	var secret = []byte(model.SettingsConf.UserConfig.Secret)
	claims = new(MyClaims)
	token, err = jwt.ParseWithClaims(tokenString, claims, func(*jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return
	}
	if !token.Valid { // 校验token
		err = errors.New("invalid token")
	}
	return
}

// RefreshToken 刷新AccessToken
func RefreshToken(aToken, rToken string, oUserID int64) (newAToken, newRToken string, err error) {
	newAToken = aToken
	newRToken = rToken
	// aToken 失效的话就继续 没失效就返回
	var secret = []byte(model.SettingsConf.UserConfig.Secret)
	if _, err = jwt.Parse(aToken, func(*jwt.Token) (interface{}, error) {
		return secret, nil
	}); err == nil {
		zap.L().Info("aToken 未失效")
		return
	}

	// refresh token⽆效直接返回 没失效就继续
	if _, err = jwt.Parse(rToken, func(*jwt.Token) (interface{}, error) {
		return secret, nil
	}); err != nil {
		zap.L().Info("rToken 已失效")
		return
	}

	// 从旧access token中解析出claims数据
	var claims MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, func(*jwt.Token) (interface{}, error) {
		return secret, nil
	})
	v, _ := err.(*jwt.ValidationError)
	if claims.UserID == oUserID {
		// 当access token是过期错误 并且 refresh token没有过期时就创建⼀个新的access token
		if v.Errors == jwt.ValidationErrorExpired {
			if rToken == "" {
				newAToken, newRToken, err = GenToken(claims.UserID, true)
				return
			} else {
				newAToken, _, err = GenToken(claims.UserID, false)
				return
			}
		}
	} else {
		err = code.ErrorInvalidToken
	}
	return
}

func GetTokenUserId(aToken string) (userId int64, err error) {
	var secret = []byte(model.SettingsConf.UserConfig.Secret)
	var claims MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, func(*jwt.Token) (interface{}, error) {
		return secret, nil
	})
	userId = claims.UserID
	return
}
