package tzj_kefu

import (
	//"fmt"

	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"../config"
)

type Http struct {
	conf *config.ConfigInfo
}

func NewHttp() *Http {
	ret := new(Http)
	ret.conf = config.NewConfig().GetConfigInfo()
	return ret
}

//开启服务器
func (c *Http) Listen() {
	// 设置服务文件
	http.Handle(CSS_CLIENT_PATH, http.FileServer(http.Dir(CSS_SVR_PATH)))
	http.Handle(DART_CLIENT_PATH, http.FileServer(http.Dir(DART_SVR_PATH)))
	http.Handle(IMAGE_CLIENT_PATH, http.FileServer(http.Dir(IMAGE_SVR_PATH)))
	http.HandleFunc("/prox", c.Prox)
	http.HandleFunc("/", c.HomePage)

	// 网址与处理逻辑对应起来
	//绑定socket方法
	c.setNotepadHandleFunc()
	http.ListenAndServe(c.conf.HttpIpPort, nil) //设置监听的端口
}

func (c *Http) Prox(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	body, err := ioutil.ReadAll(req.Body)
	var info ProxInfo
	json.Unmarshal(body, &info)
	fmt.Println("req.Body", string(body), info)
	url := info.Url
	data := info.Data
	body_type := info.ContentType // "application/json;charset=utf-8"
	if err != nil || data == "" || url == "" {
		c.sendBack(w, Err_parame, "")
		return
	}
	fmt.Println("data", url, data)
	reqest, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		c.sendBack(w, Err_post, err.Error())
		return
	}
	client := &http.Client{}
	reqest.Header.Set("Content-Type", body_type)
	//	reqest.Header.Set("Authorization", token)
	//处理返回结果
	response, err2 := client.Do(reqest)
	if err2 != nil {
		c.sendBack(w, Err_postDo, err)
		return
	}
	//结果返回
	datas, _ := ioutil.ReadAll(response.Body)
	c.sendBack(w, response.StatusCode, string(datas))
	// //将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
	//tdout := os.Stdout
	//_, err = io.Copy(stdout, response.Body)
	//   if response.StatusCode != 200 {
	//fmt.Println("err", "post result", response.StatusCode)
	// // }
	// //log.PrintErr("postMsg", body)
	// // //返回的状态码
	// status := response.StatusCode
}

func (c *Http) HomePage(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" && req.Method != "get" {
		c.sendBack(w, Err_MethodGet, "")
		fmt.Println("HomePage", "err", req.URL.Path)
		return
	}
	path := req.URL.Path
	if req.URL.Path != "/" {
		c.sendBack(w, Err_req, "")
		return
	}
	path = DART_SVR_PATH + c.conf.DefaultHtml
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

//校验
func (c *Http) headCheck(w http.ResponseWriter, req *http.Request) (string, []byte) {
	if req.Method != "POST" && req.Method != "post" {
		c.sendBack(w, Err_MethodPost, "")
		return "", nil
	}
	// req.ParseForm()
	// tokenStr := req.Form.Get("proid")
	// proid, err1 := c.getTokensData(TokenKey, tokenStr)
	// if err1 != Err_null {
	// 	c.sendBack(w, err1, len(proid))
	// 	return "", nil
	// }
	// body, err := ioutil.ReadAll(req.Body)
	// if err != nil || len(proid) < 3 || len(proid) > 64 {
	// 	c.sendBack(w, Err_Ummarshal, len(proid))
	// 	return "", nil
	// }
	// fmt.Println("headCheck", string(body))
	// return proid, body
	return "", nil
}

//返回消息
func (c *Http) sendBack(w http.ResponseWriter, code int, data interface{}) {
	var rst S2CBody
	rst.Code = code
	rst.CodeMsg = c.getCodeStr(code)
	rst.Data = data
	dataStr, _ := json.Marshal(rst)
	fmt.Println("sendBack", string(dataStr))
	io.WriteString(w, string(dataStr))
}
