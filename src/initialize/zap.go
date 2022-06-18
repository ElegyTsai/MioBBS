package initialize

import (
	"Miyo_Assignment/global"
	"go.uber.org/zap"
)

func InitZap() {
	logger, _ := zap.NewProduction()
	global.GV_LOG = logger.Sugar()
}
