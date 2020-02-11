package data_json

import "fmt"

type RunInfo struct {
	Name, Path string
}
type RunInfos struct {
	datas    []RunInfo
	jsonData *JsonData
}

type RunCall func(name, path string)
type StopCall func(name string)

func NewRunInfos() *RunInfos {
	ret := new(RunInfos)
	ret.jsonData = NewJsonData()
	ret.jsonData.GetData(&ret.datas)
	return ret
}

func (c *RunInfos) AddRunInfo(name, path string) bool {
	var info RunInfo
	info.Name = name
	info.Path = path
	for _, k := range c.datas {
		if k.Name == name {
			return false
		}
	}
	c.datas = append(c.datas, info)
	c.jsonData.UpData(c.datas)
	fmt.Println(c.datas)
	return true
}
func (c *RunInfos) AltRunInfo(oldName, name, path string) bool {
	if c.GetRunInfo(name) != nil {
		return false
	}
	var info RunInfo
	info.Name = name
	info.Path = path
	for v, k := range c.datas {
		if k.Name == oldName {
			k.Name = name
			k.Path = path
			c.datas[v] = k
			return true
		}
	}
	return false
}

func (c *RunInfos) DelRunInfo(name string) bool {

	for index, k := range c.datas {
		if k.Name == name {
			c.datas = append(c.datas[:index], c.datas[index+1:]...)
			c.jsonData.UpData(c.datas)
			return true
		}
	}
	return false
}
func (c *RunInfos) GetRunInfo(name string) interface{} {

	for _, k := range c.datas {
		if k.Name == name {
			return k
		}
	}
	return nil
}
func (c *RunInfos) GetRunInfos() (rst []string) {

	for _, k := range c.datas {
		rst = append(rst, k.Name)
	}
	return rst
}
func (c *RunInfos) RunSingle(name string, callBack RunCall) bool {

	for _, k := range c.datas {
		if k.Name == name {
			callBack(k.Name, k.Path)
			return true
		}
	}
	return false
}

func (c *RunInfos) RunAll(callBack RunCall) {

	for _, k := range c.datas {
		callBack(k.Name, k.Path)
	}
}

func (c *RunInfos) StopSingle(name string, callBack StopCall) bool {
	for _, k := range c.datas {
		if k.Name == name {
			callBack(name)
			return true
		}
	}
	return false
}

func (c *RunInfos) StopAll(callBack StopCall) {

	for _, k := range c.datas {
		callBack(k.Path)
	}
}
