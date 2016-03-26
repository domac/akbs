package core

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var ConfigInfo *Conf

//配置信息结构
type Conf struct {
	REDIS *Redis `yaml:"redis"`
	MYSQL *Mysql `yaml:"mysql"`
}

type Redis struct {
	Addrs     string `yaml:"addrs"`
	IsCluster bool   `yaml:"iscluster"`
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
