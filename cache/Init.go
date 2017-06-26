package cache

import (
	"time"

	"github.com/FlyCynomys/gear/conf"
	"github.com/FlyCynomys/tools/log"
	redis "github.com/go-redis/redis"
)

type CacheClient struct {
	RedisClient *redis.Client
}

var cc *CacheClient

func Init() {
	cc = new(CacheClient)
	cc.RedisClient = redis.NewClient(&redis.Options{
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
				_, err := cc.RedisClient.Ping().Result()
				if err != nil {
					log.Error(err)
				}
			}
		}
	}()
}
