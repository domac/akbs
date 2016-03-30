package core

import (
	"errors"
	rc "github.com/chasex/redis-go-cluster"
	"github.com/garyburd/redigo/redis"
	"time"
)

//缓存接口
type Cache interface {
	Do(cmd string, args ...interface{}) (reply interface{}, err error)
}

//根据配置创建RedisCache
func NewRedisCache() (cache Cache, err error) {
	if ConfigInfo == nil {
		return nil, errors.New("can't load config info ... ")
	}
	isCluster := ConfigInfo.IsCluster //是否集群

	if !isCluster {
		addrs := ConfigInfo.GetRedisConfig().Addrs
		cache, err = OpenRedis(addrs) //单实例连接
	} else {
		addrs := ConfigInfo.GetRedisClusterConfig().Addrs
		cache, err = OpenRedisCluster([]string{addrs}) //集群连接
	}

	return cache, err
}

//单实例Redis连接
type RedisClient struct {
	pool *redis.Pool
}

//创建单实例Redis客户端
func OpenRedis(address string) (*RedisClient, error) {
	_, err := redis.Dial("tcp", address)
	pool := &redis.Pool{
		//连接方法
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", address)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		//从连接池中获取redis连接时做的校验操作
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	cr := ConfigInfo.GetRedisConfig()

	if cr.MaxIdle > 0 {
		pool.MaxIdle = cr.MaxIdle
	}

	if cr.IdleTimeout > 0 {
		pool.IdleTimeout = time.Second * time.Duration(cr.IdleTimeout)
	}

	return &RedisClient{
		pool: pool,
	}, err
}

//操作命令
func (r *RedisClient) Do(cmd string, args ...interface{}) (reply interface{}, err error) {

	//从连接池中获取redis连接
	conn := r.pool.Get()
	defer conn.Close()
	return conn.Do(cmd, args...)
}

// 获取连接
func (r *RedisClient) Get() redis.Conn {
	return r.pool.Get()
}

//Redis集群客户端
type RedisClusterClient struct {
	conn *rc.Cluster
}

func OpenRedisCluster(address []string) (*RedisClusterClient, error) {

	ops := &rc.Options{
		StartNodes: address,
	}

	ccl := ConfigInfo.GetRedisClusterConfig()
	keep_alive := ccl.KeepAlive
	alive_time := ccl.AliveTime
	conn_timeout := ccl.ConnTimeout
	read_timeout := ccl.ReadTimeout
	write_timeout := ccl.WriteTimeout

	if keep_alive > 0 {
		ops.KeepAlive = keep_alive
	}

	if alive_time > 0 {
		ops.AliveTime = time.Duration(alive_time) * time.Second
	}

	if conn_timeout > 0 {
		ops.ConnTimeout = time.Duration(conn_timeout) * time.Millisecond
	}

	if read_timeout > 0 {
		ops.ReadTimeout = time.Duration(read_timeout) * time.Millisecond
	}

	if write_timeout > 0 {
		ops.WriteTimeout = time.Duration(write_timeout) * time.Millisecond
	}

	cluster, err := rc.NewCluster(ops)
	return &RedisClusterClient{
		conn: cluster,
	}, err
}

func (r *RedisClusterClient) Do(cmd string, args ...interface{}) (reply interface{}, err error) {
	return r.conn.Do(cmd, args...)
}
