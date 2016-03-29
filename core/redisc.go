package core

import (
	"github.com/chasex/redis-go-cluster"
	"time"
)

type RedisClusterConn struct {
	conn *redis.Cluster
}

func OpenRedisCluster(address []string) (*RedisClusterConn, error) {
	cluster, err := redis.NewCluster(
		&redis.Options{
			StartNodes:   address,
			ConnTimeout:  50 * time.Millisecond,
			ReadTimeout:  50 * time.Millisecond,
			WriteTimeout: 50 * time.Millisecond,
			KeepAlive:    16,
			AliveTime:    60 * time.Second,
		})
	return &RedisClusterConn{
		conn: cluster,
	}, err
}

func (r *RedisClusterConn) Do(cmd string, args ...interface{}) (reply interface{}, err error) {
	return r.conn.Do(cmd, args...)
}
