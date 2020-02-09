package tcpClient

import (
	"../define"
)

func (c *WebTcp) BroadMsg(msg interface{}) {
	for _, k := range c.login.GetUsers() {
		//	if v.(string) != name {
		con := k.(*TcpClient)
		con.NoticeClient(define.Msg_pub, msg)
		//	}
	}
}

func (c *WebTcp) RouteMsg(to string, msg interface{}) {
	if con := c.login.GetCon(to); con != nil {
		cons := con.(*TcpClient)
		cons.NoticeClient(define.Msg_route, msg)
	}
}
