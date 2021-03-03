package server

import (
	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/dao/mysql"
	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/model"
)

func getChildrenList(menu *model.TMenu, treeMap map[string][]model.TMenu) (err error) {
	menu.Children = treeMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

// GetMenu 根据角色获取动态路由
func GetMenu(roleName string) (err error, menus []model.TMenu) {
	err, menuTree := mysql.GetMenuTreeMap(roleName)
	if err != nil {
		return
	}
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i], menuTree)
	}
	return err, menus
}

// GetMenus 获取菜单树
func GetMenuByPId(pid string) (err error, menus []model.TMenuByPId) {
	err, menuTree := mysql.GetMenuByPId(pid)
	if err != nil {
		return
	}
	err, menuCounts := mysql.GetMenuBiCount()
	for index, val := range menuTree {
		val.HasChildren = false
		for _, val2 := range menuCounts {
			if val.ID == val2.Id && val2.Count > 0 {
				menuTree[index].HasChildren = true
			}
		}
	}
	return err, menuTree
}
