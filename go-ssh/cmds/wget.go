package cmds

import (
	"fmt"
	"strings"
)

//wget
func (c *Cmds) Cmd_downFile(url, keepPath string, isZip bool, zipType int) {
	c.BindCheck(str_wget, c.checkDownfile)
	str := "wget " + url + " -O " + keepPath + " -T 5 -t 1"
	var info DownInfo
	info.Url = url
	info.KeepPath = keepPath
	info.IsZip = isZip
	c.status[str_wget] = info
	c.Run(str_wget, str)
}

//wget
func (c *Cmds) checkDownfile(msg string) {
	if strings.Contains(msg, str_err) {
		info := c.status[str_wget]
		fmt.Println("下载失败", info.(DownInfo).Url)
		return
	}
	if strings.Contains(msg, "saved") {
		info := c.status[str_wget]
		if info.(DownInfo).IsZip {
			fmt.Println("下载完成---正在解压")
			if info.(DownInfo).ZipType != Zip_tar {
				c.cmd_unzipFile(info.(DownInfo).KeepPath)
			} else {
				c.cmd_untarFile(info.(DownInfo).KeepPath)
			}
		} else {
			fmt.Println("下载完成")
		}
	}
}
