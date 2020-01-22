package config

import (
	"fmt"

	conf "../../../util/config"
)

//-------config.jso  ------
type ConfigInfo struct {
	WebSocket string `json:""`
	SvrUrl    string
}

type Config struct {
	info ConfigInfo
	cf   *conf.Config
}

func makeInfo() ConfigInfo {
	var info ConfigInfo
	info.WebSocket = "localhost:10023"
	info.SvrUrl = "http://134.175.145.46:8025/im/info"
	return info
}

//初始化
func NewConfig() *Config {
	ret := new(Config)
	ret.cf = conf.NewConfig()
	ret.info = makeInfo()
	ret.Init()
	fmt.Println(ret.info)
	return ret
}

//初始化配置
func (c *Config) Init() {
	key := "localIm"
	if !c.cf.GetConfigInfo(key, &c.info) {
		c.cf.SetConfigInfo(key, &c.info)
	}
}

//获取配置信息
func (c *Config) GetConfigInfo() *ConfigInfo {
	return &c.info
}
