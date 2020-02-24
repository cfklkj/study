package cmds

import "fmt"

type Check func(string)

type Cmds struct {
	status map[string]interface{}
	check  map[string]Check
	Run    func(name, shell string)
}

func NewCmds() *Cmds {
	ret := new(Cmds)
	ret.status = make(map[string]interface{})
	ret.check = make(map[string]Check)
	return ret
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
