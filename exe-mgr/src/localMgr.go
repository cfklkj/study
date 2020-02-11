package main

import (
	"fmt"
	"net/http"
	"os"

	"./config"
	httpClient "./httpCon"
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
		os.Exit(0)
	}
}
func RunHttpSvr() {
	handle := httpClient.NewHttp()
	handle.Listen()
}
func main() {
	cfg := config.NewConfig().GetConfigInfo()
	webTcp := tcpClient.NewWebTcp()
	go RunWebSocketSvr(cfg.WebSocket, webTcp.AcceptFunc)
	RunHttpSvr()
}
