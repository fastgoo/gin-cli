package redis

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/go-redis/redis/v8"
)

//后续如果需要加方法，可以参考一下官网的用法
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
	ctx         = context.Background()
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

// redis单机
func newClient() error {
	redisClient.client = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr[0],
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	if err := redisClient.client.(redis.Cmdable).Ping(ctx).Err(); err != nil {
		return err
	}
	return nil
}

// redis集群
func newCluster() error {
	redisClient.client = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    cfg.Addr,
		Password: cfg.Password,
		//ReadOnly: false,
		//DialTimeout:  50 * time.Microsecond, // 设置连接超时
		//ReadTimeout:  50 * time.Microsecond, // 设置读取超时
		//WriteTimeout: 50 * time.Microsecond, // 设置写入超时
	})
	if err := redisClient.client.(redis.Cmdable).Ping(ctx).Err(); err != nil {
		return err
	}
	return nil
}

func Get(key string) string {
	val, err := redisClient.client.(redis.Cmdable).Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			//log.Panic("redis get error, " + err.Error())
		}
		return ""
	}
	return val
}

func MGet(key ...string) []interface{} {
	val, err := redisClient.client.(redis.Cmdable).MGet(ctx, key...).Result()
	if err != nil {
		if err != redis.Nil {
			//log.Panic("redis get error, " + err.Error())
		}
		return []interface{}{}
	}
	return val
}

func GetVal(key string, data interface{}) error {
	get := redisClient.client.(redis.Cmdable).Get(ctx, key)
	if err := get.Err(); err != nil {
		if err != redis.Nil {
			//log.Panic("redis getval error, " + err.Error())
		}
		return err
	}
	err := get.Scan(data)
	return err
}

func Set(key string, val interface{}, expireTime time.Duration) error {
	err := redisClient.client.(redis.Cmdable).Set(ctx, key, val, expireTime*time.Second).Err()
	if err != nil {
		//log.Panic("redis set error, " + err.Error())
		return err
	}
	return nil
}

func SetNX(key string, val interface{}, expireTime time.Duration) error {
	err := redisClient.client.(redis.Cmdable).SetNX(ctx, key, val, expireTime*time.Second).Err()
	if err != nil {
		//log.Panic("redis set error, " + err.Error())
		return err
	}
	return nil
}

func MSet(pairs ...interface{}) error {
	err := redisClient.client.(redis.Cmdable).MSet(ctx, pairs...).Err()
	if err != nil {
		//log.Panic("redis set error, " + err.Error())
		return err
	}
	return nil
}

func MSetNX(pairs ...interface{}) error {
	err := redisClient.client.(redis.Cmdable).MSetNX(ctx, pairs...).Err()
	if err != nil {
		//log.Panic("redis set error, " + err.Error())
		return err
	}
	return nil
}

func Exists(key ...string) error {
	k, err := redisClient.client.(redis.Cmdable).Exists(ctx, key...).Result()
	if err != nil {
		return err
	}
	if int(k) != len(key) {
		return errors.New("redis somekey not exist")
	}
	return nil
}

func Del(key ...string) error {
	err := redisClient.client.(redis.Cmdable).Del(ctx, key...).Err()
	if err != nil {
		//log.Panic("redis del error, " + err.Error())
		return err
	}
	return nil
}

func HGet(key, field string) (string, error) {
	val, err := redisClient.client.(redis.Cmdable).HGet(ctx, key, field).Result()
	if err != nil {
		if err != redis.Nil {
			//log.Panic("redis hget error, " + err.Error())
		}
		return "", err
	}
	return val, nil
}

func HGetAll(key string) (map[string]string, error) {
	val, err := redisClient.client.(redis.Cmdable).HGetAll(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			//log.Panic("redis hgetall error, " + err.Error())
		}
		return nil, err
	}
	return val, nil
}

func HSet(key, field string, val interface{}) error {
	err := redisClient.client.(redis.Cmdable).HSet(ctx, key, field, val).Err()
	if err != nil {
		//log.Panic("redis hset error, " + err.Error())
		return err
	}
	return nil
}

func HDel(key string, fields ...string) error {
	err := redisClient.client.(redis.Cmdable).HDel(ctx, key, fields...).Err()
	if err != nil {
		//log.Panic("redis hdel error, " + err.Error())
		return err
	}
	return nil
}

func HExists(key string, field string) error {
	has, err := redisClient.client.(redis.Cmdable).HExists(ctx, key, field).Result()
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
	resNum, err := redisClient.client.(redis.Cmdable).HIncrBy(ctx, key, field, num).Result()
	if err != nil {
		//log.Panic("redis HIncrBy error, " + err.Error())
		return 0, err
	}
	return resNum, nil
}

func HKeys(key string) ([]string, error) {
	keys, err := redisClient.client.(redis.Cmdable).HKeys(ctx, key).Result()
	if err != nil {
		//log.Panic("redis HKeys error, " + err.Error())
		return nil, err
	}
	return keys, nil
}

func HLen(key string) (int64, error) {
	resLen, err := redisClient.client.(redis.Cmdable).HLen(ctx, key).Result()
	if err != nil {
		//log.Panic("redis HLen error, " + err.Error())
		return 0, err
	}
	return resLen, nil
}

func HMGet(key string, fields ...string) ([]interface{}, error) {
	res, err := redisClient.client.(redis.Cmdable).HMGet(ctx, key, fields...).Result()
	if err != nil {
		//log.Panic("redis HMGet error, " + err.Error())
		return nil, err
	}
	return res, nil
}

func HMSet(key string, fields map[string]interface{}) error {
	_, err := redisClient.client.(redis.Cmdable).HMSet(ctx, key, fields).Result()
	if err != nil {
		//log.Panic("redis HMSet error, " + err.Error())
		return err
	}
	return nil
}

//原子自增
func Incr(key string, expireTime time.Duration) error {
	pipe := redisClient.client.(redis.Cmdable).TxPipeline()
	pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, expireTime)
	_, err := pipe.Exec(ctx)
	return err
}

//原子自减
func Decr(key string, expireTime time.Duration) error {
	pipe := redisClient.client.(redis.Cmdable).TxPipeline()
	pipe.Decr(ctx, key)
	pipe.Expire(ctx, key, expireTime)
	_, err := pipe.Exec(ctx)
	return err
}

//设置key的有效时间
func Expire(key string, expireTime time.Duration) error {
	_, err := redisClient.client.(redis.Cmdable).Expire(ctx, key, expireTime).Result()
	if err != nil {
		//log.Panic("redis HMSet error, " + err.Error())
		return err
	}
	return nil
}

//返回Key的有效时间返回秒
func TTL(key string) time.Duration {
	t, err := redisClient.client.(redis.Cmdable).TTL(ctx, key).Result()
	if err != nil {
		//log.Panic("redis HMSet error, " + err.Error())
		return 0
	}
	return t
}

//返回key的有效时间毫秒
func PTTL(key string) time.Duration {
	t, err := redisClient.client.(redis.Cmdable).PTTL(ctx, key).Result()
	if err != nil {
		//log.Panic("redis HMSet error, " + err.Error())
		return 0
	}
	return t
}
