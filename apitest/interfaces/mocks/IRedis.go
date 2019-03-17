package mocks

import (
	"strings"
	"time"

	"github.com/go-redis/redis"
	mock "github.com/stretchr/testify/mock"
)

// IRedis is an autogenerated mock type for the IBankRepository type
type IRedis struct {
	mock.Mock
}

// Get indicates an expected call of Get Data from redis
func (_m *IRedis) Get(key string) (interface{}, error) {
	// ret := _m.Called()

	// var r0 interface{}
	// if rf, ok := ret.Get(0).(func(int) interface{}); ok {
	// 	r0 = rf(0)
	// } else {
	// 	r0 = ret.Get(0).(interface{})
	// }

	// var r1 error
	// if rf, ok := ret.Get(1).(func(int) error); ok {
	// 	r1 = rf(0)
	// } else {
	// 	r1 = ret.Error(1)
	// }

	// return r0, r1
	if strings.Contains(key, "token") {
		return key, nil
	}
	return 0, nil
}

// Set indicates an expected call of set data to redis
func (_m *IRedis) Set(key string, value interface{}, ttl time.Duration) error {
	return nil
}

// OpenConnection indicates an expected call of OpenConnection redis
func (_m *IRedis) OpenConnection() (redis.Cmdable, error) {
	var r redis.Client
	return &r, nil
}

func (_m *IRedis) Ping() error {
	return nil
}
