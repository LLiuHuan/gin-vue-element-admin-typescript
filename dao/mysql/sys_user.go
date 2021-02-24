package mysql

import (
	"database/sql"

	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/model/code"

	g "54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/global"
	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/model"
	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/utils"
)

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) error {
	sqlStr := "select count(user_id) from t_user where username = ?"
	var count int
	if err := g.MDB.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return code.ErrorUserExist
	}
	return nil
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *model.User) (err error) {
	// 执行SQL入库
	sqlStr := `insert into t_user(user_id, username, password) values (?,?,?)`
	_, err = g.MDB.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func CheckUserPwd(user *model.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id, username, password from t_user where username = ?`
	err = g.MDB.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return code.ErrorUserNotExist
	}
	if err != nil {
		return err
	}

	err = utils.CheckPassword(user.Password, oPassword)
	return
}
