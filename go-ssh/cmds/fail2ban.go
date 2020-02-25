package cmds

import (
	"fmt"
	"strings"
)

func (c *Cmds) Cmd_checkFail2ban() {
	c.BindCheck(str_fail2banCheck, c.checkFail2ban)
	sh := "ls /lib/systemd/system/fail2ban.service"
	c.Run(str_fail2banCheck, sh)
}

//c.DownFile("https://codeload.github.com/fail2ban/fail2ban/tar.gz/0.9.1", "./")
func (c *Cmds) cmd_installFail2ban() {
	c.BindCheck(str_fail2ban, c.checkFail2banInstall)
	sh := "apt install fail2ban"
	c.Run(str_fail2ban, sh)
}

func (c *Cmds) Cmd_restartFail2ban() bool {
	c.BindCheck(str_fail2banRestart, c.checkFail2banRestart)
	sh := "service fail2ban restart"
	c.Run(str_fail2banRestart, sh)
	return c.getResult(str_fail2banRestart).(bool)
}

func (c *Cmds) checkFail2ban(msg string) {
	if strings.Contains(msg, "cannot access") {
		fmt.Println("正在安装fail2ban")
		c.cmd_installFail2ban()
	} else {
		fmt.Println("已安装fail2ban")
	}
}
func (c *Cmds) checkFail2banInstall(msg string) {
	if strings.Contains(msg, "Setting up fail2ban") {
		fmt.Println("安装fail2ban完成")
	} else {
		fmt.Println("安装fail2ban失败")
	}
}
func (c *Cmds) checkFail2banRestart(msg string) {
	if strings.Contains(msg, "Failed") {
		c.setResult(str_fail2banRestart, false)
	} else {
		c.setResult(str_fail2banRestart, true)
	}
}
