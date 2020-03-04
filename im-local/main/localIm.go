package main

import (
	"fmt"
	"net/http"
	"os"

	"../svr_module/svr_file"
	"../svr_module/svr_tcp"
	"../svr_module/svr_web"
)

func main() {
	svrweb := svr_web.SvrWeb{
		PatternCss:    "/css/",
		PatternJs:     "/js/",
		PatternImg:    "/image/",
		FileServerDir: "./web",
		IndexHtml:     "index.html",
	}
	svrtcp := svr_tcp.NewSvrTcp()
	svrfile := svr_file.SvrFile{
		PatternFileServer: "/download",
		PatternUpfile:     "/upload",
		PatternLog:        "/log",
		PatternFriend:     "/friend",
		FileServerDir:     "d:/www",
		FormKey_up:        "uploadfile",
	}
	svrtcp.Listen()
	svrweb.Listen()
	svrfile.Listen()
	host := "127.0.0.1:801"
	if len(os.Args) > 1 {
		host = os.Args[1]
	}
	fmt.Println("eg:\n [ip:port]\nstart:" + host)
	http.ListenAndServe(host, nil)
}
