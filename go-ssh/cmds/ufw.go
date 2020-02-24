package cmds

import (
	"fmt"
	"strconv"
	"strings"
)

func (c *Cmds) Cmd_checkUfw() {
	c.BindCheck(str_ufwCheck, c.checkUfw)
	c.Run(str_ufwCheck, "ufw version")
}

func (c *Cmds) cmd_installUfw() {
	c.BindCheck(str_ufw, c.checkUfwInstall)
	c.Run(str_ufw, "apt install ufw -y")
}

func (c *Cmds) Cmd_ufwAllow(port int) {
	c.BindCheck(str_ufwAllow, c.checkAllow)
	c.Run(str_ufwAllow, "ufw allow "+strconv.Itoa(port))
}
func (c *Cmds) cmd_ufwEnable() {
	c.BindCheck(str_ufwEnable, c.checkEnable)
	c.Run(str_ufwEnable, "echo y|ufw enable")
}
func (c *Cmds) Cmd_ufwDelAllow(port int) {
	c.BindCheck(str_ufwDelAllow, c.checkDeleteAllow)
	c.Run(str_ufwDelAllow, "ufw delete allow "+strconv.Itoa(port))
}

//-------------------check
func (c *Cmds) checkUfw(msg string) {
	if !strings.Contains(msg, str_err) {
		fmt.Println("已安装ufw")
		return
	}
	fmt.Println("正在安装ufw...")
	c.cmd_installUfw()
}

func (c *Cmds) checkUfwInstall(msg string) {
	if strings.Contains(msg, "Setting up ufw") {
		fmt.Println("安装ufw成功")
		c.Cmd_ufwAllow(22) //打开ssh端口
		c.cmd_ufwEnable()
	}
}

func (c *Cmds) checkEnable(msg string) {
	if strings.Contains(msg, "Firewall is active") {
		fmt.Println("ufw", "已启动ufw")
	} else {
		fmt.Println("ufw", msg)
	}
}

func (c *Cmds) checkAllow(msg string) {
	if strings.Contains(msg, "Rule added") {
		fmt.Println("ufw", "Rule added")
	} else {
		fmt.Println("ufw", msg)
	}
}
func (c *Cmds) checkDeleteAllow(msg string) {
	if strings.Contains(msg, "Rule deleted") {
		fmt.Println("ufw", "Rule deleted")
	} else {
		fmt.Println("ufw", msg)
	}
}
