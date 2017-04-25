package cache

import (
	"time"

	"github.com/FlyCynomys/gear/conf"
	"github.com/FlyCynomys/tools/log"
	redis "github.com/go-redis/redis"
)

var RedisClient *redis.Client

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         conf.GetCfg().RedisAddress,
		Password:     "",
		DB:           0,
		MaxRetries:   3,
		DialTimeout:  30,
		ReadTimeout:  10,
		WriteTimeout: 10,
		PoolSize:     10,
	})
	go func() {
		ti := time.NewTicker(time.Second * 60)
		for {
			select {
			case <-ti.C:
				_, err := RedisClient.Ping().Result()
				if err != nil {
					log.Error(err)
				}
			}
		}
	}()
}
