package cmds

import "fmt"

type Check func(string)

type Cmds struct {
	status map[string]interface{}
	check  map[string]Check
	result map[string]interface{}
	Run    func(name, shell string) bool
}

func NewCmds() *Cmds {
	ret := new(Cmds)
	ret.status = make(map[string]interface{})
	ret.check = make(map[string]Check)
	ret.result = make(map[string]interface{})
	return ret
}

func (c *Cmds) setResult(name string, value interface{}) {
	c.result[name] = value
}
func (c *Cmds) getResult(name string) interface{} {
	return c.result[name]
}

func (c *Cmds) BindCheck(name string, callBack Check) {
	if ok := c.check[name]; ok != nil {
		return
	}
	c.check[name] = callBack
}

func (c *Cmds) PrintMsg(name, msg string) {
	if call := c.check[name]; call != nil {
		call(msg)
	} else {
		fmt.Println(name, msg)
	}
}
