package model

import "54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/global"

type TRole struct {
	global.GMODEL
	Name  string `gorm:"comment:角色名称" json:"name"`
	Desc  string `gorm:"comment:角色备注" json:"desc"`
	PsIds string `gorm:"comment:权限标识id" json:"ps_ids"`
}

func (TRole) TableName() string {
	return "t_role"
}

type TRoleMenu struct {
	Id   uint   `gorm:"comment:主键id" json:"id"`
	Role string `gorm:"comment:角色" json:"role"`
	Menu string `gorm:"comment:菜单" json:"menu"`
}

func (t TRoleMenu) TableName() string {
	return "t_role_menu"
}
