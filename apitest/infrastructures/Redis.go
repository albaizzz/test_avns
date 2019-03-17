package infrastructures

import (
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var redisClient redis.Cmdable

// Redis defines parameter to open connection
type Redis struct {
	Host     string
	DB       int
	Port     int
	clientMu sync.Mutex
}

// OpenConnection if a function to open redis connection
func (r *Redis) OpenConnection() (redis.Cmdable, error) {
	if redisClient == nil {
		r.clientMu.Lock()
		connString := fmt.Sprintf("%s:%d", r.Host, r.Port)
		redisClient = redis.NewClient(&redis.Options{
			Addr: connString,
			DB:   r.DB,
		})

		r.clientMu.Unlock()
	}

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Error(err)
		return redisClient, err
	}

	return redisClient, nil
}

// Close closes redis cmdable
func (r *Redis) Close() error {
	if err := redisClient.(io.Closer).Close(); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// Get gets data by key from redis
func (r *Redis) Get(key string) (interface{}, error) {
	// Open redis connection
	client, err := r.OpenConnection()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// Get data from redis
	result := client.Get(key)
	if err := result.Err(); err != nil && err != redis.Nil {
		log.Error(err)
		return nil, err
	} else if err == redis.Nil {
		return nil, nil
	}

	return result.Val(), nil
}

// Set sets redis key, value and ttl
func (r *Redis) Set(key string, value interface{}, ttl time.Duration) error {
	// Open redis connection
	client, err := r.OpenConnection()
	if err != nil {
		log.Error(err)
		return err
	}

	// Set data from redis
	result := client.Set(key, value, ttl)
	if err := result.Err(); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// Set sets redis key, value and ttl
func (r *Redis) Del(key string) error {
	// Open redis connection
	client, err := r.OpenConnection()
	if err != nil {
		log.Error(err)
		return err
	}

	// Set data from redis
	result := client.Del(key)
	if err := result.Err(); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (r *Redis) Ping() error {

	if redisClient == nil {
		_, err := r.OpenConnection()
		if err != nil {
			return err
		}
	}
	return redisClient.Ping().Err()
}
