package globals

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"shop-srv/basic/config"
)

var (
	AppConfig *config.Config
	DB        *gorm.DB
	Rdb       *redis.Client
)
