package nginx

import (
	"../define"
	"../run"
)

const (
	str_nginx        = "nginx"
	str_nginxCheck   = "nginxCheck"
	str_unzip        = "unzip"
	str_unzipCheck   = "unzipCheck"
	str_wget         = "wget"
	str_unzipFile    = "unzipFile"
	str_nginxReStart = "nginxReStart"
	str_nginxReload  = "nginxReload"
	str_certbot      = "certbot"
	str_certbotCheck = "certbotCheck"
	str_certNew      = "certNew"
	str_certRenew    = "certRenew"
)

type Nginx struct {
	keepPath        string
	isZip, isReload bool
	run             *run.DoRunInfo
	PrintMsg        func(act int, msg string)
}

func NewNginx() *Nginx {
	ret := new(Nginx)
	ret.run = run.NewDoRunInfo()
	ret.run.PrintMsg = ret.printMsg
	return ret
}

func (c *Nginx) IsInstallNginx() {
	if c.run.IsRun(str_nginx) || c.run.IsRun(str_nginxCheck) {
		c.PrintMsg(define.Msg_nginx, "正在安装...")
	} else {
		c.run.CheckRun(str_nginxCheck, "nginx", []string{"-v"})
	}
}
func (c *Nginx) ReStartNginx() {
	if c.run.IsRun(str_nginxReStart) || c.run.IsRun(str_nginxReload) {
		c.PrintMsg(define.Msg_nginxRestart, "正在重启...")
	} else {
		c.run.Run(str_nginxReStart, "nginx", []string{"-t"})
	}
}
func (c *Nginx) reloadNginx() {
	c.run.Run(str_nginxReload, "nginx", []string{"-s", "reload"})
}

func (c *Nginx) installNginx() {
	if c.run.IsRun(str_nginx) {
		return
	}
	c.run.Run(str_nginx, "apt", []string{"install", "nginx", "-y"})
}

func (c *Nginx) IsInstallunzip() {
	c.run.CheckRun(str_unzipCheck, "unzip", []string{})
}

func (c *Nginx) installUnzip() {
	if c.run.IsRun(str_unzip) {
		return
	}
	c.run.Run(str_unzip, "apt", []string{"install", "unzip", "-y"})
}

func (c *Nginx) IsInstallcertbot() {
	c.run.CheckRun(str_certbotCheck, "certbot", []string{"-h"})
}

func (c *Nginx) installcertbot() {
	if c.run.IsRun(str_certbot) {
		return
	}
	c.run.Run(str_certbot, "apt", []string{"install", "certbot", "python-certbot-nginx", "-y"})
}
