package tcpClient

import (
	"encoding/json"
	"fmt"

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
		case define.Msg_set:
			writeLine(info.Sender, info.Data)
		case define.Msg_get:
			info.Data = readLine(info.Sender, info.Index)
		case define.Msg_del:
		case define.Msg_getLenth:
			info.Index = countLine(info.Sender)
			fmt.Println("have--", info.Sender, info.Index)
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
