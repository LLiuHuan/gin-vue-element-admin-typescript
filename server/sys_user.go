package server

import (
	"fmt"

	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/dao/mysql"
	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/dao/redis"
	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/model"
	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/model/request"
	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/pkg/jwt"
	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/pkg/snowflake"
	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/utils"
)

// SignUp 注册业务处理
func SignUp(p *request.ParamSignUp) (err error) {
	// 判断用户存不存在
	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		// 数据库查询出错
		return err
	}
	// 生成UID
	userID := snowflake.GenID()
	// 构造user实例
	user := &model.User{
		UserID:   userID,
		Username: p.Username,
		Password: utils.SetPassword(p.Password),
	}

	// 保存进数据库
	return mysql.InsertUser(user)
}

// Login 登录
func Login(p *request.ParamLogin) (aToken, rToken string, err error) {
	user := &model.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递的是指针，就能拿到user.UserID
	if err := mysql.CheckUserPwd(user); err != nil {
		return "", "", err
	}
	// 生成JWT
	aToken, rToken, err = jwt.GenToken(user.UserID, true)
	err = redis.SetUserToken(user.UserID, aToken)
	return
}

func RefreshToken(aToken, rToken string, userID int64) (nAToken, nRToken string, err error) {
	fmt.Println(userID, aToken, rToken)
	nAToken, nRToken, err = jwt.RefreshToken(aToken, rToken, userID)
	fmt.Println(nAToken, nRToken, err)
	if err == nil {
		err = redis.SetUserToken(userID, aToken)
	}
	return
}
