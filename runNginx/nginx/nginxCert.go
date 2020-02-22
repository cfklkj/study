package nginx

import (
	"../define"
	util "../module/util/file"
)

func (c *Nginx) MakeCert(domain string) {
	if c.run.IsRun(str_certNew) {
		c.PrintMsg(define.Msg_certNew, "正在生成证书...")
		return
	}
	if util.FindTag("/etc/letsencrypt/live/", domain) {
		c.PrintMsg(define.Msg_certNew, "证书已存在...")
		return
	}
	c.PrintMsg(define.Msg_certNew, "正在生成证书...")
	c.run.Run(str_certNew, "certbot", []string{"--nginx", "certonly", "-d", domain})
}

func (c *Nginx) RenewCert() {
	if c.run.IsRun(str_certRenew) {
		c.PrintMsg(define.Msg_certRenew, "正在更新证书...")
		return
	}
	c.PrintMsg(define.Msg_certRenew, "正在更新证书...")
	c.run.Run(str_certRenew, "certbot", []string{"renew"})
}
