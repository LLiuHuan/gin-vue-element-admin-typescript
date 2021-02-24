package initialize

import (
	"fmt"

	"54cc.cc/LLiuHuan/gin-vue-element-admin-typescript/model"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

// Settings 初始化配置文件
func Settings(filePath string) (err error) {
	// 方式1：指定配置文件名和配置文件位置 viper自行查找可用的配置文件
	//viper.SetConfigName("config") // 指定配置名称（不需要带后缀）
	//viper.AddConfigPath(".")      // 指定查找配置文件的路径
	//viper.AddConfigPath("./conf")
	// 远程查找
	//viper.SetConfigType("yaml")   // 指定文件类型（专用于从远程获取配置信息时指定配置文件类型的）

	// 方式2：直接指定配置文件路径
	viper.SetConfigFile(filePath) // 指定配置文件

	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {            // 读取配置信息失败
		fmt.Printf("viper.ReadInConfig failed, err: %v\n", err)
		return err
		//panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := viper.Unmarshal(model.SettingsConf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err: %#v \n", err)
	}

	// 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了")
		if err := viper.Unmarshal(model.SettingsConf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err: %#v \n", err)
		}
	})
	return
}
