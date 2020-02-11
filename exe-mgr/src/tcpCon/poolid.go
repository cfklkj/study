package tcpClient

import (
	"../module/queue"
)

type ServerIdMgr struct {
	idPool *queue.LinkListNode
}

func NewServerIdMgr() *ServerIdMgr {
	ret := new(ServerIdMgr)
	ret.initIdPool()
	return ret
}

func (c *ServerIdMgr) initIdPool() {
	c.idPool = queue.NewLinkListNode()
	for i := 1; i < 50; i++ {
		c.idPool.Enque(i)
	}
}

//---
func (c *ServerIdMgr) popId() int {
	v, _ := c.idPool.Deque()
	return v.(int)
}

func (c *ServerIdMgr) pushId(id int) {
	c.idPool.Enque(id)
}
