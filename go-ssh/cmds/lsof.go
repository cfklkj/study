package cmds

import (
	"strconv"
	"strings"
)

func (c *Cmds) Cmd_lsofPort(port int) string {
	c.BindCheck(str_lsofPort, c.checkRunPort)
	sh := "lsof -R -i:" + strconv.Itoa(port) + " | grep \" 1 \""
	c.Run(str_lsofPort, sh)
	return c.getResult(str_lsofPort).(string)
}

func (c *Cmds) Cmd_lsofPID(pid string) bool {
	c.BindCheck(str_lsofPid, c.checkRunPid)
	sh := "lsof -i" + " | grep " + pid
	c.Run(str_lsofPid, sh)
	return c.getResult(str_lsofPid).(bool)
}

func (c *Cmds) checkRunPid(msg string) {
	if strings.Contains(msg, str_err) {
		c.setResult(str_lsofPid, true)
	} else {
		c.setResult(str_lsofPid, false)
	}
}

func (c *Cmds) checkRunPort(msg string) {
	c.setResult(str_lsofPort, "")
	if !strings.Contains(msg, str_err) {
		c.setResult(str_lsofPort, msg)
	}
}
