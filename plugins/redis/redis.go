package redis

import (
	"errors"
	"log"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/go-redis/redis"
)

type config struct {
	Addr      []string `env:"REDIS_ADDRS,required" envSeparator:","`
	Password  string   `env:"REDIS_PW" envDefault:""`
	DB        int      `env:"REDIS_DB" envDefault:"0"`
	IsCluster bool
}

type redisCli struct {
	client interface{}
}

var (
	cfg         = &config{}
	redisClient = &redisCli{}
)

func init() {
	var err error
	if err = env.Parse(cfg); err != nil {
		log.Fatal(err)
	}
	if len(cfg.Addr) > 1 {
		cfg.IsCluster = true
		err = newCluster()
	} else {
		err = newClient()
	}
	if err != nil {
		log.Fatal("Unable to connect to redis " + err.Error())
	}
}

func newClient() error {
	redisClient.client = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr[0],
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	if err := redisClient.client.(redis.Cmdable).Ping().Err(); err != nil {
		return err
	}
	return nil
}

func newCluster() error {
	redisClient.client = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    cfg.Addr,
		Password: cfg.Password,
		//ReadOnly: false,
		//DialTimeout:  50 * time.Microsecond, // 设置连接超时
		//ReadTimeout:  50 * time.Microsecond, // 设置读取超时
		//WriteTimeout: 50 * time.Microsecond, // 设置写入超时
	})
	if err := redisClient.client.(redis.Cmdable).Ping().Err(); err != nil {
		return err
	}
	return nil
}

func Get(key string) string {
	val, err := redisClient.client.(redis.Cmdable).Get(key).Result()
	if err != nil {
		if err != redis.Nil {
			//log.Panic("redis get error, " + err.Error())
		}
		return ""
	}
	return val
}

func GetVal(key string, data interface{}) error {
	get := redisClient.client.(redis.Cmdable).Get(key)
	if err := get.Err(); err != nil {
		if err != redis.Nil {
			//log.Panic("redis getval error, " + err.Error())
		}
		return err
	}
	err := get.Scan(&data)
	return err
}

func Set(key string, val interface{}, expireTime time.Duration) error {
	err := redisClient.client.(redis.Cmdable).SetNX(key, val, expireTime*time.Second).Err()
	if err != nil {
		//log.Panic("redis set error, " + err.Error())
		return err
	}
	return nil
}

func Exists(key ...string) error {
	err := redisClient.client.(redis.Cmdable).Exists(key...).Err()
	if err != nil {
		//log.Panic("redis exists error, " + err.Error())
		return err
	}
	return nil
}

func Del(key ...string) error {
	err := redisClient.client.(redis.Cmdable).Del(key...).Err()
	if err != nil {
		//log.Panic("redis del error, " + err.Error())
		return err
	}
	return nil
}

func HGet(key, field string) (string, error) {
	val, err := redisClient.client.(redis.Cmdable).HGet(key, field).Result()
	if err != nil {
		if err != redis.Nil {
			//log.Panic("redis hget error, " + err.Error())
		}
		return "", err
	}
	return val, nil
}

func HGetAll(key string) (map[string]string, error) {
	val, err := redisClient.client.(redis.Cmdable).HGetAll(key).Result()
	if err != nil {
		if err != redis.Nil {
			//log.Panic("redis hgetall error, " + err.Error())
		}
		return nil, err
	}
	return val, nil
}

func HSet(key, field string, val map[string]string) error {
	err := redisClient.client.(redis.Cmdable).HSet(key, field, val).Err()
	if err != nil {
		//log.Panic("redis hset error, " + err.Error())
		return err
	}
	return nil
}

func HDel(key string, fields ...string) error {
	err := redisClient.client.(redis.Cmdable).HDel(key, fields...).Err()
	if err != nil {
		//log.Panic("redis hdel error, " + err.Error())
		return err
	}
	return nil
}

func HExists(key string, field string) error {
	has, err := redisClient.client.(redis.Cmdable).HExists(key, field).Result()
	if err != nil {
		//log.Panic("redis HExists error, " + err.Error())
		return err
	}
	if !has {
		return errors.New("redis key field not found")
	}
	return nil
}

func HIncrBy(key string, field string, num int64) (int64, error) {
	resNum, err := redisClient.client.(redis.Cmdable).HIncrBy(key, field, num).Result()
	if err != nil {
		//log.Panic("redis HIncrBy error, " + err.Error())
		return 0, err
	}
	return resNum, nil
}

func HKeys(key string) ([]string, error) {
	keys, err := redisClient.client.(redis.Cmdable).HKeys(key).Result()
	if err != nil {
		//log.Panic("redis HKeys error, " + err.Error())
		return nil, err
	}
	return keys, nil
}

func HLen(key string) (int64, error) {
	resLen, err := redisClient.client.(redis.Cmdable).HLen(key).Result()
	if err != nil {
		//log.Panic("redis HLen error, " + err.Error())
		return 0, err
	}
	return resLen, nil
}

func HMGet(key string, fields ...string) ([]interface{}, error) {
	res, err := redisClient.client.(redis.Cmdable).HMGet(key, fields...).Result()
	if err != nil {
		//log.Panic("redis HMGet error, " + err.Error())
		return nil, err
	}
	return res, nil
}

func HMSet(key string, fields map[string]interface{}) error {
	_, err := redisClient.client.(redis.Cmdable).HMSet(key, fields).Result()
	if err != nil {
		//log.Panic("redis HMSet error, " + err.Error())
		return err
	}
	return nil
}
