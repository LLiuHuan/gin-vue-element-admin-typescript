package initialize

import (
	"time"

	g "54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/global"

	sf "github.com/bwmarrin/snowflake"
)

func Snowflake(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return err
	}
	sf.Epoch = st.UnixNano() / 1000000
	g.NODE, err = sf.NewNode(machineID)
	return
}
