package initialize

import (
	"Miyo_Assignment/global"
	"Miyo_Assignment/utils"
	"github.com/spf13/viper"
	"path"
)

// 从viper.yaml文件中读取相关配置信息

func InitViper() {
	global.GV_VIP = viper.New()                                          //初始化viper
	workDir := utils.GetProjectPath()                                    //当前目录
	global.GV_VIP.SetConfigFile(path.Join(workDir, "config/viper.yaml")) //载入配置
	if err := global.GV_VIP.ReadInConfig(); err != nil {
		global.GV_LOG.Errorf("Fatal error of config file: %s \n ", err) //失败的话写入日志
		panic("stopped")
	}
	if err := global.GV_VIP.Unmarshal(&global.GV_Config); err != nil {
		global.GV_LOG.Errorf("Fatal error when mapping config file: %s \n ", err) //失败的话写入日志
		panic("stopped")
	}

}
