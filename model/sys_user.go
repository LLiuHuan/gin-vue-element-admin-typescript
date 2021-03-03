package model

import "54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/global"

//
type User struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type TUser struct {
	global.GMODEL
	UserId       int64  `gorm:"comment:雪花id" json:"user_id"`
	Username     string `gorm:"comment:用户名" json:"username"`
	Password     string `gorm:"comment:密码" json:"password"`
	Email        string `gorm:"comment:邮箱" json:"email"`
	Gender       uint   `gorm:"comment:性别" json:"gender"`
	Avatar       string `gorm:"comment:头像" json:"avatar"`
	Name         string `gorm:"comment:名称" json:"name"`
	Introduction string `gorm:"comment:签名" json:"introduction"`
	Phone        string `gorm:"comment:手机号" json:"phone"`
	Roles        string `gorm:"comment:角色" json:"roles"`
}

func (TUser) TableName() string {
	return "t_user"
}

//type SysUser struct {
//	Username    string       `json:"userName" gorm:"comment:用户登录名"`
//	Password    string       `json:"-"  gorm:"comment:用户登录密码"`
//	NickName    string       `json:"nickName" gorm:"default:系统用户;comment:用户昵称" `
//	HeaderImg   string       `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
//	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
//	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:用户角色ID"`
//}
