package ssh

import (
	"fmt"
	"net"
	"os"
	"path"
	"time"

	"github.com/pkg/sftp"
	gossh "golang.org/x/crypto/ssh"
)

type clientInfo struct {
	user, pwd, addr string
}

type CSSH struct {
	info     clientInfo
	client   *gossh.Client
	PrintMsg func(name, msg string)
}

func NewCSSH() *CSSH {
	ret := new(CSSH)
	return ret
}

func (c *CSSH) Clear() {
	if c.client != nil {
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

func (c *CSSH) Run(name, shell string) bool {
	if c.client == nil {
		if _, err := c.Connect(c.info.user, c.info.pwd, c.info.addr); err != nil {
			c.PrintMsg(name, err.Error())
			return false
		}
	}
	session, err := c.client.NewSession()
	if err != nil {
		c.PrintMsg(name, err.Error())
		return false
	}
	defer session.Close()
	buf, err := session.CombinedOutput(shell)
	if err != nil {
		c.PrintMsg(name, string(buf)+"---err---:"+err.Error())
		return false
	}
	c.PrintMsg(name, string(buf))
	return true
}

func (c *CSSH) SCPupFile(localFilePath, remoteDir string) bool {
	if c.client == nil {
		if _, err := c.Connect(c.info.user, c.info.pwd, c.info.addr); err != nil {
			return false
		}
	}
	sftpClient, err := sftp.NewClient(c.client)
	if err != nil {
		return false
	}
	srcFile, err := os.Open(localFilePath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer srcFile.Close()

	var remoteFileName = path.Base(localFilePath)
	dstFile, err := sftpClient.Create(path.Join(remoteDir, remoteFileName))
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer dstFile.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf[0:n])
	}
	return true
}

func (c *CSSH) SCPDownFile(remoteFilePath, localDir string) bool {
	if c.client == nil {
		if _, err := c.Connect(c.info.user, c.info.pwd, c.info.addr); err != nil {
			return false
		}
	}
	sftpClient, err := sftp.NewClient(c.client)
	if err != nil {
		return false
	}

	srcFile, err := sftpClient.Open(remoteFilePath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer srcFile.Close()

	var remoteFileName = path.Base(remoteFilePath)
	dstFile, err := os.Create(path.Join(localDir, remoteFileName))
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer dstFile.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf[0:n])
	}
	return true
}
