package nginx

import (
	"fmt"

	"../define"
)

func (c *Nginx) printMsg(name, msg string) {
	fmt.Println(name, msg)
	if name == str_certNew {
		if msg == "err" {
			c.PrintMsg(define.Msg_certNew, "生成cert失败")
		} else {
			c.PrintMsg(define.Msg_certNew, msg)
		}
		return
	}
	if name == str_certRenew {
		if msg == "err" {
			c.PrintMsg(define.Msg_certRenew, "更新cert失败...")
		} else {
			c.PrintMsg(define.Msg_certRenew, msg)
		}
		return
	}
	if name == str_certbotCheck {
		if msg == "unInstall" {
			c.PrintMsg(define.Msg_local, "正在安装certbot...")
			c.installcertbot()
		} else {
			c.PrintMsg(define.Msg_local, "certbot已安装")
		}
		return
	}
	if name == str_certbot {
		if msg == "err" {
			c.PrintMsg(define.Msg_local, "安装certbot失败...")
		} else {
			c.PrintMsg(define.Msg_local, "安装certbot成功...")
		}
		return
	}
	if name == str_nginxCheck {
		if msg == "unInstall" {
			c.PrintMsg(define.Msg_nginx, "正在安装nginx...")
			c.installNginx()
		} else {
			c.PrintMsg(define.Msg_nginx, "nginx已安装")
		}
		return
	}
	if name == str_nginx {
		if msg == "err" {
			c.PrintMsg(define.Msg_nginx, "安装nginx失败...")
		} else {
			c.PrintMsg(define.Msg_nginx, "安装nginx成功...")
		}
		return
	}
	if name == str_unzipCheck {
		if msg == "unInstall" {
			c.PrintMsg(define.Msg_local, "正在安装unzip...")
			c.installUnzip()
		} else {
			c.PrintMsg(define.Msg_local, "unzip已安装")
		}
		return
	}
	if name == str_unzip {
		if msg == "err" {
			c.PrintMsg(define.Msg_local, "安装unzip失败...")
		} else {
			c.PrintMsg(define.Msg_local, "安装unzip成功")
		}
		return
	}
	if name == str_wget {
		if msg == "saved" {
			c.PrintMsg(define.Msg_down, "已下载完成...")
			if c.isZip {
				c.PrintMsg(define.Msg_down, "正在解压...")
				c.unZip()
			}
		} else if msg == "err" {
			c.PrintMsg(define.Msg_down, "下载失败...")
		}
		return
	}
	if name == str_unzipFile {
		if msg == "err" {
			c.PrintMsg(define.Msg_down, "解压失败...")
		} else {
			c.PrintMsg(define.Msg_down, "解压成功...")
		}
	}
	if name == str_nginxReStart {
		if msg == "err" {
			c.PrintMsg(define.Msg_nginxRestart, "重启nginx失败...")
		} else {
			c.PrintMsg(define.Msg_nginxRestart, "正在重启nginx...")
			c.reloadNginx()
		}
		return
	}
	if name == str_nginxReload {
		if msg == "err" {
			c.PrintMsg(define.Msg_nginxRestart, "重启nginx失败...")
		} else {
			c.PrintMsg(define.Msg_nginxRestart, "重启nginx成功...")
		}
	}
}
