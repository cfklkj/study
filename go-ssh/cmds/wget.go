package cmds

import (
	"fmt"
	"strings"
)

//wget
func (c *Cmds) Cmd_downFile(url, keepPath string, isZip bool, zipType int) bool {
	c.BindCheck(str_wget, c.checkDownfile)
	str := "wget " + url + " -O " + keepPath + " -T 5 -t 1"
	var info DownInfo
	info.Url = url
	info.KeepPath = keepPath
	info.IsZip = isZip
	c.status[str_wget] = info
	c.Run(str_wget, str)
	return c.getResult(str_wget).(bool)
}

//wget
func (c *Cmds) checkDownfile(msg string) {
	if strings.Contains(msg, "saved") {
		info := c.status[str_wget]
		rst := false
		if info.(DownInfo).IsZip {
			fmt.Println("下载完成---正在解压")
			if info.(DownInfo).ZipType != Zip_tar {
				rst = c.cmd_unzipFile(info.(DownInfo).KeepPath)
			} else {
				rst = c.cmd_untarFile(info.(DownInfo).KeepPath)
			}
		} else {
			rst = true
			fmt.Println("下载完成")
		}
		c.setResult(str_wget, rst)
	} else {
		c.setResult(str_wget, false)
		info := c.status[str_wget]
		fmt.Println("下载失败", info.(DownInfo).Url)
	}
}
