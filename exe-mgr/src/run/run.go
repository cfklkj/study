package run

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/axgle/mahonia"
)

type RunAdd func(name string, id int)
type RunDel func(name string, id int)
type Record func(name string, reader io.ReadCloser)

//执行命令--
func RunCmd(ctx context.Context, name, path string, callLog Record, callAdd RunAdd, callDel RunDel) {
	dir := filepath.Dir(path)
	cmd := exec.CommandContext(ctx, "cmd", "/c", path)
	cmdStdoutPipe, _ := cmd.StdoutPipe()
	cmdStderrPipe, _ := cmd.StderrPipe()
	defer cmdStdoutPipe.Close() // 保证关闭输出流
	defer cmdStderrPipe.Close() // 保证关闭输出流
	cmd.Dir = dir
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}
	go callLog(name, cmdStdoutPipe)
	go callLog(name, cmdStderrPipe)
	callAdd(name, cmd.Process.Pid)
	err := cmd.Wait()
	if err != nil {
		fmt.Println(err)
	}
	callDel(name, cmd.Process.Pid)
	fmt.Println("server-over", name)
}

/*参考
https://blog.csdn.net/zhuxinquan61/article/details/89716301
*/
//通过管道同步获取日志的函数
func (c *DoRunInfo) syncLog(name string, reader io.ReadCloser) {
	buf := make([]byte, 1024, 1024)
	for c.IsRun(name) {
		strNum, err := reader.Read(buf)
		if strNum > 0 {
			outputByte := buf[:strNum]
			//c.PrintToWeb(name, string(tls.DoZlibCompress(outputByte)))
			if !validUTF8(string(outputByte)) { //判断字符编码
				c.PrintToWeb(name, ConvertByte2String(outputByte, GB18030)) //控制台的字符编码 GB18030
			} else {
				c.PrintToWeb(name, string(outputByte))
			}
		}
		if err != nil {
			//读到结尾
			if err == io.EOF || strings.Contains(err.Error(), "file already closed") {
				err = nil
			}
		}
	}
}

/*
编码转换
import "golang.org/x/text/encoding/simplifiedchinese"
参考https://blog.csdn.net/jeffrey11223/article/details/79287010
*/
type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByte2String(src []byte, charset Charset) string {
	tagCoder := mahonia.NewDecoder(string(charset))
	_, cdata, _ := tagCoder.Translate(src, true)
	return string(cdata)
}
func ConvertToByte(src string, srcCode string, targetCode string) []byte {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(targetCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	return cdata
}

/*
判断字符编码
原文链接：https://blog.csdn.net/cxzzxc123456/article/details/83153945
*/
func validUTF8(src string) bool {
	buf := []byte(src)
	nBytes := 0

	for i := 0; i < len(buf); i++ {

		if nBytes == 0 {

			if (buf[i] & 0x80) != 0 { //与操作之后不为0，说明首位为1

				for (buf[i] & 0x80) != 0 {

					buf[i] <<= 1 //左移一位

					nBytes++ //记录字符共占几个字节

				}

				if nBytes < 2 || nBytes > 6 { //因为UTF8编码单字符最多不超过6个字节

					return false

				}

				nBytes-- //减掉首字节的一个计数

			}

		} else { //处理多字节字符

			if buf[i]&0xc0 != 0x80 { //判断多字节后面的字节是否是10开头

				return false

			}

			nBytes--

		}

	}
	return nBytes == 0
}
