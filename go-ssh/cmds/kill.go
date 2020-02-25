package cmds

import "strings"

func (c *Cmds) Cmd_killRunPortPID(port int) bool {
	pidInfo := c.Cmd_lsofPort(port)
	if pidInfo == "" {
		return true
	}
	data := strings.Split(pidInfo, "\n")
	pid := ""
	for _, v := range data {
		info := strings.Fields(v)
		if len(info) > 2 && pid != info[1] {
			pid = info[1]
			c.Cmd_killPID(pid)
			return c.Cmd_lsofPID(pid)
		}
	}
	return false
}

func (c *Cmds) Cmd_killPID(pid string) bool {
	c.BindCheck(str_kill, c.checkKill)
	sh := "kill " + pid
	c.Run(str_kill, sh)
	return c.getResult(str_kill).(bool)
}
func (c *Cmds) checkKill(msg string) {
	c.setResult(str_kill, true)
}
