package ssh

import (
	"net"
	"time"

	gossh "golang.org/x/crypto/ssh"
)

type clientInfo struct {
	user, pwd, addr string
}

type CSSH struct {
	info     clientInfo
	client   *gossh.Client
	session  *gossh.Session
	PrintMsg func(name, msg string)
}

func NewCSSH() *CSSH {
	ret := new(CSSH)
	return ret
}

func (c *CSSH) Clear() {
	if c.client != nil {
		if c.session != nil {
			c.session.Close()
		}
		c.client.Close()
	}
}

func (c *CSSH) Connect(user, pwd, addr string) (*CSSH, error) {
	config := &gossh.ClientConfig{}
	config.SetDefaults()
	config.User = user
	config.Timeout = 5 * time.Second
	config.Auth = []gossh.AuthMethod{gossh.Password(pwd)}
	config.HostKeyCallback = func(hostname string, remote net.Addr, key gossh.PublicKey) error { return nil }
	client, err := gossh.Dial("tcp", addr, config)
	if nil != err {
		return c, err
	}
	c.info = clientInfo{user, pwd, addr}
	c.client = client
	return c, nil
}

func (c *CSSH) Run(name, shell string) {
	if c.client == nil {
		if _, err := c.Connect(c.info.user, c.info.pwd, c.info.addr); err != nil {
			c.PrintMsg(name, err.Error())
			return
		}
	}
	session, err := c.client.NewSession()
	if err != nil {
		c.PrintMsg(name, err.Error())
		return
	}
	defer session.Close()
	buf, err := session.CombinedOutput(shell)
	if err != nil {
		c.PrintMsg(name, string(buf)+"---err---:"+err.Error())
		return
	}
	c.PrintMsg(name, string(buf))
}
