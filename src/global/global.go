package global

import (
	"Miyo_Assignment/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GV_DB     *gorm.DB
	GV_LOG    *zap.SugaredLogger
	GV_VIP    *viper.Viper
	GV_Config config.Config
)
