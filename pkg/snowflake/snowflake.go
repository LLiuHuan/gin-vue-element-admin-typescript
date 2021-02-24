package snowflake

import g "54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/global"

func GenID() int64 {
	return g.NODE.Generate().Int64()
}
