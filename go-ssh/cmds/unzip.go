package cmds

import (
	"fmt"
	"path/filepath"
	"strings"
)

//---unzip
func (c *Cmds) Cmd_checkUnzip() {
	c.BindCheck(str_unzipCheck, c.checkUnzip)
	c.Run(str_unzipCheck, "unzip")
}
func (c *Cmds) cmd_installUnzip() {
	c.BindCheck(str_unzip, c.chckUnzipInstall)
	c.Run(str_unzip, "apt install unzip -y")
}

func (c *Cmds) cmd_unzipFile(filePath string) bool {
	c.BindCheck(str_unzipFile, c.chckUnzipFile)
	dir, _ := filepath.Split(filePath)
	str := "unzip -o " + filePath + " -d " + dir
	c.Run(str_unzipFile, str)
	return c.getResult(str_unzipFile).(bool)
}
func (c *Cmds) cmd_untarFile(filePath string) bool {
	c.BindCheck(str_unzipFile, c.chckUnzipFile)
	dir, _ := filepath.Split(filePath)
	str := "cd " + dir +
		"\ntar zxvf " + filePath
	c.Run(str_unzipFile, str)
	return c.getResult(str_unzipFile).(bool)
}

//unzip
func (c *Cmds) checkUnzip(msg string) {
	if !strings.Contains(msg, str_err) {
		fmt.Println("已安装unzip")
	} else {
		fmt.Println("正在安装unzip")
		c.cmd_installUnzip()
	}
}
func (c *Cmds) chckUnzipInstall(msg string) {
	if strings.Contains(msg, "Setting up unzip") {
		fmt.Println("安装unzip成功")
	} else {
		fmt.Println("安装unzip失败")
	}
}
func (c *Cmds) chckUnzipFile(msg string) {
	if strings.Contains(msg, str_err) {
		fmt.Println("解压失败", msg)
		c.setResult(str_unzipFile, false)
		return
	}
	c.setResult(str_unzipFile, true)
	fmt.Println("解压成功")
}
