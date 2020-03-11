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
		FileServerDir: "./",
		IndexHtml:     "index.html",
	}
	svrtcp := svr_tcp.NewSvrTcp()
	svrfile := svr_file.SvrFile{
		PatternFileServer: "/download/",
		PatternUpfile:     "/upload",
		PatternLog:        "/log",
		PatternFriend:     "/friend",
		FileServerDir:     "d:/www",
		FormKey_up:        "uploadfile",
	}

	host := "127.0.0.1:801"
	if len(os.Args) > 1 {
		host = os.Args[1]
	}
	if len(os.Args) > 2 {
		svrweb.FileServerDir = os.Args[2]
	}
	if len(os.Args) > 3 {
		svrfile.FileServerDir = os.Args[3]
	}
	fmt.Println("eg:\n [ip:port] [webLocalPath] [fileLocalPath]\nstart:" + host)

	svrtcp.Listen()
	svrweb.Listen()
	svrfile.Listen()
	http.ListenAndServe(host, nil)
}
