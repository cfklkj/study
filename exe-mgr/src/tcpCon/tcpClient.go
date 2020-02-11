package tcpClient

import (
	"encoding/json"
	"net"

	"tzj.com/svr_http/src/define"
)

//用户信息
type ClientInfo struct {
	AppIndex, UserIndex int
}

//客户端
type TcpClient struct {
	conn        net.Conn //net.Conn  *websocket.Conn  与  net.Conn  存在继承关系？
	userInfo    ClientInfo
	stopCode    int            //停止原因
	allBytes    []byte         //接收的tcpbuff
	nextBytes   []byte         //额外数据
	s2cBody     define.S2CBody //返回客户端的消息体
	upTimeCount int            //更新心跳检测的次数
	allBytesNum int            //数据大小
	msgChan     chan []byte
	readEnd     bool
}

func NewTcpClient(con net.Conn) *TcpClient {
	ret := new(TcpClient)
	ret.conn = con
	ret.upTimeCount = 0
	ret.msgChan = make(chan []byte)
	return ret
}

//con act-----
func (c *TcpClient) ConClose() {
	c.conn.Close()
}
func (c *TcpClient) ConWrite(msg []byte) (int, error) {
	return c.conn.Write(msg)
}
func (c *TcpClient) ConRead(buff []byte) (int, error) {
	return c.conn.Read(buff)
}

//通知客户端服务器关闭连接
func (c *TcpClient) NoticeClient(opt int, msg interface{}) {
	c.s2cBody.Code = opt
	c.s2cBody.Data = msg
	rst, _ := json.Marshal(c.s2cBody)
	c.WriteToClient(rst)
}

//异步消息
func (c *TcpClient) WriteToClient(msg []byte) bool {
	_, err := c.ConWrite(msg)
	if err != nil {
		return false
	}
	return true
}

func (c *TcpClient) StopTcp() {
	if c.stopCode != 0 { //不重复 处理
		return
	}
	c.ConClose()
}
