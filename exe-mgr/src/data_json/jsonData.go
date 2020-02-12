package data_json

import (
	conf "../module/util/config"
)

type JsonData struct {
	cf  *conf.Config
	key string
}

//初始化
func NewJsonData() *JsonData {
	ret := new(JsonData)
	ret.Init()
	return ret
}

//初始化配置
func (c *JsonData) Init() {
	c.cf = conf.NewConfig()
	c.key = "RunInfo"
}
func (c *JsonData) UpData(data []RunInfo) {
	c.cf.SetConfigInfo(c.key, &data)
}

//获取配置信息
func (c *JsonData) GetData(data *[]RunInfo) bool {
	return c.cf.GetConfigInfo(c.key, data)
}
