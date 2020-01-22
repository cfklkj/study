package tcpClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"../config"
	"../define"
	"golang.org/x/net/websocket"
)

var clientCount = 0

func AcceptFunc(conn *websocket.Conn) {
	clientCount += 1
	for {
		client := NewTcpClient(conn)
		if !client.readMsg(false) {
			break
		}
		var info define.MsgInfo
		json.Unmarshal(client.allBytes, &info)
		switch info.Act {
		case define.Msg_login:
			cfg := config.NewConfig().GetConfigInfo()
			info.Data = Prox(cfg.SvrUrl, "application/x-www-form-urlencoded;", info.Data)
		case define.Msg_set:
			writeLine(info.Sender, info.Data)
		case define.Msg_get:
			info.Data = readLine(info.Sender, info.Index)
		case define.Msg_del:
		case define.Msg_getLenth:
			info.Index = countLine(info.Sender)
			fmt.Println("have--", info.Sender, info.Index)
		case define.Msg_fileUp:
		default:
		}
		client.NoticeClient(info.Act, info)
	}
	clientCount -= 1
	// if clientCount == 0 {
	// 	os.Exit(0)
	// }
	fmt.Println("over")
}

func readLine(user string, index int) string {
	path := "./" + user + ".log"
	return ReadLine(path, index)
}

func writeLine(user, msg string) int {
	path := "./" + user + ".log"
	return WriteLine(path, msg)
}
func countLine(user string) int {
	path := "./" + user + ".log"
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
		return "Err_postDo"
	}
	//结果返回
	datas, _ := ioutil.ReadAll(response.Body)
	return string(datas)
}
