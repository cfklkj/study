package config

import (
	"fmt"

	conf "../util/config"
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
	info.WebSocket = ":10023"
	info.HttpIpPort = ":10024"
	info.UpUrl = ":10024/upload"
	info.DownUrl = ":10024/download"
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
	key := "localIm"
	if !c.cf.GetConfigInfo(key, &c.info) {
		c.cf.SetConfigInfo(key, &c.info)
	}
}

//获取配置信息
func (c *Config) GetConfigInfo() *ConfigInfo {
	return &c.info
}
