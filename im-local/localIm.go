package main

import (
	"net/http"

	httpClient "./httpCon"
	tcpClient "./tcpCon"
	"golang.org/x/net/websocket"
)

func RunHttpSvr() {
	handle := httpClient.NewHttp()
	webTcp := tcpClient.NewWebTcp()
	http.Handle("/wss", websocket.Handler(webTcp.AcceptFunc))
	handle.Listen()
}
func main() {
	RunHttpSvr()
}
