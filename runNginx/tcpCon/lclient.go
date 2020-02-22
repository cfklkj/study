package tcpClient

import (
	"encoding/json"
	"fmt"

	"../config"
	define "../define"
	"../nginx"
	"golang.org/x/net/websocket"
)

type WebTcp struct {
	conf      *config.ConfigInfo
	client    *TcpClient
	nginx     *nginx.Nginx
	loginUser string
}

func NewWebTcp() *WebTcp {
	ret := new(WebTcp)
	ret.nginx = nginx.NewNginx()
	ret.nginx.PrintMsg = ret.printMsg
	ret.nginx.IsInstallunzip()
	ret.nginx.IsInstallcertbot()
	return ret
}
func (c *WebTcp) getLoginUser() string {
	return c.loginUser
}
func (c *WebTcp) DataToStruct(data interface{}, out interface{}) {
	rst, _ := json.Marshal(data)
	json.Unmarshal(rst, out)
}
func (c *WebTcp) AcceptFunc(conn *websocket.Conn) {
	c.client = NewTcpClient(conn)
	for {
		if !c.client.readMsg(false) {
			break
		}
		var info define.MsgInfo
		json.Unmarshal(c.client.allBytes, &info)
		fmt.Println("info", info)
		// if c.loginUser == "" && define.Msg_login != info.Act {
		// 	c.client.NoticeClient(define.Msg_lastErr, "请先登入！")
		// 	continue
		// }
		switch info.Act {
		case define.Msg_login:
			var data define.LoginInfo
			c.DataToStruct(info.Data, &data)
			// if c.loginUser == data.Username { //重复登入
			// 	c.client.NoticeClient(define.Msg_login, true)
			// }
			c.loginUser = data.Name
			continue
		case define.Msg_nginx:
			c.nginx.IsInstallNginx()
			continue
		case define.Msg_nginxRestart:
			c.nginx.ReStartNginx()
			continue
		case define.Msg_down:
			var data define.DownInfo
			c.DataToStruct(info.Data, &data)
			c.nginx.DownFile(data.DownUrl, data.KeepPath, data.FileType == "zip")
			continue
		case define.Msg_certNew:
			var data define.CertInfo
			c.DataToStruct(info.Data, &data)
			c.nginx.MakeCert(data.Domain)
			continue
		case define.Msg_certRenew:
			c.nginx.RenewCert()
			continue
		default:
			c.client.NoticeClient(info.Act, info)
		}
	}
}
