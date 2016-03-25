package core

import (
	"errors"
	"gopkg.in/redis.v3"
)

type MyRedisClient struct {
	redisClient   *redis.Client
	clusterClient *redis.ClusterClient
	isCluster     bool
}

func NewSimpleRedisClient() *MyRedisClient {
	return NewMyRedisClient(false)
}

//构建Redis客户端
func NewMyRedisClient(isCluster bool) *MyRedisClient {
	var redisClient *redis.Client
	var clusterClient *redis.ClusterClient

	if !isCluster {
		redisClient = redis.NewClient(&redis.Options{
			Addr: "192.168.139.139:6699",
		})
	} else {
		clusterClient = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs: []string{"192.168.139.139:6699"},
		})
	}
	return &MyRedisClient{
		redisClient:   redisClient,
		clusterClient: clusterClient,
		isCluster:     isCluster}
}

func (r *MyRedisClient) GetConfigConn() (interface{}, error) {
	if r.isCluster {
		return r.GetClusterConn()
	} else {
		return r.GetConn()
	}
}

func (r *MyRedisClient) GetConn() (*redis.Client, error) {
	if r.isCluster {
		return nil, errors.New("could not get conn")
	}

	err := r.redisClient.Ping().Err()
	if err != nil {
		return nil, err
	}
	return r.redisClient, nil
}

func (r *MyRedisClient) GetClusterConn() (*redis.ClusterClient, error) {
	if !r.isCluster {
		return nil, errors.New("could not get cluster conn")
	}

	err := r.clusterClient.Ping().Err()
	if err != nil {
		return nil, err
	}
	return r.clusterClient, nil
}
