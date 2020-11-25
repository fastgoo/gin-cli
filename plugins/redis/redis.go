package redis

import (
	"context"
	"github.com/caarlos0/env/v6"
	"log"
	"time"
)

type config struct {
	Addr      []string `env:"REDIS_ADDRS,required" envSeparator:","`
	Password  string   `env:"REDIS_PW" envDefault:""`
	DB        int      `env:"REDIS_DB" envDefault:"0"`
	IsCluster bool
}

var (
	cfg = &config{}
	ctx = context.Background()
)

func init() {
	if err := env.Parse(cfg); err != nil {
		log.Fatal(err)
	}
	if len(cfg.Addr) > 1 {
		cfg.IsCluster = true
	}
	Initialize()
}

func Initialize() {
	var err error
	if !cfg.IsCluster {
		err = newClient()
	} else {
		err = newCluster()
	}
	if err != nil {
		log.Fatal("Unable to connect to redis " + err.Error())
	}
}

func Get(key string) string {
	if cfg.IsCluster {
		return redisCluster.Get(key)
	}
	return redisClient.Get(key)
}

func GetVal(key string, data interface{}) error {
	if cfg.IsCluster {
		return redisCluster.GetVal(key, data)
	}
	return redisClient.GetVal(key, data)
}

func Exists(key ...string) error {
	if cfg.IsCluster {
		return redisCluster.Exists(key...)
	}
	return redisClient.Exists(key...)
}

func Set(key string, val interface{}, expireTime time.Duration) error {
	if cfg.IsCluster {
		return redisCluster.Set(key, val, expireTime)
	}
	return redisClient.Set(key, val, expireTime)
}

func Del(key ...string) error {
	if cfg.IsCluster {
		return redisCluster.Del(key...)
	}
	return redisClient.Del(key...)
}

func HGet(key, field string) (string, error) {
	if cfg.IsCluster {
		return redisCluster.HGet(key, field)
	}
	return redisClient.HGet(key, field)
}

func HGetAll(key string) (map[string]string, error) {
	if cfg.IsCluster {
		return redisCluster.HGetAll(key)
	}
	return redisClient.HGetAll(key)
}

func HSet(key string, val map[string]string) error {
	if cfg.IsCluster {
		return redisCluster.HSet(key, val)
	}
	return redisClient.HSet(key, val)
}

func HDel(key string, fields ...string) error {
	if cfg.IsCluster {
		return redisCluster.HDel(key, fields...)
	}
	return redisClient.HDel(key, fields...)
}

func HExists(key string, field string) error {
	if cfg.IsCluster {
		return redisCluster.HExists(key, field)
	}
	return redisClient.HExists(key, field)
}

func HIncrBy(key string, field string, num int64) (int64, error) {
	if cfg.IsCluster {
		return redisCluster.HIncrBy(key, field, num)
	}
	return redisClient.HIncrBy(key, field, num)
}

func HKeys(key string) ([]string, error) {
	if cfg.IsCluster {
		return redisCluster.HKeys(key)
	}
	return redisClient.HKeys(key)
}

func HLen(key string) (int64, error) {
	if cfg.IsCluster {
		return redisCluster.HLen(key)
	}
	return redisClient.HLen(key)
}

func HMGet(key string, fields ...string) ([]interface{}, error) {
	if cfg.IsCluster {
		return redisCluster.HMGet(key, fields...)
	}
	return redisClient.HMGet(key, fields...)
}

func HMSet(key string, values ...interface{}) (bool, error) {
	if cfg.IsCluster {
		return redisCluster.HMSet(key, values...)
	}
	return redisClient.HMSet(key, values...)
}
