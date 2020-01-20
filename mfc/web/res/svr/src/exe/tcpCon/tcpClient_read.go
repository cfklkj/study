package tcpClient

import (
	"net"
)

const (
	g_FirstHdTime  = 5    //秒  第一次等待
	g_hdTimeOut    = 3    ////3 分钟无通信
	g_readBuffSize = 1024 //每次读取的字节数
	g_maxLenth     = 4096 //每条消息最大长度
)

//读取信息
func (c *TcpClient) readMsg(isFisrt bool) bool {
	c.allBytesNum = 0
	c.allBytes = nil
	for {
		if c.nextBytes != nil {
			if index := c.findEnd(c.nextBytes); index > -1 { //找末尾 \n
				c.allBytesNum = index
				c.allBytes = append(c.allBytes, c.nextBytes[:index]...)
				if len(c.nextBytes) > index+1 {
					c.nextBytes = c.nextBytes[index+1:]
				} else {
					c.nextBytes = nil
				}
				return true
			}
			c.allBytes = c.nextBytes
			c.nextBytes = nil
			c.allBytesNum = len(c.nextBytes)
			if c.readEnd {
				c.readEnd = false
				if c.allBytesNum > 0 {
					return true
				}
			}
		}
		//读取连接内存
		buf := make([]byte, g_readBuffSize)
		readBytesNum, err := c.ConRead(buf)
		if err != nil { //读取异常
			errs, ok := err.(net.Error)
			if !isFisrt {
				if !ok || !errs.Temporary() {
					c.StopTcp() //端断开连接 -- 连接错误
					break
				}
				continue
			} else {
				c.StopTcp() //端断开连接 -- 超时未操作
			}
			break
		}
		if readBytesNum != 0 {
			if index := c.findEnd(buf); index > -1 { //找末尾 \n
				c.allBytesNum += index
				c.allBytes = append(c.allBytes, buf[:index]...)
				if index != readBytesNum {
					c.nextBytes = buf[index+1 : readBytesNum]
				}
				if readBytesNum != g_readBuffSize { //接收的长度没满，说明已经读到最后了
					c.readEnd = true
				}
				return true
			}
			c.allBytesNum += readBytesNum
			c.allBytes = append(c.allBytes, buf[:readBytesNum]...)
		}
		if readBytesNum != g_readBuffSize { //接收的长度没满，说明已经读到最后了
			return true
		}
	}
	return false
}

func (c *TcpClient) findEnd(buff []byte) int {
	for index, data := range buff {
		if data == '\n' {
			return index
		}
	}
	return -1
}
