package config

import (
	"fmt"

	conf "../module/util/config"
)

//-------config.jso  ------
type ConfigInfo struct {
	WebSocket   string `json:""`
	HttpIpPort  string
	UpUrl       string
	DownUrl     string
	DefaultHtml string
}

type Config struct {
	info ConfigInfo
	cf   *conf.Config
}

func makeInfo() ConfigInfo {
	var info ConfigInfo
	info.WebSocket = "127.0.0.1:10033"
	info.HttpIpPort = "127.0.0.1:10034"
	info.UpUrl = ":10034/upload"
	info.DownUrl = ":10034/download"
	info.DefaultHtml = "index.html"
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
	key := "localMgr"
	if !c.cf.GetConfigInfo(key, &c.info) {
		c.cf.SetConfigInfo(key, &c.info)
	}
}

//获取配置信息
func (c *Config) GetConfigInfo() *ConfigInfo {
	return &c.info
}
