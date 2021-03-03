package request

// ParamRole 根据角色获取菜单
type ParamRole struct {
	RoleName string `json:"roleName" binding:"required"`
}
