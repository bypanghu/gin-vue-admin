package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	pkgRedis "github.com/flipped-aurora/gin-vue-admin/server/pkg/redis"
	"github.com/redis/go-redis/v9"
)

func Redis() {
	redisClient, err := pkgRedis.InitRedisClient(global.GVA_CONFIG.Redis)
	if err != nil {
		panic(err)
	}
	global.GVA_REDIS = redisClient
}

func RedisList() {
	redisMap := make(map[string]redis.UniversalClient)

	for _, redisCfg := range global.GVA_CONFIG.RedisList {
		client, err := pkgRedis.InitRedisClient(redisCfg)
		if err != nil {
			panic(err)
		}
		redisMap[redisCfg.Name] = client
	}

	global.GVA_REDISList = redisMap
}
