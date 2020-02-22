package tcpClient

import "fmt"

func (c *WebTcp) printMsg(act int, msg string) {
	//fmt.Println("printmsg", name, msg)
	//	act :=
	if c.client != nil {
		c.client.NoticeClient(act, msg)
	} else {
		fmt.Println("printmsg", msg)
	}
}
