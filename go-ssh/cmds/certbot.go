package cmds

import (
	"fmt"
	"strings"
)

//certbot
func (c *Cmds) Cmd_checkCertbot() {
	c.BindCheck(str_certbotCheck, c.checkCertbot)
	c.Run(str_certbotCheck, "certbot -h")
}

func (c *Cmds) cmd_installCertbot() {
	c.BindCheck(str_certbot, c.checkCertbotInstall)
	c.Run(str_certbot, "apt install certbot python-certbot-nginx -y")
}

func (c *Cmds) Cmd_makeCert(domain string) {
	c.BindCheck(str_certNew, c.checkNewCert)
	str := "certbot --nginx certonly -d " + domain
	c.Run(str_certNew, str)
}

func (c *Cmds) cmd_RenewCert() {
	c.BindCheck(str_certRenew, c.checkRenewCert)
	c.Run(str_certRenew, "certbot renew")
}

func (c *Cmds) cmd_lsCert() {
	c.Run(str_certls, "ls /etc/letsencrypt/renewal")
}

//certbot
func (c *Cmds) checkCertbot(msg string) {
	if !strings.Contains(msg, str_err) {
		fmt.Println("已安装certbot")
		return
	} else {
		fmt.Println("正在安装certbot")
		c.cmd_installCertbot()
	}
	return
}
func (c *Cmds) checkCertbotInstall(msg string) {
	if strings.Contains(msg, "Setting up certbot") {
		fmt.Println("安装certbot成功")
	}
}
func (c *Cmds) checkNewCert(msg string) {
	if strings.Contains(msg, str_err) {
		c.cmd_RenewCert()
	} else {
		c.cmd_lsCert()
	}
}
func (c *Cmds) checkRenewCert(msg string) {
	c.cmd_lsCert()
}
