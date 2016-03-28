package core

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

type RedisConn struct {
	pool *redis.Pool
}

func OpenRedis(network, address, password string) (*RedisConn, error) {
	_, err := redis.Dial(network, address)
	pool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(network, address)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return &RedisConn{
		pool: pool,
	}, err
}

func (r *RedisConn) Do(cmd string, args ...interface{}) (reply interface{}, err error) {

	conn := r.pool.Get()
	defer conn.Close()

	return conn.Do(cmd, args...)
}

// get a redis connection
func (r *RedisConn) Get() redis.Conn {
	return r.pool.Get()
}
