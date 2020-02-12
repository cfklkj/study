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
	"../data_json"
	define "../define"
	"../run"
	"golang.org/x/net/websocket"
)

type WebTcp struct {
	conf     *config.ConfigInfo
	dataJson *data_json.RunInfos
	run      *run.DoRunInfo
	clients  map[int]*TcpClient
	buildId  *ServerIdMgr
}

func NewWebTcp() *WebTcp {
	ret := new(WebTcp)
	ret.dataJson = data_json.NewRunInfos()
	ret.conf = config.NewConfig().GetConfigInfo()
	ret.run = run.NewDoRunInfo()
	ret.run.PrintToWeb = ret.printToWeb
	ret.clients = make(map[int]*TcpClient)
	ret.buildId = NewServerIdMgr()
	return ret
}

func (c *WebTcp) printToWeb(name, msg string) {
	var info define.MsgInfo
	info.Act = define.Msg_print
	for _, con := range c.clients {
		if con == nil {
			continue
		}
		info.ConversationId = name
		info.Data = msg
		con.NoticeClient(info.Act, info)
	}
}

func (c *WebTcp) AcceptFunc(conn *websocket.Conn) {
	username := ""
	client := NewTcpClient(conn)
	id := c.buildId.popId()
	c.clients[id] = client
	for {
		if !client.readMsg(false) {
			break
		}
		var info define.MsgInfo
		json.Unmarshal(client.allBytes, &info)
		switch info.Act {
		case define.Msg_Login:
			username = info.ConversationId
			continue
		case define.Msg_set:
			writeLine(username, info.ConversationId, info.Data.(string))
			continue
		case define.Msg_get:
			info.Data = readLine(username, info.ConversationId, info.Index)
		case define.Msg_del:
		case define.Msg_getLenth:
			info.Index = countLine(username, info.ConversationId)
		case define.Msg_run:
			data, _ := json.Marshal(info.Data)
			var runInfo define.Run
			json.Unmarshal(data, &runInfo)
			switch runInfo.Act {
			case define.Act_start:
				if !c.run.IsRun(runInfo.Name) {
					runInfo.Data = c.dataJson.RunSingle(runInfo.Name, c.run.Run)
				} else {
					runInfo.Data = true
				}
			case define.Act_stop:
				runInfo.Data = c.dataJson.StopSingle(runInfo.Name, c.run.Stop)
			case define.Act_startAll:
				//runInfo.Data = c.dataJson.RunAll(c.run.Run)
			case define.Act_stopAll:
				//runInfo.Data = c.dataJson.StopAll(c.run.Stop)
			case define.Act_run_statu:
				runInfo.Data = c.run.IsRun(runInfo.Name)
			}
			info.Data = runInfo
		case define.Msg_setRunInfo:
			data, _ := json.Marshal(info.Data)
			var runInfo define.RunInfo
			json.Unmarshal(data, &runInfo)
			switch runInfo.Act {
			case define.Act_add:
				if c.dataJson.AddRunInfo(runInfo.Name, runInfo.Path) {
					runInfo.Data = runInfo.Name
				} else {
					runInfo.Data = ""
				}
			case define.Act_alt:
				runInfo.Data = c.dataJson.AltRunInfo(runInfo.OldName, runInfo.Name, runInfo.Path)
				fmt.Println(runInfo.OldName, runInfo.Name, runInfo.Path, runInfo.Data)
			case define.Act_del:
				c.dataJson.StopSingle(runInfo.Name, c.run.Stop)
				runInfo.Data = c.dataJson.DelRunInfo(runInfo.Name)
			}
			info.Data = runInfo
		case define.Msg_getRunInfo:
			data, _ := json.Marshal(info.Data)
			var runInfo define.GetRunInfo
			json.Unmarshal(data, &runInfo)
			if runInfo.Act != define.Act_single {
				var info []define.RunStatu
				for _, name := range c.dataJson.GetRunInfos() {
					var statu define.RunStatu
					statu.Name = name
					statu.IsRun = c.run.IsRun(statu.Name)
					info = append(info, statu)
				}
				runInfo.Data = info
			} else {
				if k := c.dataJson.GetRunInfo(runInfo.Name); k != nil {
					var statu define.RunStatu
					statu.Name = k.(data_json.RunInfo).Name
					statu.Path = k.(data_json.RunInfo).Path
					statu.IsRun = c.run.IsRun(statu.Name)
					runInfo.Data = statu
				} else {
					runInfo.Data = nil
				}
			}
			info.Data = runInfo
		default:
		}
		client.NoticeClient(info.Act, info)
	}
	c.clients[id] = nil
	c.buildId.pushId(id)
	fmt.Println("web-client-out", username)
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
