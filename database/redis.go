package database

import (
	"fmt"
	"log"
	"strconv"
	"time"

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

func (r *Redis) IsTokenValid(token string) bool {
	_, err := r.redis.Get("token:" + token).Result()
	return err != redis.Nil
}

func (r *Redis) AddToken(token string, userId int64, time time.Duration) {
	r.redis.Set("token:"+token, fmt.Sprint(userId), time)
}

func (r *Redis) GetUserId(token string) int64 {
	val, _ := r.redis.Get("token:" + token).Result()
	n, _ := strconv.ParseInt(val, 10, 64)
	return n
}
