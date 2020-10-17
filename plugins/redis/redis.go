package redis

import (
	"context"
	"github.com/caarlos0/env/v6"
	"github.com/go-redis/redis/v8"
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
	cfg          = &config{}
	ctx          = context.Background()
	redisCluster *redis.ClusterClient
	redisClient  *redis.Client
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
		err = newClusterClient()
	}
	if err != nil {
		log.Fatal("Unable to connect to redis " + err.Error())
	}
}

func newClient() error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr[0],
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	if err := redisClient.Ping(ctx).Err(); err != nil {
		return err
	}
	return nil
}

func newClusterClient() error {
	redisCluster = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    cfg.Addr,
		Password: cfg.Password,
		//ReadOnly: false,
		//DialTimeout:  50 * time.Microsecond, // 设置连接超时
		//ReadTimeout:  50 * time.Microsecond, // 设置读取超时
		//WriteTimeout: 50 * time.Microsecond, // 设置写入超时
	})
	if err := redisCluster.Ping(ctx).Err(); err != nil {
		return err
	}
	return nil
}

func Get() string {
	val, err := redisClient.Get(ctx, "key").Result()
	if err != nil {
		if err == redis.Nil {
			return ""
		}
		log.Panic("redis get error, " + err.Error())
		return ""
	}
	return val
}

func GetVal(data interface{}) error {
	get := redisClient.Get(ctx, "key")
	if err := get.Err(); err != nil {
		if err == redis.Nil {
			return err
		}
		log.Panic("redis get error, " + err.Error())
		return err
	}
	err := get.Scan(&data)
	return err
}

func Set(key string, val interface{}, expireTime time.Duration) error {
	err := redisClient.SetNX(ctx, key, val, expireTime*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func Exists(key ...string) error {
	err := redisClient.Exists(ctx, key...).Err()
	if err != nil {
		return err
	}
	return nil
}

func Del(key ...string) error {
	err := redisClient.Del(ctx, key...).Err()
	if err != nil {
		return err
	}
	return nil
}

func HGet() {

}

func HGetAll() {

}

func HSet() {

}

func IncrBy() {

}

func Incr() {

}

func HIncrByFloat() {

}

func HIncrByFloat() {

}

func HIncrByFloat() {

}

func HIncrByFloat() {

}

func HIncrByFloat() {

}

func HIncrByFloat() {

}

func HIncrByFloat() {

}

func HIncrByFloat() {

}

func HIncrByFloat() {

}

func HIncrByFloat() {

}

func HIncrByFloat() {

}
