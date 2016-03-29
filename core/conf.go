package core

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var ConfigInfo *Conf

//配置信息结构
type Conf struct {
	REDIS        *Redis        `yaml:"redis"`
	REDISCLUSTER *RedisCluster `yaml:"redis_cluster"`
	MYSQL        *Mysql        `yaml:"mysql"`
	IsCluster    bool          `yaml:"iscluster"`
}

type Redis struct {
	Addrs       string `yaml:"addrs"`
	MaxIdle     int    `yaml:"max_idle"`
	IdleTimeout int    `yaml:"idle_timeout"`
}

type RedisCluster struct {
	Addrs        string `yaml:"addrs"`
	KeepAlive    int    `yaml:"keep_alive"`
	AliveTime    int    `yaml:"alive_time"`
	ConnTimeout  int    `yaml:"conn_timeout"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func (c *Conf) GetRedisConfig() *Redis {
	return c.REDIS
}

func (c *Conf) GetRedisClusterConfig() *RedisCluster {
	return c.REDISCLUSTER
}

func (c *Conf) GetMysqlConfig() *Mysql {
	return c.MYSQL
}

func SetConfig(conf *Conf) {
	ConfigInfo = conf
}

func GetConfig() *Conf {
	return ConfigInfo
}

//解析配置文件
func ParseConfigFile(fileName string) (*Conf, error) {
	var cfg Conf
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
