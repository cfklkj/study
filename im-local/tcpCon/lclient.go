package tcpClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"../act"
	"../config"
	define "../define"
	"golang.org/x/net/websocket"
)

type WebTcp struct {
	login *act.Login
	conf  *config.ConfigInfo
}

func NewWebTcp() *WebTcp {
	ret := new(WebTcp)
	ret.login = act.NewLogin()
	ret.conf = config.NewConfig().GetConfigInfo()
	return ret
}

func (c *WebTcp) AcceptFunc(conn *websocket.Conn) {
	username := ""
	client := NewTcpClient(conn)
	for {
		if !client.readMsg(false) {
			break
		}
		var info define.MsgInfo
		json.Unmarshal(client.allBytes, &info)
		fmt.Println("func", info)
		switch info.Act {
		case define.Msg_login:
			var data define.LoginMsg
			data.Code = define.Err_null
			data.Data = c.login.GetUserId(info.Data.(string))
			if v := c.login.GetCon(data.Data); v != nil {
				data.Code = define.Err_logined
				info.Data, _ = json.Marshal(data)
				client.NoticeClient(info.Act, info)
				return
			}
			data.UpUrl = c.conf.UpUrl
			data.DownUrl = c.conf.DownUrl
			username = data.Data
			info.Data, _ = json.Marshal(data)
			c.login.SetCon(username, client)
		case define.Msg_pub:
			c.BroadMsg(info.Data)
			continue
		case define.Msg_route:
			to := info.ConversationId
			info.ConversationId = username
			c.RouteMsg(to, info)
			continue
		case define.Msg_set:
			writeLine(username, info.ConversationId, info.Data.(string))
			continue
		case define.Msg_get:
			info.Data = readLine(username, info.ConversationId, info.Index)
		case define.Msg_del:
		case define.Msg_getLenth:
			info.Index = countLine(username, info.ConversationId)
			fmt.Println("have--", info.ConversationId, info.Index)
		case define.Msg_fileUp:
		default:
		}
		client.NoticeClient(info.Act, info)
	}
	c.login.DelCon(username)
	fmt.Println("over", username)
}

func readLine(user, to string, index int) string {
	path := "./document/" + user + "_" + to + ".log"
	return ReadLine(path, index)
}

func writeLine(user, to, msg string) int {
	path := "./document/" + user + "_" + to + ".log"
	return WriteLine(path, msg)
}
func countLine(user, to string) int {
	path := "./document/" + user + "_" + to + ".log"
	return CountsLine(path)
}

func Prox(url, body_type, data string) string {
	if data == "" || url == "" {
		return "Err_parame"
	}

	reqest, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return "Err_post"
	}
	reqest.Header.Set("Content-Type", body_type) // "application/json;charset=utf-8"
	//处理返回结果
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {

				conn, err := net.DialTimeout(netw, addr, time.Second*2) //设置建立连接超时

				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * 2)) //设置发送接受数据超时

				return conn, nil
			},

			ResponseHeaderTimeout: time.Second * 2,
		},
	}

	response, err2 := client.Do(reqest)
	if err2 != nil {
		fmt.Println("Err_postDo", err2)
		return "Err_postDo"
	}
	//结果返回
	datas, _ := ioutil.ReadAll(response.Body)
	return string(datas)
}
