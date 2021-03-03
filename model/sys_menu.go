package model

import (
	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/global"
	"gorm.io/gorm"
)

type TMenu struct {
	ID          string         `gorm:"primarykey" json:"id"`
	CreateTime  global.MyTime  `db:"create_time" json:"create_time"`
	UpdateTime  global.MyTime  `db:"update_time" json:"update_time"`
	DeletedTime gorm.DeletedAt `gorm:"delete_time" json:"-"`
	PId         string         `gorm:"comment:父菜单ID" json:"pid"`
	Path        string         `gorm:"comment:路由path" json:"path"`
	Component   string         `gorm:"comment:对应前端文件路径" json:"component"`
	Redirect    string         `gorm:"comment:重定向" json:"redirect"`
	Name        string         `gorm:"comment:路由name" json:"name"`
	Level       int            `gorm:"comment:级别" json:"level"`
	Sort        int            `gorm:"comment:排序标记" json:"sort"`
	Meta        `json:"meta" gorm:"comment:附加属性"`
	Children    []TMenu `gorm:"-" json:"children"`
}

func (TMenu) TableName() string {
	return "t_menu"
}

type Meta struct {
	MetaTitle      string `gorm:"comment:菜单名" json:"title"`
	MetaIcon       string `gorm:"comment:菜单图标" json:"icon"`
	MetaRoles      string `gorm:"comment:菜单权限? 暂时可能不需要" json:"roles"`
	MetaAlwaysShow bool   `gorm:"comment:一直留在根菜单" json:"always_show"`
	MetaAffix      bool   `gorm:"comment:贴上？" json:"affix"`
	MetaHidden     bool   `gorm:"comment:显示？" json:"hidden"`
	MetaCache      bool   `gorm:"comment:隐藏？" json:"cache"`
	MetaBreadcrumb bool   `gorm:"comment:面包屑导航？" json:"breadcrumb"`
}

type TMenuByPId struct {
	ID             string         `gorm:"primarykey" json:"id"`
	CreateTime     global.MyTime  `db:"create_time" json:"create_time"`
	UpdateTime     global.MyTime  `db:"update_time" json:"update_time"`
	DeletedTime    gorm.DeletedAt `gorm:"delete_time" json:"-"`
	PId            string         `gorm:"comment:父菜单ID" json:"pid"`
	Path           string         `gorm:"comment:路由path" json:"path"`
	Component      string         `gorm:"comment:对应前端文件路径" json:"component"`
	Redirect       string         `gorm:"comment:重定向" json:"redirect"`
	Name           string         `gorm:"comment:路由name" json:"name"`
	Level          int            `gorm:"comment:级别" json:"level"`
	Sort           int            `gorm:"comment:排序标记" json:"sort"`
	MetaTitle      string         `gorm:"comment:菜单名" json:"title"`
	MetaIcon       string         `gorm:"comment:菜单图标" json:"icon"`
	MetaRoles      string         `gorm:"comment:菜单权限? 暂时可能不需要" json:"roles"`
	MetaAlwaysShow bool           `gorm:"comment:一直留在根菜单" json:"always_show"`
	MetaAffix      bool           `gorm:"comment:贴上？" json:"affix"`
	MetaHidden     bool           `gorm:"comment:显示？" json:"hidden"`
	MetaCache      bool           `gorm:"comment:隐藏？" json:"cache"`
	MetaBreadcrumb bool           `gorm:"comment:面包屑导航？" json:"breadcrumb"`
	HasChildren    bool           `json:"hasChildren"`
}

type TMenusCount struct {
	Id    string `gorm:"primarykey" json:"id"`
	Count int    `gorm:"comment:数量" json:"count"`
}

func (TMenuByPId) TableName() string {
	return "t_menu"
}
