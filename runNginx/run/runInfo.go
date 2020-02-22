package run

import (
	"context"
	"sync"
)

type DoRunInfo struct {
	ctx      context.Context
	cancel   context.CancelFunc
	mem      map[string]int
	lock     sync.Mutex
	PrintMsg func(name, msg string)
}

func NewDoRunInfo() *DoRunInfo {
	ret := new(DoRunInfo)
	ret.mem = make(map[string]int)
	ret.ctx, ret.cancel = context.WithCancel(context.Background())
	return ret
}
func (c *DoRunInfo) Close() {
	c.cancel()
}

func (c *DoRunInfo) Run(name, command string, args []string) {
	go c.RunCmd(c.ctx, name, command, args, c.syncLog, c.addRunInfo, c.delRunInfo)
}

func (c *DoRunInfo) CheckRun(name, cmd string, args []string) {
	go c.commandCheck(c.ctx, name, cmd, args)
}

func (c *DoRunInfo) Stop(name string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if pid := c.mem[name]; pid > 0 {
		c.mem[name] = 0
		//KillAll(pid)
	}
}

func (c *DoRunInfo) addRunInfo(name string, pid int) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.mem[name] = pid
}

func (c *DoRunInfo) delRunInfo(name string, pid int) {
	c.Stop(name)
}

func (c *DoRunInfo) IsRun(name string) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	if pid := c.mem[name]; pid > 0 {
		return true
	}
	return false
}
