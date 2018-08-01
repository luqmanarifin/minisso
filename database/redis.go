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

func NewRedis(opt RedisOption) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     opt.Host + ":" + opt.Port,
		Password: opt.Password,
		DB:       opt.Database,
	})
	pong, err := client.Ping().Result()
	log.Println(pong, err)
	if err != nil {
		return &redis.Client{}, err
	}

	log.Printf("Success connecting Redis to %s:%s with pass %s\n", opt.Host, opt.Port, opt.Password)
	return client, nil
}
