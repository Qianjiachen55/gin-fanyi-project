package initalize

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func RedisClient() *redis.Client {
	ip := viper.GetString("redis.ip")
	port := viper.GetString("redis.port")
	DB := viper.GetInt("redis.DB")
	passwd := viper.GetString("redis.passwd")
	rdb := redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               fmt.Sprintf("%s:%s", ip, port),
		Dialer:             nil,
		OnConnect:          nil,
		Username:           "",
		Password:           passwd,
		DB:                 DB,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolFIFO:           false,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
		Limiter:            nil,
	})
	return rdb
}

func InitRedis(rdb *redis.Client) {
	//暂无配置
}
