package cmds

import (
	"fmt"
	"strings"
)

func (c *Cmds) Cmd_checkNginx() {
	c.BindCheck(str_nginxCheck, c.checkNginx)
	c.Run(str_nginxCheck, "nginx -v")
}

func (c *Cmds) cmd_installNginx() {
	c.BindCheck(str_nginx, c.checkDeleteAllow)
	c.Run(str_nginx, "apt install nginx -y")
}

func (c *Cmds) Cmd_nginxReStart() bool {
	c.BindCheck(str_nginxReStart, c.checkNginxConf)
	c.Run(str_nginxReStart, "nginx -t")
	return c.getResult(str_nginxReStart).(bool)
}

func (c *Cmds) Cmd_relaodNginx() bool {
	c.BindCheck(str_nginxReload, c.checkNginxReload)
	c.Run(str_nginxReload, "nginx -s reload")
	return c.getResult(str_nginxReload).(bool)
}

//----------------------

func (c *Cmds) checkNginx(msg string) {
	if !strings.Contains(msg, str_err) {
		fmt.Println("已安装ngix")
	} else {
		fmt.Println("正在安装ngix...")
		c.cmd_installNginx()
	}
}

func (c *Cmds) checkNginxInstall(msg string) {
	if strings.Contains(msg, "Setting up nginx") {
		fmt.Println("安装nginx成功")
	} else {
		fmt.Println("安装nginx失败")
	}
}

func (c *Cmds) checkNginxConf(msg string) {
	if strings.Contains(msg, str_err) {
		fmt.Println("配置文件错误")
		c.setResult(str_nginxReStart, false)
		return
	}
	rst := c.Cmd_relaodNginx()
	c.setResult(str_nginxReStart, rst)
}
func (c *Cmds) checkNginxReload(msg string) {
	if msg != "" {
		fmt.Println("失败", msg)
		c.setResult(str_nginxReload, false)
		return
	}
	c.setResult(str_nginxReload, true)
	fmt.Println("已重启nginx")
}
