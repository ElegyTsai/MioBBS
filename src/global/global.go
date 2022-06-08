package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GV_DB  *gorm.DB
	GV_LOG *zap.Logger
	GV_VIP *viper.Viper
)
