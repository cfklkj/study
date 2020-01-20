package main

import (
	"fmt"
	"net/http"

	tcpClient "./tcpCon"
	"golang.org/x/net/websocket"
)

type AcceptWebSoketFunc func(cn *websocket.Conn)

func RunWebSocketSvr(thisWebSocketUrl string, callback AcceptWebSoketFunc) {
	fmt.Println("启动websocket服务:", thisWebSocketUrl)

	http.Handle("/v1/tzj", websocket.Handler(callback))
	err := http.ListenAndServe(thisWebSocketUrl, nil)
	if err != nil {
		fmt.Println(err, "ListenAndServe")
	}
}

func main() {
	RunWebSocketSvr("localhost:10023", tcpClient.AcceptFunc)
}