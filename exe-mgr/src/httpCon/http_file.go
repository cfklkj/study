package httpCon

import (
	"io"
	"math/rand"
	"net/http"
	"os"

	"../define"
)

func (c *Http) randName(lenth int, tail string) string {
	rand.Seed(int64(rand.Intn(10000)))
	name := make([]byte, lenth)
	for i := 0; i < lenth; i++ {
		isUp := rand.Intn(2)
		index := rand.Intn(25)
		if isUp < 1 {
			name[i] = byte('a' + index)
		} else {
			name[i] = byte('A' + index)
		}
	}
	return string(name) + tail
}

func (c *Http) upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		c.sendBack(w, define.Err_MethodGet, "")
		return
	} else {
		//r.ParseMultipartForm(32 << 20)
		fileName := c.randName(9, ".data")
		file, _, err := r.FormFile(define.FILE_upName)
		if err != nil {
			c.sendBack(w, define.Err_post, "")
			return
		}
		defer file.Close()
		f, err := os.OpenFile(define.FILE_SVR_PATH+define.FILE_PATH+fileName, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			c.sendBack(w, define.Err_open, fileName)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		c.sendBack(w, define.Err_null, fileName)
	}
}
func (c *Http) listenFile() {
	http.Handle(define.FILE_PATH, http.FileServer(http.Dir(define.FILE_SVR_PATH)))
	http.HandleFunc("/upload", c.upload) //设置访问的路由
}
