package initialize

import (
	"Miyo_Assignment/global"
	"go.uber.org/zap"
)

func InitZap() {
	global.GV_LOG = zap.NewProduction()
}
