package initialize

import (
	"Miyo_Assignment/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
	初始化全局变量中的mysql
*/
func InitGORM() {
	MySQLConfig := global.GV_Config.Mysql
	if db, err := gorm.Open(mysql.Open(MySQLConfig.Dsn()), &gorm.Config{}); err != nil {
		global.GV_LOG.Errorf("fatal error when connecting gorm with mysql: %s \n", err)
		panic("stopped, please check 'viper.yaml'")
	} else {
		global.GV_DB = db
	}

}
