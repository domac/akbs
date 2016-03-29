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
		MaxIdle:     3,                 //最大空闲数
		IdleTimeout: 240 * time.Second, //空闲超时时间
		//连接方法
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
		//从连接池中获取redis连接时做的校验操作
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return &RedisConn{
		pool: pool,
	}, err
}

//操作命令
func (r *RedisConn) Do(cmd string, args ...interface{}) (reply interface{}, err error) {

	//从连接池中获取redis连接
	conn := r.pool.Get()
	defer conn.Close()

	return conn.Do(cmd, args...)
}

// 获取连接
func (r *RedisConn) Get() redis.Conn {
	return r.pool.Get()
}
