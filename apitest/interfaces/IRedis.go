package interfaces

import (
	"time"

	"github.com/go-redis/redis"
)

//IRedis redis interface
type IRedis interface {
	OpenConnection() (redis.Cmdable, error)
	Get(key string) (interface{}, error)
	Del(key string) error
	Set(key string, value interface{}, ttl time.Duration) error
	Ping() error
}
