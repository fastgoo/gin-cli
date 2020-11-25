package redis

import (
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisClusterCli struct {
	client *redis.ClusterClient
}

var redisCluster = &redisClusterCli{}

func newCluster() error {
	redisCluster.client = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    cfg.Addr,
		Password: cfg.Password,
		//ReadOnly: false,
		//DialTimeout:  50 * time.Microsecond, // 设置连接超时
		//ReadTimeout:  50 * time.Microsecond, // 设置读取超时
		//WriteTimeout: 50 * time.Microsecond, // 设置写入超时
	})
	if err := redisCluster.client.Ping(ctx).Err(); err != nil {
		return err
	}
	return nil
}

func (r *redisClusterCli) Get(key string) string {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			//log.Panic("redis get error, " + err.Error())
		}
		return ""
	}
	return val
}

func (r *redisClusterCli) GetVal(key string, data interface{}) error {
	get := r.client.Get(ctx, key)
	if err := get.Err(); err != nil {
		if err != redis.Nil {
			//log.Panic("redis getval error, " + err.Error())
		}
		return err
	}
	err := get.Scan(&data)
	return err
}

func (r *redisClusterCli) Set(key string, val interface{}, expireTime time.Duration) error {
	err := r.client.SetNX(ctx, key, val, expireTime*time.Second).Err()
	if err != nil {
		//log.Panic("redis set error, " + err.Error())
		return err
	}
	return nil
}

func (r *redisClusterCli) Exists(key ...string) error {
	err := r.client.Exists(ctx, key...).Err()
	if err != nil {
		//log.Panic("redis exists error, " + err.Error())
		return err
	}
	return nil
}

func (r *redisClusterCli) Del(key ...string) error {
	err := r.client.Del(ctx, key...).Err()
	if err != nil {
		//log.Panic("redis del error, " + err.Error())
		return err
	}
	return nil
}

func (r *redisClusterCli) HGet(key, field string) (string, error) {
	val, err := r.client.HGet(ctx, key, field).Result()
	if err != nil {
		if err != redis.Nil {
			//log.Panic("redis hget error, " + err.Error())
		}
		return "", err
	}
	return val, nil
}

func (r *redisClusterCli) HGetAll(key string) (map[string]string, error) {
	val, err := r.client.HGetAll(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			//log.Panic("redis hgetall error, " + err.Error())
		}
		return nil, err
	}
	return val, nil
}

func (r *redisClusterCli) HSet(key string, val map[string]string) error {
	err := r.client.HSet(ctx, key, val).Err()
	if err != nil {
		//log.Panic("redis hset error, " + err.Error())
		return err
	}
	return nil
}

func (r *redisClusterCli) HDel(key string, fields ...string) error {
	err := r.client.HDel(ctx, key, fields...).Err()
	if err != nil {
		//log.Panic("redis hdel error, " + err.Error())
		return err
	}
	return nil
}

func (r *redisClusterCli) HExists(key string, field string) error {
	has, err := r.client.HExists(ctx, key, field).Result()
	if err != nil {
		//log.Panic("redis HExists error, " + err.Error())
		return err
	}
	if !has {
		return errors.New("redis key field not found")
	}
	return nil
}

func (r *redisClusterCli) HIncrBy(key string, field string, num int64) (int64, error) {
	resNum, err := r.client.HIncrBy(ctx, key, field, num).Result()
	if err != nil {
		//log.Panic("redis HIncrBy error, " + err.Error())
		return 0, err
	}
	return resNum, nil
}

func (r *redisClusterCli) HKeys(key string) ([]string, error) {
	keys, err := r.client.HKeys(ctx, key).Result()
	if err != nil {
		//log.Panic("redis HKeys error, " + err.Error())
		return nil, err
	}
	return keys, nil
}

func (r *redisClusterCli) HLen(key string) (int64, error) {
	resLen, err := r.client.HLen(ctx, key).Result()
	if err != nil {
		//log.Panic("redis HLen error, " + err.Error())
		return 0, err
	}
	return resLen, nil
}

func (r *redisClusterCli) HMGet(key string, fields ...string) ([]interface{}, error) {
	res, err := r.client.HMGet(ctx, key, fields...).Result()
	if err != nil {
		//log.Panic("redis HMGet error, " + err.Error())
		return nil, err
	}
	return res, nil
}

func (r *redisClusterCli) HMSet(key string, values ...interface{}) (bool, error) {
	status, err := r.client.HMSet(ctx, key, values...).Result()
	if err != nil {
		//log.Panic("redis HMSet error, " + err.Error())
		return false, err
	}
	return status, nil
}
