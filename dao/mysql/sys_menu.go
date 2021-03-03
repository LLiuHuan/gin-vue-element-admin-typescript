package mysql

import (
	g "54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/global"
	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/model"
)

//type User struct {
//	ID   int64
//	Name string `gorm:"default:'小王子'"`
//	Age  int64
//}

// 获取全部路由
func GetMenuTreeMap(roleName string) (err error, treeMap map[string][]model.TMenu) {
	var AllMenus []model.TMenu
	treeMap = make(map[string][]model.TMenu)
	err = g.GROM.Where("id in (?)", g.GROM.Where("role = ?", roleName).Table("t_role_menu").Select("menu")).Find(&AllMenus).Error
	for _, v := range AllMenus {
		treeMap[v.PId] = append(treeMap[v.PId], v)
	}
	return
}

// 获取路由
func GetMenuByPId(pid string) (err error, treeMenu []model.TMenuByPId) {
	err = g.GROM.Where("p_id = ?", pid).Order("sort asc").Find(&treeMenu).Error
	return
}

func GetMenuBiCount() (err error, treeMenu []model.TMenusCount) {
	err = g.GROM.Raw("select id, (select count(1) from t_menu where p_id = tm.id) count from t_menu tm").Scan(&treeMenu).Error
	return
}
