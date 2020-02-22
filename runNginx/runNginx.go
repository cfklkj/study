package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"./config"
	"./define"
	tcpClient "./tcpCon"
	"golang.org/x/net/websocket"
)

func main() {
	runSvr()
}

func runSvr() {
	conf := config.NewConfig().GetConfigInfo()
	webTcp := tcpClient.NewWebTcp()
	http.Handle("/wss", websocket.Handler(webTcp.AcceptFunc))
	fmt.Println("服务地址:", conf.WebSocket)
	http.Handle(define.CSS_CLIENT_PATH, http.FileServer(http.Dir(define.CSS_SVR_PATH)))
	http.Handle(define.DART_CLIENT_PATH, http.FileServer(http.Dir(define.DART_SVR_PATH)))
	http.Handle(define.IMAGE_CLIENT_PATH, http.FileServer(http.Dir(define.IMAGE_SVR_PATH)))

	http.HandleFunc("/", HomePage)
	if err := http.ListenAndServe(conf.WebSocket, nil); err != nil {
		fmt.Println("err", err)
	}
}

func HomePage(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" && req.Method != "get" {
		fmt.Println("HomePage", "err", req.URL.Path)
		return
	}
	path := req.URL.Path
	if req.URL.Path != "/" {
		return
	}
	path = define.DART_SVR_PATH + "index.html"
	t, err := template.ParseFiles(path)
	if err != nil {
		fmt.Println("sdd", err)
		return
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err = w.Write(buf.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
}

func sendBack(w http.ResponseWriter, code int, data interface{}) {
	var rst define.S2CBody_http
	rst.Code = code
	rst.Data = data
	dataStr, _ := json.Marshal(rst)
	io.WriteString(w, string(dataStr))
}
