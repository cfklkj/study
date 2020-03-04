package Config

/*
配置信息
加载配置LoadConfig
获取配置结构体 GetConfigInfo

*/
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	node string
	Info interface{}
	path string
}

//初始化
func NewConfig() *Config {
	ret := new(Config)
	ret.path = "./svrConfig.json"
	return ret
}

func (c *Config) SetConfigPath(path string) {
	if path == "" {
		return
	}
	c.path = path
	c.Info = nil
}

//配置文件路径
func (c *Config) getConfigPath() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	return c.path
}

//获取配置信息
func (c *Config) GetConfigInfo(node string, res interface{}) bool {
	if c.node != node {
		c.Info = c.getValue(c.getConfigPath(), node)
		c.node = node
	}
	if c.Info == nil {
		return false
	}
	//解析信息到对象
	data, err := json.Marshal(c.Info)
	if err != nil {
		fmt.Println("GetConfigInfo", "err", err)
		return false
	}
	err = json.Unmarshal(data, res)
	return err == nil
}

//设置配置信息
func (c *Config) SetConfigInfo(key string, value interface{}) {
	c.Info = value
	c.setValue(c.getConfigPath(), key, value)
}

//创建配置文件
func (c *Config) createConfigFile() error {
	// 创建文件
	filePtr, err := os.Create(c.getConfigPath())
	if err != nil {
		return err
	}
	defer filePtr.Close()
	filePtr.Write([]byte("{}"))
	return err
}

//添加值
func (c *Config) setValue(jsonFile string, key string, value interface{}) bool {
	byteValue, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		fmt.Println("setValue", "err", err)
		return false
	}
	//读取原结构
	var result map[string]interface{}
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		fmt.Println("setValue", "err2", err)
		return false
	}
	//添加值
	result[key] = value
	byteValue, err = json.Marshal(result)
	if err != nil {
		fmt.Println("setValue", "err3", err)
		return false
	}
	fmt.Println("setValue", result, string(byteValue))
	err = ioutil.WriteFile(jsonFile, byteValue, 0644)
	return err == nil
}
func (c *Config) printErrMsg() {
	fmt.Println("please change the svrConfig.json info")
	os.Exit(1)
}
func (c *Config) getValue(jsonFile string, key string) interface{} {
	// Read json buffer from jsonFile
	byteValue, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		c.createConfigFile()
		return nil
	}
	//读取原结构
	var result map[string]interface{}
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return nil
	}
	if len(result) < 1 {
		return nil
	}
	if value, ok := result[key]; ok {
		return value
	}
	return nil
}

//-----test
func main() {
	// ret := NewConfig()
	// var tmp define.SvrConfig
	// t := ret.GetConfigInfo("svr_discover1", &tmp)
	// if !t {
	// 	ret.SetConfigInfo("svr_discover1", tmp)
	// }
	// fmt.Println(t, tmp)
}
