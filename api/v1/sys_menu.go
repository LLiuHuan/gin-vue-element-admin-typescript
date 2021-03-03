package v1

import (
	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/model/response"
	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/server"
	"github.com/gin-gonic/gin"
)

// GetMenu 根据角色获取动态路由
// @Summary 根据角色获取动态路由
// @Description 根据角色获取动态路由
// @Tags AuthorityMenu
// @Accept application/json
// @Produce application/json
// @Param who query request.ParamRole true "角色名称"
// @Security ApiKeyAuth
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/menu/getMenu [get]
func GetMenu(c *gin.Context) {
	roleName := c.DefaultQuery("roleName", "")
	if roleName == "" {
		response.FailWithMessage("缺少角色参数", c)
		return
	}
	err, menus := server.GetMenu(roleName)
	if err != nil {
		response.FailWithMessage("查询菜单失败", c)
		return
	}
	response.OkWithData(menus, c)
}

// GetMenuByPId 获取菜单树
// @Summary 获取菜单树
// @Description 获取动态菜单树
// @Tags AuthorityMenu
// @Accept application/json
// @Produce application/json
// @Param who query request.ParamRole true "角色名称"
// @Security ApiKeyAuth
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/menu/getMenuByPId [get]
func GetMenuByPId(c *gin.Context) {
	pid := c.DefaultQuery("pid", "")
	if pid == "" {
		response.FailWithMessage("缺少角色参数", c)
		return
	}
	err, menus := server.GetMenuByPId(pid)
	if err != nil {
		response.FailWithMessage("查询菜单失败", c)
		return
	}
	response.OkWithData(menus, c)
}
