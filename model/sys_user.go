package model

//
type User struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

//type SysUser struct {
//	Username    string       `json:"userName" gorm:"comment:用户登录名"`
//	Password    string       `json:"-"  gorm:"comment:用户登录密码"`
//	NickName    string       `json:"nickName" gorm:"default:系统用户;comment:用户昵称" `
//	HeaderImg   string       `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
//	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
//	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:用户角色ID"`
//}
