package redis

import (
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

