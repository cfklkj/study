package nginx

import (
	"path/filepath"

	"../define"
)

func (c *Nginx) IsDownFile() {
	if c.run.IsRun(str_wget) {
		c.PrintMsg(define.Msg_down, "正在下载...")
	} else {
		c.run.Run(str_wget, "wget", []string{"install", "unzip"})
	}
}
func (c *Nginx) DownFile(url, keepPath string, isZip bool) {
	if c.run.IsRun(str_wget) || c.run.IsRun(str_unzipFile) {
		c.PrintMsg(define.Msg_down, "正在下载...")
	} else {
		c.PrintMsg(define.Msg_down, "正在执行下载...")
		c.keepPath = keepPath
		c.isZip = isZip
		c.run.Run(str_wget, "wget", []string{url, "-O", keepPath, "-T", "5", "-t", "1"})
	}
}

func (c *Nginx) unZip() {
	dir, _ := filepath.Split(c.keepPath)
	c.run.Run(str_unzipFile, "unzip", []string{"-o", c.keepPath, "-d", dir})
}
