package initialize

import (
	"Miyo_Assignment/global"
	"github.com/spf13/viper"
	"os"
	"path"
)

func LoadConfig() {
	global.GV_VIP = viper.New()
	workDir, _ := os.Getwd()
	global.GV_VIP.SetConfigFile(path.Join(workDir, "config/viper.yaml"))
	if err := global.GV_VIP.ReadInConfig(); err != nil {

	}
}
