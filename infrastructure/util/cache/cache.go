package cache

import (
	"github.com/go-redis/redis"
	"link-art-api/infrastructure/config"
	"log"
)

var CACHE *redis.Client

func SetupCache() {
	CACHE = redis.NewClient(&redis.Options{
		Addr:     config.RedisConfig.Host,
		Password: config.RedisConfig.Password,
		DB:       config.RedisConfig.DB,
	})
	_, err := CACHE.Ping().Result()
	if err != nil {
		log.Panic(err)
	}
}
