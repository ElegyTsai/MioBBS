package initialize

import "Miyo_Assignment/global"

func Init() {
	InitZap()

	InitViper()

	InitGORM()

	global.GV_LOG.Infof("Successfully initilize configure")
}
