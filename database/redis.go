package database

import (
	"log"

	"github.com/go-redis/redis"
)

// Option holds all necessary options for Redis
type RedisOption struct {
	Host     string
	Port     string
	Password string
	Database int
}

type Redis struct {
	redis *redis.Client
}

func NewRedis(opt RedisOption) (Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     opt.Host + ":" + opt.Port,
		Password: opt.Password,
		DB:       opt.Database,
	})
	pong, err := client.Ping().Result()
	log.Println(pong, err)
	if err != nil {
		return Redis{}, err
	}

	log.Printf("Success connecting Redis to %s:%s with pass %s\n", opt.Host, opt.Port, opt.Password)
	return Redis{redis: client}, nil
}
