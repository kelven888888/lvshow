package global

import (
	"ginshop.com/config"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	SHOP_DB     *gorm.DB
	SHOP_LOG    *zap.Logger
	SHOP_VP     *viper.Viper
	SHOP_CONFIG config.Server
	SHOP_REDIS  *redis.Client
	BlackCache  local_cache.Cache
)
